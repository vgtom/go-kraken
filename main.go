package main

import (
	"log"
	"fmt"
	"bufio"
    "os"

	"github.com/aopoltorzhicky/go_kraken/rest"
)

func main() {
	api := rest.New("", "")

	t, err := api.Time()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(t)

	resp, err := api.AssetPairs()
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("kraken.dat")
    if err != nil{
		log.Fatal(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	for _, v := range resp {
		fmt.Println("Pair ", v.WSName)
		_, err := w.WriteString(v.WSName + "\n")
		if err != nil{
			log.Fatal(err)
		}
	}

	w.Flush()
}