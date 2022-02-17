package device

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateDeviceCombinations(t *testing.T) {
	iD := createInputsForDeviceConfs()
	deltaDeviceCapacity := 1
	for _,i := range iD{
		device, err := CreateDeviceCombinations(i.Capacity,deltaDeviceCapacity,i.Background,i.Foreground)
		assert.NoError(t,err, "There is a problem with the parse on the file that you provided")
		assert.Equal(t,i.Capacity,device.Capacity, "The capacity value from device should to be same thann the input")
		if device.DeviceConf == nil{
			assert.NotNil(t,device.NonExactDeviceConf,"There should to be possible combinations when at least there is 1 delta diference between the capacity and consumption")
			for _,dc := range *device.NonExactDeviceConf{
				assert.Equal(t,(i.Capacity-deltaDeviceCapacity),(dc.Foreground.Consumption+dc.Background.Consumption),fmt.Sprint("The sum betweenn foreground and background vs the capacity, the delta should be ",deltaDeviceCapacity))
			}
		} else{
			assert.NotNil(t,device.DeviceConf,"There should to be possible combinations when at least there is  delta diference between the capacity and consumption")
			for _,dc := range *device.DeviceConf{
				assert.Equal(t,i.Capacity,(dc.Foreground.Consumption+dc.Background.Consumption),fmt.Sprint("The sum betweenn foreground and background vs the capacity, the delta should be ",deltaDeviceCapacity))
			}
		}
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

type inputData struct {
	Capacity int
	Background map[int]int
	Foreground map[int]int
}

func createInputsForDeviceConfs() []inputData {
	return []inputData{
		inputData{
			Capacity:   7,
			Background: map[int]int{1:6, 2:2, 3:4},
			Foreground: map[int]int{1:2},
		},
		inputData{
			Capacity:   10,
			Background: map[int]int{1:5, 2:7, 3:10, 4:3},
			Foreground: map[int]int{1:5, 2:4, 3:3, 4:2},
		},
		inputData{
			Capacity:   20,
			Background: map[int]int{1:9, 2:15, 3:8},
			Foreground: map[int]int{1:11, 2:8, 3:12},
		},
		inputData{
			Capacity:   20,
			Background: map[int]int{1:7, 2:14, 3:8},
			Foreground: map[int]int{1:14, 2:5, 3:10},
		},
	}
}