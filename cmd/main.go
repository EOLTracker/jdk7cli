package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name: "jdk7",
    Usage: "Counts down the days, minutes, seconds, and Unix epoches until JDK7 is deprecated.",
    Action: func(c *cli.Context) error {
      fmt.Println("Run jdk7 --help or jdk7 -h to see the available commands.")
      return nil
    },
  }

func main() {
	app := &cli.App{
	  Flags: []cli.Flag {
		&cli.StringFlag{
		  Name: "lang",
		  Value: "english",
		  Usage: "language for the greeting",
		},
	  },

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}