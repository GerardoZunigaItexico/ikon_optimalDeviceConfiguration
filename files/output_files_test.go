package files

import (
	"fmt"
	"ikon_optimalDeviceConfiguration/device"
	"log"
	"os"
	"testing"
)

func Test_createOutputDeviceConf(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputData := ParseChallengeIn(path+"/testFiles/challenge.in")
	for _,i := range inputData{
		device := device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		fmt.Println(createOutputDeviceConf(device))
	}
}

func Test_createOutputStringForOutputFile(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputData := ParseChallengeIn(path+"/testFiles/challenge.in")
	var devices []device.Device
	for _,i := range inputData{
		devices = append(devices, device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground))
	}
	fmt.Println(createOutputStringForOutputFile(devices))
}

func Test_CreateChallengeOut(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputData := ParseChallengeIn(path+"/testFiles/challenge.in")
	var devices []device.Device
	for _,i := range inputData{
		devices = append(devices, device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground))
	}
	fmt.Println(CreateChallengeOut(devices))
}