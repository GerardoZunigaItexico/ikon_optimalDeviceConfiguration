package device

type Consumption struct {
	Id          int
	Consumption int
}

type DeviceConf struct {
	Capacity   int
	Background Consumption
	Foreground Consumption
}

func CreateDeviceCombinations(deviceCapacity, deltaDevCap int, background map[int]int, foreground map[int]int) ([]DeviceConf, []DeviceConf) {
	exactCombs := map[int]*[]DeviceConf{}
	nonExactCombs := map[int]*[]DeviceConf{}
	for bK, bV := range background {
		for fK, fV := range foreground {
			if sum := bV + fV; sum == deviceCapacity {
				appendToMap(exactCombs,sum, bK, bV, fK, fV)
			} else if (deviceCapacity>sum) && ((deviceCapacity-sum)<=deltaDevCap){
				appendToMap(&nonExactCombs,sum, bK, bV, fK, fV)
			}
		}
	}

	if len(exactCombs)>0 {return mapToSlice(exactCombs),nil}
	return nil, mapToSlice(nonExactCombs)
}

func appendToMap(combinations map[int]*[]DeviceConf, sum, bK, bV, fK, fV int) {
	devSl, ok := combinations[sum]
	newDC := DeviceConf{Capacity: sum, Background: Consumption{Id: bK, Consumption: bV}, Foreground: Consumption{Id: fK, Consumption: fV}}
	if !ok {
		devSl := &([]DeviceConf{newDC})
		combinations[sum] = *devSl
	} else {
		*devSl = append(*devSl, newDC)
	}
}

func mapToSlice(combinations map[int][]DeviceConf)(deviceConfigs []DeviceConf){
	for _,v := range combinations {
		deviceConfigs = append(deviceConfigs,v...)
	}
	return 
}
