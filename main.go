package main

import (
	"encoding/json"
	"os"
	Anlagen_Package "v1/Anlagen"

	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	AlleDaten := make([]Anlagen_Package.Anlage, 0)
	exslFileNames := GetDir("C:/Unterstufe/Programmierprojete/Fork/Go/Musterflughafen/exslfiles")
	for _, file := range exslFileNames {
		fmt.Println(file)
		f, err := excelize.OpenFile("C:/Unterstufe/Programmierprojete/Fork/Go/Musterflughafen/exslfiles/" + file)
		if err != nil {
			fmt.Println(err)

		}

		sheets := f.GetSheetList()
		//Alle dateien haben in diesem fall nur 1 Sheet
		sheet := sheets[0]
		fmt.Println(sheets[0])
		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)

		}
		for rowIndex, row := range rows {
			if rowIndex < 9 {
				continue
			}
			fmt.Printf("Current row: %v \n", rowIndex)
			for fieldIndex, field := range row {
				if fieldIndex < len(row) {
					fmt.Printf("Row: %v , field: %v value: %v \n", rowIndex, fieldIndex, field)
				}
			}
			if len(row) < 6 {
				continue
			}
			var massnahmen [2]string
			massnahmen[0] = row[4]
			massnahmen[1] = row[5]
			AlleDaten = append(AlleDaten, *&Anlagen_Package.Anlage{Name: row[0], Anlagen_nr: row[1], Bezeichnung: row[2], Raum: row[3], Massnahmen: massnahmen, Steuerung: row[6]})
		}
	}
	for _, v := range AlleDaten {
		fmt.Println(v)
	}
	datei, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer datei.Close()
	encoder := json.NewEncoder(datei)
	err = encoder.Encode(AlleDaten)
	encoder.SetIndent("", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("a total of %v rows", len(AlleDaten))
}
