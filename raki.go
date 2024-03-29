package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"raki/query"
	"text/tabwriter"
)

func main() {
	resp, err := http.Get("https://www.mastodonc.com/dc_rankings")
	if err != nil {
		log.Fatal(err)
	}
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := query.Parsejson(json)
	str := query.PrettyPrint(data)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprint(w, str)
	w.Flush()
}
