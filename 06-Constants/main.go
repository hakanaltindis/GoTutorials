package main

import "fmt"

// 1. define constant
const thy string = "Turkish Airlines"

// 2. define constant
const tai = "TAI"

/*
 * You can use constants as enum.
 * See the below to details
 */
type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE    Brand = "Google"
	DT        Brand = "Dogus Teknoloji"
)

type Season int

// Define season constants with iota
// To prevent to assign same value different constants, you can use iota
const (
	SUMMER Season = iota // 0 because of iota
	AUTUMN               // 1
	WINTER               // 2
	SPRING               // 3
)

func main() {

	// print constant variable to console
	fmt.Println(thy, tai)

	// print constant as enum to console
	PrintBrand(DT)
	PrintBrand(MICROSOFT)

	// print constant as season enum to console
	PrintSeason(WINTER)
}

func PrintBrand(brand Brand) {
	fmt.Println(brand)
}

func PrintSeason(season Season) {
	fmt.Println("Season is ", season)
}
