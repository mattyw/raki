package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type col struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Id    string `json:"id"`
}

type Location struct {
	V string `json:"v"`
	F string `json:"f"`
}

type CO2 struct {
	V float64 `json:"v"`
	F string  `json:"f"`
}

type Mixed struct {
	V float64 `json:"v"`
	N string  `json:"v"`
	F string  `json:"f"`
}

type Rank struct {
	Name     string
	Location string
	V        float64
	CO2hour  string
}

type RankingTable struct {
	Cols    []col
	Ranking []Rank
}

type MastadonCRanking struct {
	Cols []col                                 `json:"cols"`
	Rows []map[string][]map[string]interface{} `json:"rows"`
}

func RowToRanking(name, location, value, co2 interface{}) Rank {
	n, _ := name.(string)
	l, _ := location.(string)
	v, _ := value.(float64)
	c, _ := co2.(string)
	return Rank{n, l, v, c}
}

func Parsejson(jsonBlob []byte) RankingTable {
	var c MastadonCRanking
	result := RankingTable{}
	if err := json.Unmarshal(jsonBlob, &c); err != nil {
		log.Println(err)
	}
	result.Cols = c.Cols

	for _, v := range c.Rows {
		rank := RowToRanking(v["c"][0]["f"], v["c"][0]["v"],
			v["c"][1]["v"], v["c"][1]["f"])
		result.Ranking = append(result.Ranking, rank)
	}
	return result
}

func PrettyPrint(table RankingTable) string {
	buffer := bytes.Buffer{}
	for i := range table.Ranking {
		rank := table.Ranking[i]
		line := fmt.Sprintf("%s\t%s\t%02f\t%s\n", rank.Name, rank.Location, rank.V, rank.CO2hour)
		buffer.WriteString(line)
	}
	return buffer.String()

}
