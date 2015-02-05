package main

import (
	core "github.com/goloc/core"
	"testing"
)

func TestImport(t *testing.T) {
	mi := core.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(".")

	sizeLoc := bano.SizeLocalisation()
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
	memindex := core.NewMemindexFromFile("golocTest.gob")

	sizeLoc := memindex.SizeLocalisation()
	if sizeLoc != 21734 {
		t.Fail()
	}

	sizeIndex := memindex.SizeIndex()
	if sizeIndex != 9356 {
		t.Fail()
	}
}
