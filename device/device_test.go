package device

import (
	"fmt"
	"testing"
)

func Test_CreateDeviceCombinations(t *testing.T) {
	iD := CreateInputs()
	for _,i := range iD{
		exactCombs, nonExactCombs := CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		fmt.Println("Exact: ",exactCombs)
		fmt.Println("Non Exact: ",nonExactCombs)
		fmt.Println("--------------")
	}
}

type InputData struct {
	Capacity int
	Background map[int]int
	Foreground map[int]int
}

func CreateInputs() []InputData{
	return []InputData{
		//InputData{
		//	Capacity:   7,
		//	Background: map[int]int{1:6, 2:2, 3:4},
		//	Foreground: map[int]int{1:2},
		//},
		InputData{
			Capacity:   10,
			Background: map[int]int{1:5, 2:7, 3:10, 4:3},
			Foreground: map[int]int{1:5, 2:4, 3:3, 4:2},
		},
		InputData{
			Capacity:   20,
			Background: map[int]int{1:9, 2:15, 3:8},
			Foreground: map[int]int{1:11, 2:8, 3:12},
		},
		InputData{
			Capacity:   20,
			Background: map[int]int{1:7, 2:14, 3:8},
			Foreground: map[int]int{1:14, 2:5, 3:10},
		},
	}
}