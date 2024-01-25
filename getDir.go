package main

import (
	"fmt"
	"log"
	"os"
)

func GetDir(path string) (fileNames []string) {
	//C:/Unterstufe/Programmierprojete/Fork/Go/Musterflughafen/exslfiles
	dateien, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, datei := range dateien {
		fmt.Println(datei.Name())
		fileNames = append(fileNames, datei.Name())
	}
	return
}
