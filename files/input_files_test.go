package files

import (
	"fmt"
	"log"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_readfile(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines,err := readFile(path+"/testFiles/challenge.in")
	assert.NotEmpty(t,lines,"The file should to have at least 1 line")
	assert.NoError(t,err,"There are not expected errors on this unit case")
	assert.Equal(t,0,(len(lines)%3),"The file that will be read, should be defined in groups of 3 lines")
	for _,l := range lines{
		fmt.Println(l)
	}
}

func Test_parseInputLines(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines,err := readFile(path+"/testFiles/challenge.in")
	configurations := parseInputLines(lines)
	assert.NotEmpty(t,configurations,"At least one configuration")
	for _,c := range configurations{
		assert.Len(t,c,3,"Each Configuration should to have 3 lines at least")
		fmt.Println(c)
	}
}

func Test_regexForGetNumbers(t *testing.T){
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines,err := readFile(path+"/testFiles/challenge.in")
	configurations := parseInputLines(lines)
	inputDataSlice := regexForGetNumbers(configurations)
	for _, inputData := range inputDataSlice{
		fmt.Println( inputData.Capacity, "Background", inputData.Background, "Foreground", inputData.Foreground)
	}
}