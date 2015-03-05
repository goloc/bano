package main

import (
	"github.com/goloc/goloc"
	"testing"
)

func TestImport(t *testing.T) {
	mi := goloc.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(".")

	sizeLoc := bano.SizeLocation()
	if sizeLoc != 21734 {
		t.Fail()
	}

	sizeIndex := bano.SizeIndex()
	if sizeIndex != 9356 {
		t.Fail()
	}

	mi.SaveInFile("golocTest.gob")
}

func TestReload(t *testing.T) {
	memindex := goloc.NewMemindexFromFile("golocTest.gob")

	sizeLoc := memindex.SizeLocation()
	if sizeLoc != 21734 {
		t.Fail()
	}

	sizeIndex := memindex.SizeIndex()
	if sizeIndex != 9356 {
		t.Fail()
	}
}
