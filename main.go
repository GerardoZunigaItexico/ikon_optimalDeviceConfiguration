package main

import (
	"fmt"
	"github.com/urfave/cli"
	"ikon_optimalDeviceConfiguration/cmd"
	"os"
)

func main(){
	app := cli.NewApp()
	app.Name = "optimal-onfigurator"
	app.Usage = "Utility to import products"

	app.Commands = []cli.Command{
		cmd.NewGetOptimalConfiguration(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
