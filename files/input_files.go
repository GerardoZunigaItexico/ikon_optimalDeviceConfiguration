package files

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InputData struct {
	Capacity int
	Background map[int]int
	Foreground map[int]int
}

func readFile(inputFile string) (lines []string, err error) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if l := scanner.Text(); len(l)>0{
			lines = append(lines, l)
		}
	}

	f.Close()
	if (len(lines)%3)!=0{
		err = errors.New("The file that will be read, should be defined in groups of 3 lines")
	}

	return
}

func parseInputLines(lines []string) (configurations [][3]string){
	for i:=0; i<len(lines);i+=3{
		capacity:= lines[i]
		foreground:= lines[i+1]
		background :=lines[i+2]
		configurations = append(configurations, [3]string{capacity,foreground,background})
	}
	return
}

func regexForGetNumbers(configurations [][3]string) (inputData []InputData){
	for _, line := range configurations{
		capacity, err := strconv.Atoi(line[0])
		if err != nil{
			fmt.Println("there should to be only numbers")
			continue
		}
		foreground := createConfigurationMap(line[1])
		background := createConfigurationMap(line[2])
		inputData = append(inputData,InputData{
			Capacity:   capacity,
			Background: background,
			Foreground: foreground,
		})
	}
	return
}

func createConfigurationMap(line string) (m map[int]int){
	m = map[int]int{}
	pattern, err := regexp.Compile(`([0-9a-f]+), ([0-9a-f]+)`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	submatchall := pattern.FindAllString(line, -1)
	for _, subMatch := range submatchall{
		conf := strings.Split(subMatch, ", ")
		fId, err := strconv.Atoi(conf[0])
		if err != nil{
			fmt.Println("there should to be only numbers")
			return nil
		}
		fC, err := strconv.Atoi(conf[1])
		if err != nil{
			fmt.Println("there should to be only numbers")
			return nil
		}
		m[fId]=fC
	}
	return
}

func ParseChallengeIn(inputFile string) (id []InputData, err error) {
	lines,err := readFile(inputFile)
	if err != nil{
		return id, err
	}
	configurations := parseInputLines(lines)
	return regexForGetNumbers(configurations), nil
}