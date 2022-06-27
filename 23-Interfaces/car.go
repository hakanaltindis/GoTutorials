package main

import "fmt"

func CarExecute(i Carface) {
	fmt.Println(i.Information())
}

func main() {

	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Price = 1.4
	ferr.Special = true

	merco := new(Mercedes)
	merco.Brand = "Mercedes"
	merco.Model = "S500"
	merco.Color = "Black"
	merco.Price = 3.4

	CarExecute(ferr)
	CarExecute(merco)
}

// Interface
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base Struct
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

// Ferrari
type Ferrari struct {
	Car // Composition
	SpecialProduction
}

func (f Ferrari) Run() bool {
	return true
}

func (f Ferrari) Stop() bool {
	return true
}

func (f *Ferrari) Information() string {
	return f.Brand + " " + f.Model
}

// Mercedes
type Mercedes struct {
	Car // Composition
}

func (f Mercedes) Run() bool {
	return true
}

func (f Mercedes) Stop() bool {
	return true
}

func (m *Mercedes) Information() string {
	return m.Brand + " " + m.Model
}
