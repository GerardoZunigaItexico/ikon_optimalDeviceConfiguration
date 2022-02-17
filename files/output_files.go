package files

import (
	"errors"
	"fmt"
	"ikon_optimalDeviceConfiguration/device"
	"os"

	"strings"
)

func CreateChallengeOut(devices []device.Device) (fileLocation string, err error){
	path, err := os.Getwd()
	if err != nil {
		return fileLocation, err
	}
	fileLocation = path+"/challenge.out"
	err = deleteIfExists(fileLocation)
	if err != nil {
		return fileLocation, err
	}

	f, err := os.Create(fileLocation)

	if err != nil {
		return fileLocation, err
	}
	defer f.Close()

	_, err2 := f.WriteString(createOutputStringForOutputFile(devices))

	if err2 != nil {
		return fileLocation, err
	}

	return fileLocation, nil
}

func createOutputStringForOutputFile(devices []device.Device) (lines string){
	for _,device := range devices{
		lines += fmt.Sprintln(createOutputDeviceConf(device))
	}
	return lines
}


func createOutputDeviceConf(device device.Device) (idsLine string){
	var ids []string
	if device.DeviceConf != nil{
		for _,dc := range *device.DeviceConf{
			ids = append(ids,fmt.Sprintf("(%d, %d)",dc.Foreground.Id, dc.Background.Id))
		}

	} else{
		if device.NonExactDeviceConf != nil{
			for _,dc := range *device.NonExactDeviceConf{
				ids = append(ids,fmt.Sprintf("(%d, %d)",dc.Foreground.Id, dc.Background.Id))
			}

		}
	}
	idsLine = strings.Join(ids,", ")
	return
}

func deleteIfExists(fileLocation string) error {
	_, err := os.Stat(fileLocation)
	if !errors.Is(err, os.ErrNotExist){
		err := os.Remove(fileLocation)  // remove a single file
		if err != nil {
			return err
		}
	}
	return nil
}