package main

import (
	"fmt"
	"github.com/akito0107/unienv"
	"github.com/urfave/cli"
	"io"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "unienv"
	app.Usage = "unify environment variables"
	app.UsageText = "generr [subcommand] [OPTIONS]"

	app.Commands = []cli.Command{
		{
			Name:   "unify",
			Usage:  "unify env",
			Action: unify,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "env", Usage: "environment name (required)"},
				cli.StringFlag{Name: "input, i", Usage: "input file name (default: unienv.ini)", Value: "unienv.ini"},
				cli.StringFlag{Name: "output, o", Usage: "output file path (default: .env)", Value: ".env"},
				cli.BoolFlag{Name: "dry-run"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func unify(ctx *cli.Context) error {
	envname := ctx.String("env")
	if envname == "" {
		return &unienv.NoEnvName{}
	}

	f, err := os.Open(ctx.String("input"))
	if err != nil {
		return err
	}
	conf, err := unienv.Parse(f)
	if err != nil {
		return err
	}
	merged, err := unienv.Merge(envname, conf)
	if err != nil {
		return err
	}
	if ctx.Bool("dry-run") {
		print(os.Stdout, merged)
		return nil
	}

	out, err := os.Create(ctx.String("output"))
	defer out.Close()
	print(out, merged)
	return nil
}

func print(w io.Writer, merged map[string]string) {
	for k, v := range merged {
		fmt.Fprintf(w, "%s=%s\n", k, v)
	}
}
