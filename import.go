package main

import (
	"container/list"
	"encoding/csv"
	"flag"
	"fmt"
	core "github.com/goloc/core"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type add func(street *core.Street)

func main() {
	dir := flag.String("dir", "", "directory")
	outpuFile := flag.String("out", "", "output file")
	flag.Parse()
	if *dir == "" {
		fmt.Printf("Directory is mandatory\n")
	}
	if *outpuFile == "" {
		fmt.Printf("Output file is mandatory\n")
	}
	if *dir == "" || *outpuFile == "" {
		fmt.Printf("\nExecute help: bano -help\n")
		return
	}
	memindex := core.NewMemindex()
	indexDir(*dir, memindex)
	memindex.SaveInFile(*outpuFile)
}

func indexDir(dirname string, index core.Index) {
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}
	for _, info := range infos {
		name := info.Name()
		if strings.HasSuffix(name, ".csv") {
			fmt.Printf(name + "\n")
			indexFile(dirname+"/"+name, index)
		}
	}
}

func indexFile(filename string, index core.Index) {
	i := 0
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Printf("[+")

	reader := csv.NewReader(file)
	reader.TrailingComma = true

	var lastZoneId string
	var lastZone *core.Zone
	var lastStreetId string
	var lastStreet *core.Street
	var currentAddresses *list.List

	for {
		records, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
		} else {
			addressId := records[0]
			num := records[1]
			streetName := records[2]
			postcode := records[3]
			city := records[4]
			lat := records[6]
			lon := records[7]
			streetId := strings.Split(addressId, "-")[0]
			zoneId := addressId[0:5]

			if lastZoneId == zoneId {

			} else {
				if lastZone != nil {
					// TODO
				}
				lastZone = core.NewZone()
				lastZone.Id = zoneId
				lastZone.Postcode = postcode
				lastZone.City = city
			}

			if lastStreetId == streetId {
				address := core.NewAddress()
				address.Num = num
				point := new(core.Point)
				floatLat, ok := strconv.ParseFloat(lat, 64)
				if ok == nil {
					point.Lat = floatLat
				}
				floatLon, ok := strconv.ParseFloat(lon, 64)
				if ok == nil {
					point.Lon = floatLon
				}
				address.Point = point
				if currentAddresses == nil {
					currentAddresses = list.New()
				}
				currentAddresses.PushBack(address)
			} else {
				if lastStreet != nil {
					if currentAddresses != nil {
						lastStreet.Addresses = make([]*core.Address, currentAddresses.Len())
						i := 0
						for e := currentAddresses.Front(); e != nil; e = e.Next() {
							address := e.Value.(*core.Address)
							lastStreet.Addresses[i] = address
							i++
						}
						currentAddresses = nil
					}
					index.Add(lastStreet)
				}
				lastStreet = core.NewStreet()
				lastStreet.Id = streetId
				lastStreet.StreetName = streetName
				lastStreet.Zone = lastZone
				point := new(core.Point)
				floatLat, ok := strconv.ParseFloat(lat, 64)
				if ok == nil {
					point.Lat = floatLat
				}
				floatLon, ok := strconv.ParseFloat(lon, 64)
				if ok == nil {
					point.Lon = floatLon
				}
				lastStreet.Point = point
			}

			lastStreetId = streetId

			i++
			if math.Mod(float64(i), 100000) == 0 {
				fmt.Printf("+")
			}
		}
	}
	if lastStreet != nil {
		index.Add(lastStreet)
	}
	fmt.Printf("+]\n")
}
