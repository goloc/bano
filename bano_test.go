package main

import (
	"fmt"
	core "github.com/goloc/core"
	"testing"
)

func TestImport(t *testing.T) {
	mi := core.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(".")

	sizeLoc := bano.SizeLocalisation()
	fmt.Printf("size localisation %d\n", sizeLoc)
	if sizeLoc != 3572 {
		t.Fail()
	}

	sizeIndex := bano.SizeIndex()
	fmt.Printf("size index %d\n", sizeIndex)
	if sizeIndex != 900 {
		t.Fail()
	}

	mi.SaveInFile("golocTest.gob")
}

func TestReload(t *testing.T) {
	memindex := core.NewMemindexFromFile("golocTest.gob")

	sizeLoc := memindex.SizeLocalisation()
	fmt.Printf("size localisation %d\n", sizeLoc)
	if sizeLoc != 3572 {
		t.Fail()
	}

	sizeIndex := memindex.SizeIndex()
	fmt.Printf("size index %d\n", sizeIndex)
	if sizeIndex != 900 {
		t.Fail()
	}
}
