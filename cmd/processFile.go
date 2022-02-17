package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"ikon_optimalDeviceConfiguration/device"
	"ikon_optimalDeviceConfiguration/files"
)

func NewGetOptimalConfiguration() cli.Command {
	return cli.Command{
		Name:   "configure",
		Usage:  "Get Optimal Configuration for Device",
		Action: configureDeviceConfigurations,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "filePath",
				Usage: "File Input Location that will be parsed.  This one is mandatory",
				Required: true,
			},
			cli.StringFlag{
				Name:  "delta",
				Value: "0",
				Usage: "Delta Device Capacity",
			},
		},
	}
}

func configureDeviceConfigurations(c *cli.Context) error {
	filePath 	 := c.String("filePath")
	deltaDevCap := c.Int("delta")

	inputData, err := files.ParseChallengeIn(filePath)
	if err != nil {
		return err
	}
	devices, err := createDeviceCombinationsMultipleInputs(inputData,deltaDevCap)
	if err != nil{
		return err
	}
	outputFilePath,err := files.CreateChallengeOut(devices)
	if err != nil{
		return err
	}
	fmt.Println(outputFilePath)
	return nil
}

func createDeviceCombinationsMultipleInputs(inputData []files.InputData, deltaDevCap int) (devices []device.Device, err error){
	for _,i := range inputData{
		device, err := device.CreateDeviceCombinations(i.Capacity,deltaDevCap,i.Background,i.Foreground)
		if err != nil{
			return devices, err
		}
		devices = append(devices, device)
	}
	return devices, err
}