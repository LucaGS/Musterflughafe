package Anlagen_Package

type Anlage struct {
	Name        string    `json: "name"`
	Anlagen_nr  string    `json:"anlagen_nr"`
	Bezeichnung string    `json:"bezeichnung"`
	Raum        string    `json:"raum"`
	Massnahmen  [2]string `json:"massnahmen"`
	Steuerung   string    `json: "steuerung"`
}

func NewAnlage(name string, anlagen_nr string, bezeichnung string, raum string, massnahmen [2]string, steuerung string) *Anlage {
	return &(Anlage{Name: name, Anlagen_nr: anlagen_nr, Bezeichnung: bezeichnung, Raum: raum, Massnahmen: massnahmen, Steuerung: steuerung})
}
