package main

import (
	"testing"

	"github.com/goloc/goloc"
)

func TestImport(t *testing.T) {
	mi := goloc.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(".")

	sizeLoc := len(mi.Locations)
	if sizeLoc != 21734 {
		t.Fail()
	}

	sizeIndex := len(mi.Keys)
	if sizeIndex != 1720 {
		t.Fail()
	}

	sizeStopWords := mi.GetStopWords().Size()
	if sizeStopWords != 37 {
		t.Fail()
	}

	sizeEncodedStopWords := mi.GetEncodedStopWords().Size()
	if sizeEncodedStopWords != 25 {
		t.Fail()
	}

	mi.SaveInFile("golocTest.gob")
}

func TestReload(t *testing.T) {
	mi := goloc.NewMemindexFromFile("golocTest.gob")

	sizeLoc := len(mi.Locations)
	if sizeLoc != 21734 {
		t.Fail()
	}

	sizeIndex := len(mi.Keys)
	if sizeIndex != 1720 {
		t.Fail()
	}

	sizeStopWords := mi.GetStopWords().Size()
	if sizeStopWords != 37 {
		t.Fail()
	}

	sizeEncodedStopWords := mi.GetEncodedStopWords().Size()
	if sizeEncodedStopWords != 25 {
		t.Fail()
	}
}
