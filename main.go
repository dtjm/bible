package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	esvBaseURL = "http://www.esvapi.org/v2/rest"
	version    = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "ment"
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
				query := url.Values{
					"key":                        {"IP"},
					"output-format":              {"plain-text"},
					"passage":                    {strings.Join([]string(c.Args()), " ")},
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
