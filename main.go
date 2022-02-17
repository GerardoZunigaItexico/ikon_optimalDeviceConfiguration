package main

import (
	"fmt"
	"ikon_optimalDeviceConfiguration/device"
	"ikon_optimalDeviceConfiguration/files"
	"log"
	"os"
)

func main(){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	inputData := files.ParseChallengeIn(path+"/files/testFiles/challenge.in")
	for _,i := range inputData{
		exactCombs, nonExactCombs := device.CreateDeviceCombinations(i.Capacity,1,i.Background,i.Foreground)
		fmt.Println("Exact: ",exactCombs)
		fmt.Println("Non Exact: ",nonExactCombs)
		fmt.Println("--------------")
	}
}
