package main

import (
	"fmt"
	core "github.com/goloc/core"
	"testing"
)

func TestImport(t *testing.T) {
	memindex := core.NewMemindex()
	indexDir(".", memindex)
	memindex.SaveInFile("golocTest.gob")

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
