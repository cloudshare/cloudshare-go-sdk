package main

import (
	"fmt"
	cs "github.com/cloudshare/go-sdk/cloudshare"
	"github.com/urfave/cli"
	neturl "net/url"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Name = "CSCURL"

	app.Description = "Invoke CloudShare REST API calls from command-line"
	app.Usage = "CloudShare REST API CLI Utility"
	app.Authors = []cli.Author{{
		Name:  "Assaf Lavie",
		Email: "assaf@cloudshare.com",
	},
	}

	app.Version = "1.1.3"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "method, m",
			Value: "get",
		},
		cli.StringFlag{
			Name:   "api-key",
			Value:  "",
			Usage:  "CloudShare API key",
			EnvVar: "CLOUDSHARE_API_KEY",
		},
		cli.StringFlag{
			Name:   "api-id",
			Value:  "",
			Usage:  "CloudShare API ID",
			EnvVar: "CLOUDSHARE_API_ID",
		},
		cli.BoolFlag{
			Name:  "headers, I",
			Usage: "Print response headers",
		},
		cli.StringFlag{
			Name:  "data, d",
			Value: "",
			Usage: "JSON document",
		},
	}

	app.Action = func(c *cli.Context) error {
		apiKey := c.String("api-key")
		apiID := c.String("api-id")
		if apiKey == "" {
			return fmt.Errorf("api-key must be set")
		}

		if apiID == "" {
			return fmt.Errorf("api-id must be set")
		}

		if c.NArg() < 0 {
			cli.ShowAppHelp(c)
			return fmt.Errorf("Expecting URL argument")
		}

		url := c.Args().Get(0)

		method := c.String("method")

		showHeaders := c.Bool("headers")

		client := &cs.Client{
			APIKey: apiKey,
			APIID:  apiID,
			Tags:   "cscurl",
		}

		data := c.String("data")
		parsed, err := neturl.Parse(url)
		if err != nil {
			return err
		}
		query := parsed.Query()
		client.APIHost = parsed.Host
		path := strings.Replace(parsed.Path, "api/v3/", "", 1)

		response, err := client.Request(method, path, &query, &data)
		if showHeaders {
			fmt.Printf("Status code: %d\n", response.StatusCode)
			for key, value := range response.Headers {
				for _, x := range value {
					fmt.Printf("%s: %s\n", key, x)
				}
			}
			fmt.Println("\n\n")
		}
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(response.Body))
			return err

		}

		fmt.Println(string(response.Body))
		return nil
	}

	app.Run(os.Args)
}
