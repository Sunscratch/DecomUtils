package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start processing...")
	rawMwt := readFromFile("mvt.csv", ';')
	mvtList := getMvtList(rawMwt)
	devicesToDecom := getDecomedDevices(mvtList)
	rawCnmdb := readFromFile("CNMDB.csv", ',')
	cnmdbList := getCnmdbList(rawCnmdb)
	newMvtList := deviceProcessor(cnmdbList, devicesToDecom)
	writeToFile(toCsvData(newMvtList), "export.csv", ';')
	fmt.Printf("Devices from cnmdb: %d\n", len(cnmdbList))
	fmt.Printf("Job completed, %d devices processed, %d devices to decom\n", len(mvtList), len(devicesToDecom))
}
