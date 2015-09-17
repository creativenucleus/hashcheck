package main

import (
	"crypto/md5"
	"fmt"
	"github.com/codegangsta/cli"
	"io"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hashcheck"
	app.Usage = "Get the MD5 hash of a file"
	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "Get the hash",
			Action: cliActionGet(),
		},
	}

	app.Run(os.Args)
}

func cliActionGet() func(*cli.Context) {
	return func(c *cli.Context) {
		filepath := c.Args().First()
		if filepath == "" {
			fmt.Println("Please supply a file")
			return
		}

		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println("Could not open the file")
		}
		defer file.Close()

		var result []byte
		hash := md5.New()
		if _, err := io.Copy(hash, file); err != nil {
			return
		}

		fmt.Printf("MD5 %x", hash.Sum(result))
	}
}
