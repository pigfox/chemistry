package main

import "fmt"

type Atom struct {
	Name   string `json:name`
	Charge int    `json:charge`
}

type Atoms struct {
	Atom   Atom `json:atom`
	Number int  `json:number`
}

type Molecule struct {
	Atoms []Atoms `json:atoms`
}

func main() {
	fmt.Println("Chemistry")
	hydrogenAtom := Atom{Name: "Hydrogen", Charge: 1}
	clorineAtom := Atom{Name: "Chlorine", Charge: -1}
	oxygenAtom := Atom{Name: "Oxygen", Charge: -2}
	phosphorus := Atom{Name: "Phosphorus", Charge: -3}

	hydrogenAtoms := Atoms{Atom: hydrogenAtom, Number: 3}
	phosphorusAtoms := Atoms{Atom: phosphorus, Number: 1}
	oxygenAtoms := Atoms{Atom: oxygenAtom, Number: 4}

	//H3PO4
	phosphoricAcid := Molecule{}
	phosphoricAcid.Atoms = append(phosphoricAcid.Atoms, hydrogenAtoms)
	phosphoricAcid.Atoms = append(phosphoricAcid.Atoms, phosphorusAtoms)
	phosphoricAcid.Atoms = append(phosphoricAcid.Atoms, oxygenAtoms)

	if bond(hydrogenAtom, clorineAtom) {
		fmt.Println("Bond")
	} else {
		fmt.Println("No Bond")
	}

	if stable(phosphoricAcid) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not Stable")
	}
}

func bond(a1 Atom, a2 Atom) bool {
	if a1.Charge+a2.Charge == 0 {
		return true
	}
	return false
}

func stable(molecule Molecule) bool {
	charge := 0
	for c := range molecule.Atoms {
		charge += c
	}

	fmt.Println("Charge:", charge)

	if charge == 0 {
		return true
	}
	return false
}