package files

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
	inputData,err := ParseChallengeIn(path+"/testFiles/challenge.in")
	assert.NoError(t,err, "There is a problem with the parse on the file that you provided")
	for _,i := range inputData{
		device, err := device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		assert.NoError(t,err, "There is a problem with the parse on the file that you provided")
		fmt.Println(createOutputDeviceConf(device))
	}
}

func Test_createOutputStringForOutputFile(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputData, err := ParseChallengeIn(path+"/testFiles/challenge.in")
	assert.NoError(t,err, "There is a problem with the parse on the file that you provided")
	var devices []device.Device
	for _,i := range inputData{
		device, err := device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		assert.NoError(t,err, "There is a problem with the parse on the file that you provided")
		devices = append(devices, device)
	}
	fmt.Println(createOutputStringForOutputFile(devices))
}

func Test_CreateChallengeOut(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputData, err := ParseChallengeIn(path+"/testFiles/challenge.in")
	assert.NoError(t,err, "There is a problem with the parse on the file that you provided")
	var devices []device.Device
	for _,i := range inputData{
		d, err := device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		assert.NoError(t,err, "There is a problem with the Device Parse")
		devices = append(devices, d)
	}
	fmt.Println(CreateChallengeOut(devices))
}