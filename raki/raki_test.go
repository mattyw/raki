package raki

import (
	"os"
	"reflect"
	"testing"
)

func TestCols(t *testing.T) {
	expected0 := col{"string", "Location", "Location"}
	expected1 := col{"number", "CO2/server hr", "MickeyGs"}
	expected := MastadonCRanking{Cols: []col{expected0, expected1}}

	file, err := os.Open("dc_rankings.json")
	if err != nil {
		t.Fatal(err)
	}
	jsonData := make([]byte, 1892)
	_, err = file.Read(jsonData)
	if err != nil {
		t.Fatal(err)
	}
	data := Parsejson(jsonData)
	if !reflect.DeepEqual(data.Cols, expected.Cols) {
		t.Errorf("%s != %s", data.Cols, expected.Cols)
	}
}

func TestFirstPlace(t *testing.T) {
	expected := "Iceland (Greenqloud)"

	file, err := os.Open("dc_rankings.json")
	if err != nil {
		t.Fatal(err)
	}
	jsonData := make([]byte, 1892)
	_, err = file.Read(jsonData)
	if err != nil {
		t.Fatal(err)
	}
	data := Parsejson(jsonData)
	firstPlace := data.Ranking[0].Name
	if !reflect.DeepEqual(firstPlace, expected) {
		t.Errorf("%s != %s", firstPlace, expected)
	}
}
