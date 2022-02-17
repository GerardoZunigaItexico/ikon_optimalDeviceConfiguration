package device

type Consumption struct {
	Id int
	Consumption int
}

type DeviceConf struct {
	Capacity int
	Background Consumption
	Foreground Consumption
}

func CreateDeviceCombinations(deviceCapacity int, background map[int]int, foreground map[int]int ) map[int][]DeviceConf{
	combinations := map[int][]DeviceConf{}
	for bK, bV := range background{
		for fK, fV := range foreground {
			if sum:= bV + fV; sum<=deviceCapacity{
				devSl,ok := combinations[sum]
				newDC := DeviceConf{ Capacity:sum,Background: Consumption{Id:bK,Consumption: bV}, Foreground: Consumption{Id:fK,Consumption: fV}}
				if !ok{
					devSl := []DeviceConf{newDC}
					combinations[sum]=devSl
				} else{
					devSl = append(devSl,newDC)
				}
			}
		}
	}

	return combinations
}

func appendToMap(combinations map[int]*[]*DeviceConf)  func (sum,bK, bV, fK, fV int){
	return func(sum,bK, bV, fK, fV int){
		devSl,ok := combinations[sum]
		newDC := &DeviceConf{ Capacity:sum,Background: Consumption{Id:bK,Consumption: bV}, Foreground: Consumption{Id:fK,Consumption: fV}}
		if !ok{
			devSl := &([]*DeviceConf{newDC})
			combinations[sum]=devSl
		} else{
			*devSl = append(*devSl,newDC)
		}
	}
}
