package main

import (
	"log"
	"os"

	//Anlagen_Package "v1/Anlagen"

	"fmt"
)

func main() {
	fmt.Println("Hello World")
	//AlleAnlagen := make([]Anlagen_Package.Anlage,0)
	exslFileNames := make([]string, 0)
	dateien, err := os.ReadDir("C:/Unterstufe/Programmierprojete/Fork/Go/Musterflughafen/exslfiles")
	if err != nil {
		log.Fatal(err)
	}
	for _, datei := range dateien {
		fmt.Println(datei.Name())
		exslFileNames = append(exslFileNames, datei.Name())
	}

}
