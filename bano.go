package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/goloc/goloc"
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
	mi := goloc.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(*dir)
	mi.SaveInFile(*outputFile)
}

type Bano struct {
	goloc.Index
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
	var loc goloc.Location
	var street *goloc.Street
	var zone *goloc.Zone
	var streetNumberedPoint *goloc.StreetNumberedPoint
	var float64Lat, float64Lon float64
	var records []string
	var i int

	badPos := float32(0)

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

		addressId := records[0]
		num := records[1]
		streetName := records[2]
		postcode := records[3]
		city := records[4]
		lat := records[6]
		lon := records[7]
		streetId := addressId[:10]
		zoneId := addressId[:5]

		float64Lat, err = strconv.ParseFloat(lat, 64)
		float32Lat := badPos
		if err == nil {
			float32Lat = float32(float64Lat)
		}
		float64Lon, err = strconv.ParseFloat(lon, 64)
		float32Lon := badPos
		if err == nil {
			float32Lon = float32(float64Lon)
		}
		if float32Lat == badPos && float32Lon == badPos {
			fmt.Printf("Bad position : %v", records)
		}

		loc = b.Get(zoneId)
		if loc == nil {
			zone = goloc.NewZone(zoneId)
			zone.Postcode = postcode
			zone.City = city
			b.Add(zone)
		} else {
			zone = loc.(*goloc.Zone)
		}
		if float32Lat != badPos || float32Lon != badPos {
			if float32Lat < zone.PointMin.Lat || zone.PointMin.Lat == badPos {
				zone.PointMin.Lat = float32Lat
			}
			if float32Lat > zone.PointMax.Lat || zone.PointMax.Lat == badPos {
				zone.PointMax.Lat = float32Lat
			}
			if float32Lon < zone.PointMin.Lon || zone.PointMin.Lon == badPos {
				zone.PointMin.Lon = float32Lon
			}
			if float32Lon > zone.PointMax.Lon || zone.PointMax.Lon == badPos {
				zone.PointMax.Lon = float32Lon
			}
		}

		loc = b.Get(streetId)
		if loc == nil {
			street = goloc.NewStreet(streetId)
			street.StreetName = streetName
			street.Zone = zone
			if float32Lat != badPos || float32Lon != badPos {
				street.Lat = float32Lat
				street.Lon = float32Lon
			}
			b.Add(street)
		} else {
			street = loc.(*goloc.Street)
		}

		streetNumberedPoint = goloc.NewStreetNumberedPoint(num)
		if float32Lat != badPos || float32Lon != badPos {
			streetNumberedPoint.Lat = float32Lat
			streetNumberedPoint.Lon = float32Lon
		}
		street.NumberedPoints[num] = streetNumberedPoint

		i++
		if math.Mod(float64(i), 20000) == 0 {
			fmt.Printf("+")
		}
	}
	fmt.Printf("]\n")
}

func NewBano(index goloc.Index) *Bano {
	b := new(Bano)
	b.Index = index
	b.AddStopWord("D", "DE", "DU", "DES", "L", "LE", "LA", "LES")
	b.AddStopWord("RUE", "ROUTE", "ALLEE", "PLACE", "CHEMIN", "IMPASSE", "AVENUE", "BOULEVARD")
	return b
}
