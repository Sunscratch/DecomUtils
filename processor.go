package main

import (
	"strings"
)

func deviceProcessor(cnmdbList []Cnmdb, mvtList []Mvt) []Mvt {
	for i, mvt := range mvtList {
		for _, cnmdb := range cnmdbList {
			if strings.Contains(strings.ToLower(cnmdb.fqdn), strings.ToLower(mvt.fqdn)) {
				mvtList[i].ip = cnmdb.ip
			}
		}
	}
	return mvtList
}
