package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "openintextedit",
		Usage: "Open a file in TextEdit",
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.Exit("Please provide a filename", 1)
			}
			filename := c.Args().Get(0)
			return openInTextEdit(filename)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func openInTextEdit(filename string) error {
	cmd := exec.Command("open", "-a", "TextEdit", filename)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	fmt.Printf("Opened %s in TextEdit\n", filename)
	return nil
}
