package main

import (
	"github.com/goloc/goloc"
	"testing"
)

func TestImport(t *testing.T) {
	mi := goloc.NewMemindex()
	bano := NewBano(mi)
	bano.IndexDir(".")

	sizeLoc := bano.Index.(*goloc.Memindex).SizeLocation()
	if sizeLoc != 21734 {
		t.Fail()
	}

	sizeIndex := bano.Index.(*goloc.Memindex).SizeIndex()
	if sizeIndex != 1720 {
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
	if sizeIndex != 1720 {
		t.Fail()
	}
}
