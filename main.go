package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Conventional Commits Autoversion"
	app.Usage = "Utility to obtain the new semantic version of a git repository based on Conventional Commits"
	app.Author = "Local Line"
	app.Version = "0.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Commands to obtain versions from tags in a git repository",
			Subcommands: []cli.Command{
				{
					Name:      "previous",
					Aliases:   []string{"p"},
					Usage:     "Obtains the last semver tag in the git history",
					ArgsUsage: "<path>",
					Action: func(c *cli.Context) {
						repo := openRepo(c.Args().Get(0))
						tag, _ := latestSemverTag(repo)
						fmt.Println(tag.toString())
					},
				},
				{
					Name:      "next",
					Aliases:   []string{"n"},
					Usage:     "Returns the next version number based on commit names",
					ArgsUsage: "<path>",
					Action: func(c *cli.Context) {
						repo := openRepo(c.Args().Get(0))
						currentSemver, _ := latestSemverTag(repo)
						sinceRef := tagRef(repo, currentSemver)
						commits := commitsBetweenRefs(repo, sinceRef, nil)
						newSemver := semverBump(commits, currentSemver)
						fmt.Println(newSemver.toString())
					},
				},
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
