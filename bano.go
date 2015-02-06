package main

import (
	//"container/list"
	"encoding/csv"
	"flag"
	"fmt"
	core "github.com/goloc/goloc-core"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir := flag.String("dir", "", "directory")
	outputFile := flag.String("out", "", "output file")
	flag.Parse()
	if *dir == "" {
		fmt.Printf("Directory is mandatory\n")
	}
	if *outputFile == "" {
		fmt.Printf("Output file is mandatory\n")
	}
	if *dir == "" || *outputFile == "" {
		fmt.Printf("\nExecute help: bano -help\n")
		return
	}
	mi := core.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(*dir)
	mi.SaveInFile(*outputFile)
}

type Bano struct {
	core.Index
}

func (b *Bano) IndexDir(dirname string) {
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}
	for _, info := range infos {
		name := info.Name()
		if strings.HasSuffix(name, ".csv") {
			fmt.Printf(name + "\n")
			b.IndexFile(dirname + "/" + name)
		}
	}
}

func (b *Bano) IndexFile(filename string) {
	var loc core.Localisation
	var street *core.Street
	var zone *core.Zone
	var address *core.Address
	var addressId, num, streetName, postcode, city, lat, lon, streetId, zoneId string
	var floatLat, floatLon float64
	var records []string
	var i int

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Printf("[+")

	reader := csv.NewReader(file)
	reader.TrailingComma = true

	for {
		records, err = reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		addressId = records[0]
		num = records[1]
		streetName = records[2]
		postcode = records[3]
		city = records[4]
		lat = records[6]
		lon = records[7]
		streetId = addressId[:10]
		zoneId = addressId[:5]

		loc = b.Get(zoneId)
		if loc == nil {
			zone = core.NewZone()
			zone.Id = zoneId
			zone.Postcode = postcode
			zone.City = city
			b.Add(zone)
		} else {
			zone = loc.(*core.Zone)
		}

		loc = b.Get(streetId)
		if loc == nil {
			street = core.NewStreet()
			street.Id = streetId
			street.StreetName = streetName
			street.Zone = zone
			floatLat, err = strconv.ParseFloat(lat, 64)
			if err == nil {
				street.Lat = floatLat
			}
			floatLon, err = strconv.ParseFloat(lon, 64)
			if err == nil {
				street.Lon = floatLon
			}
			b.Add(street)
		} else {
			street = loc.(*core.Street)
		}

		address = core.NewAddress()
		address.Num = num
		floatLat, err = strconv.ParseFloat(lat, 64)
		if err == nil {
			address.Lat = floatLat
		}
		floatLon, err = strconv.ParseFloat(lon, 64)
		if err == nil {
			address.Lon = floatLon
		}

		// street.LinkedAddress = core.NewLinkedElement(address, street.LinkedAddress)

		i++
		if math.Mod(float64(i), 20000) == 0 {
			fmt.Printf("+")
		}
	}
	fmt.Printf("]\n")
}

func NewBano(index core.Index) *Bano {
	b := new(Bano)
	b.Index = index
	return b
}
