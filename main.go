package main

import (
	"fmt"
)

type Centimeters float64
type Kilograms float64

type Classification string

const Standard Classification = "STANDARD"
const Special Classification = "SPECIAL"
const Rejected Classification = "REJECTED"

const BulkyVolumeThreshold = 1_000_000 // cm^3
const BulkySpatialThreshold = 150      // cm
const HeavyWeightThreshold = 20        // kg

type Package struct {
	Width  Centimeters
	Height Centimeters
	Length Centimeters
	Mass   Kilograms
}

func (p Package) IsBulky() bool {
	if (p.Width * p.Height * p.Length) >= BulkyVolumeThreshold {
		return true
	}
	return p.Width >= BulkySpatialThreshold || p.Height >= BulkySpatialThreshold || p.Length >= BulkySpatialThreshold
}

func (p Package) IsHeavy() bool {
	return p.Mass >= HeavyWeightThreshold
}

func (p Package) Sort() Classification {
	isBulky := p.IsBulky()
	isHeavy := p.IsHeavy()
	//fmt.Printf("bulky? %v, heavy? %v\n", isBulky, isHeavy)
	if isBulky && isHeavy {
		return Rejected
	}
	if !isBulky && !isHeavy {
		return Standard
	}
	return Special
}

func Sort(width Centimeters, height Centimeters, length Centimeters, mass Kilograms) Classification {
	p := Package{
		Width:  width,
		Height: height,
		Length: length,
		Mass:   mass,
	}
	return p.Sort()
}

func main() {
	fmt.Println(Sort(1, 2, 3, 4))
	fmt.Println(Sort(1, 1, 1, HeavyWeightThreshold))
	fmt.Println(Sort(1000, 1000, 1000, 1))
	fmt.Println(Sort(1000, 1000, 1000, HeavyWeightThreshold))
}
