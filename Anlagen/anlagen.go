package Anlagen_Package

type Anlage struct {
	Name        string
	Anlagen_nr  string
	Bezeichnung string
	Raum        string
	Massnahmen  [2]string
	Steuerung   string
}

func NewAnlage(name string, anlagen_nr string, bezeichnung string, raum string, massnahmen [2]string, steuerung string) *Anlage {
	return &(Anlage{Name: name, Anlagen_nr: anlagen_nr, Bezeichnung: bezeichnung, Raum: raum, Massnahmen: massnahmen, Steuerung: steuerung})
}
