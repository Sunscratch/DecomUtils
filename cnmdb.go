package main

type Cnmdb struct {
	did      string
	hostname string
	fqdn     string
	ip       string
	date     string
	online   string
	model    string
}

func mapCnmdb(entity []string) Cnmdb {
	result := Cnmdb{
		entity[0],
		entity[1],
		entity[2],
		entity[3],
		entity[4],
		entity[5],
		entity[6],
	}
	return result
}

func getCnmdbList(csvData [][]string) []Cnmdb {
	var result []Cnmdb
	for _, line := range csvData {
		result = append(result, mapCnmdb(line))
	}
	return result
}
