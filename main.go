package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/facebookgo/counting"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.google.com/p/portaudio-go/portaudio"
	"github.com/BurntSushi/toml"
	"github.com/Wessie/audec/mp3"
	"github.com/codegangsta/cli"
	"github.com/dtjm/bible/ref"
)

const (
	esvBaseURL = "http://www.esvapi.org/v2/rest"
	version    = "0.0.4"
)

var (
	configFile = os.Getenv("HOME") + "/.bible"
)

type config struct {
	Bookmarks map[string]string `toml:"bookmarks"`
}

func readConfig() *config {
	c := config{Bookmarks: make(map[string]string)}
	_, err := os.Stat(configFile)

	if os.IsNotExist(err) {
		log.Printf("config file does not exist, creating %q", configFile)
		c.Bookmarks["next"] = "Genesis 1"
		return &c
	}

	_, err = toml.DecodeFile(configFile, &c)
	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func (c *config) write() error {
	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	return enc.Encode(c)
}

func nextRef(s string) string {
	r, err := ref.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	// nextRef := r.NextChapter()
	return r.NextChapter().String()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conf := readConfig()

	app := cli.NewApp()
	app.Name = "bible"
	app.Usage = "a command-line tool for reading the Bible"
	app.Author = "Sam Nguyen"
	app.Email = "samxnguyen@gmail.com"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "verbose", Usage: "enable verbose logging"},
	}

	app.Before = func(c *cli.Context) error {
		if c.Bool("verbose") {
			log.SetOutput(os.Stderr)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:      "read",
			ShortName: "r",
			Usage:     "Read a passage",
			Action: func(c *cli.Context) {
				var refString = strings.Join([]string(c.Args()), " ")
				if strings.ToLower(refString) == "next" {
					refString = conf.Bookmarks["next"]
					conf.Bookmarks["next"] = nextRef(refString)
					log.Printf("next:\t%s", refString)
				}

				if refString == "" {
					cli.ShowCommandHelp(c, c.Command.Name)
					return
				}

				query := url.Values{
					"key":                        {"IP"},
					"output-format":              {"plain-text"},
					"passage":                    {refString},
					"include-headings":           {"0"},
					"include-subheadings":        {"0"},
					"include-passage-references": {"0"},
					"include-verse-numbers":      {"0"},
					"include-footnotes":          {"0"},
				}

				resp, err := http.Get(esvBaseURL + "/passageQuery?" + query.Encode())
				if err != nil {
					log.Fatal(err)
				}

				io.Copy(os.Stdout, resp.Body)
				resp.Body.Close()
				fmt.Print("\n\n")

				parsedRef, err := ref.Parse(refString)
				if err != nil {
					log.Fatal(err)
				}
				conf.Bookmarks["last"] = parsedRef.String()
				conf.write()
			},
		},

		{
			Name:      "mark",
			ShortName: "m",
			Usage:     "Bookmark a passage",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					for m, r := range conf.Bookmarks {
						fmt.Printf("%s:\t%s\n", m, r)
					}
					return
				}

				if len(c.Args()) == 1 {
					mark := c.Args()[0]
					if r, ok := conf.Bookmarks[mark]; ok {
						fmt.Printf("%s:\t%s\n", mark, r)
					} else {
						log.Printf("You don't have a bookmark called %q", mark)
					}
					return
				}

				mark := c.Args()[0]
				var refString = strings.Join([]string(c.Args()[1:]), " ")
				r, err := ref.Parse(refString)
				if err != nil {
					log.Fatal(err)
				}

				conf.Bookmarks[mark] = r.String()
				conf.write()
			},
		},

		{
			Name:      "play",
			ShortName: "p",
			Usage:     "Play a reading of a passage",
			Action: func(c *cli.Context) {

				var refString = strings.Join([]string(c.Args()), " ")
				if strings.ToLower(refString) == "next" {
					refString = conf.Bookmarks["next"]
					conf.Bookmarks["next"] = nextRef(refString)
					log.Printf("next:\t%s", refString)
				}

				if refString == "" {
					cli.ShowCommandHelp(c, c.Command.Name)
					return
				}

				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					query := url.Values{
						"key":                        {"IP"},
						"output-format":              {"plain-text"},
						"passage":                    {refString},
						"include-headings":           {"0"},
						"include-subheadings":        {"0"},
						"include-passage-references": {"0"},
						"include-verse-numbers":      {"0"},
						"include-footnotes":          {"0"},
					}

					resp, err := http.Get(esvBaseURL + "/passageQuery?" +
						query.Encode())
					if err != nil {
						log.Fatal(err)
					}

					io.Copy(os.Stdout, resp.Body)
					resp.Body.Close()
					fmt.Print("\n\n")

					parsedRef, err := ref.Parse(refString)
					if err != nil {
						log.Fatal(err)
					}
					conf.Bookmarks["last"] = parsedRef.String()
					conf.write()
					wg.Done()
				}()

				wg.Add(1)
				go func() {
					portaudio.Initialize()
					defer portaudio.Terminate()

					query := url.Values{
						"key":           {"IP"},
						"output-format": {"mp3"},
						"passage":       {refString},
					}

					resp, err := http.Get(esvBaseURL + "/passageQuery?" +
						query.Encode())
					if err != nil {
						log.Fatal(err)
					}

					defer resp.Body.Close()

					countingReader := counting.NewReader(resp.Body)
					mp3Dec, err := mp3.NewDecoder(countingReader)
					if err != nil {
						log.Fatal(err)
					}

					// buf := bytes.NewBuffer(nil)
					// n, err := io.Copy(buf, mp3Dec)
					// if err != nil {
					// 	log.Fatal(err)
					// }
					// log.Printf("Copied %d bytes into buffer", n)
					contentLength, err := strconv.Atoi(resp.Header.Get("Content-Length"))
					if err != nil {
						log.Fatal(err)
					}

					mp3Stream := mp3Stream{
						done:       make(chan struct{}),
						counter:    countingReader,
						br:         bufio.NewReader(mp3Dec),
						totalBytes: contentLength,
					}

					outDevice, err := portaudio.DefaultOutputDevice()
					if err != nil {
						log.Fatal(err)
					}

					paStream, err := portaudio.OpenStream(portaudio.StreamParameters{
						Input: portaudio.StreamDeviceParameters{},
						Output: portaudio.StreamDeviceParameters{
							Device:   outDevice,
							Channels: 1,
							Latency:  0,
						},
						SampleRate:      88200,
						FramesPerBuffer: 16384,
					}, mp3Stream.ProcessAudio)

					if err != nil {
						log.Fatal(err)
					}

					paStream.Start()
					<-mp3Stream.done
					paStream.Stop()
					wg.Done()
				}()

				wg.Wait()
			},
		},
	}

	app.Run(os.Args)
}

type counter interface {
	Count() int
}

type mp3Stream struct {
	br         *bufio.Reader
	once       sync.Once
	done       chan struct{}
	counter    counter
	totalBytes int
}

func (m *mp3Stream) ProcessAudio(_, out []float32) {

	var pack float32
	for i := range out {
		binary.Read(m.br, binary.LittleEndian, &pack)
		out[i] = pack
	}

	// Signal that the playback has started, so we can also signal when it's over
	m.once.Do(func() {
		go func() {
			playTime := time.Duration(float64(m.totalBytes)/4000.0)*time.Second + time.Second
			log.Printf("streaming for %s", playTime)
			time.Sleep(playTime)
			m.done <- struct{}{}
		}()
	})
}
