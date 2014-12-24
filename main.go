package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/dtjm/bible/ref"
)

const (
	esvBaseURL = "http://www.esvapi.org/v2/rest"
	version    = "0.0.3"
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
		c.Bookmarks["next"] = "Gen 1"
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

	nextRef := r.NextChapter()
	return nextRef.String()
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

	app.Commands = []cli.Command{
		{
			Name:      "read",
			ShortName: "r",
			Usage:     "Read a passage",
			Action: func(c *cli.Context) {
				var ref = strings.Join([]string(c.Args()), " ")
				if strings.ToLower(ref) == "next" {
					ref = conf.Bookmarks["next"]

					conf.Bookmarks["next"] = nextRef(ref)
					conf.write()
				}

				query := url.Values{
					"key":                        {"IP"},
					"output-format":              {"plain-text"},
					"passage":                    {ref},
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
			},
		},
	}

	app.Run(os.Args)
}
