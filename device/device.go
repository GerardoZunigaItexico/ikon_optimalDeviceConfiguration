package device

type Consumption struct {
	Id          int
	Consumption int
}

type DeviceConf struct {
	Background Consumption
	Foreground Consumption
}

type Device struct {
	Capacity   int
	DeviceConf []DeviceConf
}

func CreateDeviceCombinations(deviceCapacity, deltaDevCap int, background map[int]int, foreground map[int]int) (*Device, *Device) {
	exactCombs    := map[int]*[]DeviceConf{}
	nonExactCombs := map[int]*[]DeviceConf{}
	for bK, bV := range background {
		for fK, fV := range foreground {
			if sum := bV + fV; sum == deviceCapacity {
				appendToMap(exactCombs,sum, bK, bV, fK, fV)
			} else if (deviceCapacity>sum) && ((deviceCapacity-sum)<=deltaDevCap){
				appendToMap(nonExactCombs,sum, bK, bV, fK, fV)
			}
		}
	}

	if len(exactCombs)>0 {return &Device{deviceCapacity,mapToSlice(exactCombs)},nil}
	return nil, &Device{deviceCapacity,mapToSlice(nonExactCombs)}
}

func appendToMap(combinations map[int]*[]DeviceConf, sum, bK, bV, fK, fV int) {
	devSl, ok := combinations[sum]
	newDC := DeviceConf{Background: Consumption{Id: bK, Consumption: bV}, Foreground: Consumption{Id: fK, Consumption: fV}}
	if !ok {
		devSl := &([]DeviceConf{newDC})
		combinations[sum] = devSl
	} else {
		*devSl = append(*devSl, newDC)
	}
}

func mapToSlice(combinations map[int]*[]DeviceConf)(deviceConfigs []DeviceConf){
	for _, slice := range combinations {
		for _,vv := range *slice {
			deviceConfigs = append(deviceConfigs, vv)
		}
	}
	return 
}
