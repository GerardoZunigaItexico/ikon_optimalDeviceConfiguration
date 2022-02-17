package device

import (
	"fmt"
	"ikon_optimalDeviceConfiguration/files"
	"testing"
)

func Test_CreateDeviceCombinations(t *testing.T) {
	iD := CreateInputsForDeviceConfs()
	for _,i := range iD{
		device := CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		printDevice(device)
	}
}

func printDevice(device Device){
	fmt.Println(device.Capacity)
	if device.DeviceConf != nil{
		printDeviceConf(device.DeviceConf)
	} else {
		printDeviceConf(device.NonExactDeviceConf)
	}
}

func printDeviceConf(deviceConfigs *[]DeviceConf){
	for _,dc := range *deviceConfigs{
		fmt.Print("\tForeground:", dc.Foreground, " // ")
		fmt.Println("Background:", dc.Background)
	}
}

func CreateInputsForDeviceConfs() []files.InputData{
	return []files.InputData{
		files.InputData{
			Capacity:   7,
			Background: map[int]int{1:6, 2:2, 3:4},
			Foreground: map[int]int{1:2},
		},
		files.InputData{
			Capacity:   10,
			Background: map[int]int{1:5, 2:7, 3:10, 4:3},
			Foreground: map[int]int{1:5, 2:4, 3:3, 4:2},
		},
		files.InputData{
			Capacity:   20,
			Background: map[int]int{1:9, 2:15, 3:8},
			Foreground: map[int]int{1:11, 2:8, 3:12},
		},
		files.InputData{
			Capacity:   20,
			Background: map[int]int{1:7, 2:14, 3:8},
			Foreground: map[int]int{1:14, 2:5, 3:10},
		},
	}
}