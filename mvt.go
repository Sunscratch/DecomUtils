package main

type Mvt struct {
	fqdn     string
	ip       string
	activity string
	ticket   string
	ping     string
	dns      string
}

func mapMvt(entity []string) Mvt {
	mvt := Mvt{
		entity[0],
		entity[1],
		entity[2],
		entity[3],
		entity[4],
		entity[5],
	}
	return mvt
}

func (mvt *Mvt) toDecom() bool {
	return mvt.ping == "Fail" && mvt.dns == "Fail" && mvt.ip == ""
}

func getDecomedDevices(mvtList []Mvt) []Mvt {
	var result []Mvt
	for _, device := range mvtList {
		if device.toDecom() {
			result = append(result, device)
		}
	}
	return result
}

func (mvt *Mvt) unMap() []string {
	var result []string
	result = append(result, string(mvt.fqdn))
	result = append(result, string(mvt.ip))
	result = append(result, string(mvt.activity))
	result = append(result, string(mvt.ticket))
	result = append(result, string(mvt.ping))
	result = append(result, string(mvt.dns))
	return result
}

func getMvtList(csvData [][]string) []Mvt {
	var result []Mvt
	for _, line := range csvData {
		result = append(result, mapMvt(line))
	}
	return result
}

func toCsvData(mvtList []Mvt) [][]string {
	var result [][]string
	for _, device := range mvtList {
		result = append(result, device.unMap())
	}
	return result
}
