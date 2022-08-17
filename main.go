package main

import (
	"fmt"

	"github.com/maeldevfr/gomod/footypes"
)

func main() {
	var fooVar footypes.Types
	fooVar.TypeBoolean = true
	fooVar.TypeInt = 14
	fooVar.TypeString = "Hello !"
	fooVar.TypeFloat32 = 9.999
	fooVar.TypeFloat64 = 99.9999999999
	fooVar.TypeFloat64 = 2.456787
	fooVar.TypeUInt = 7
	fooVar.TypeUInt8 = 255
	fooVar.TypeUInt16 = 65535
	fooVar.TypeUInt32 = 3999999999
	fooVar.TypeUInt64 = 10999999999999999999
	fooVar.TypeInt16 = 29999
	fooVar.TypeInt32 = 1999999999
	fooVar.TypeInt64 = 8999999999999999999

	var fooBeasts footypes.Beasts
	fooBeasts.Legs = 4
	fooBeasts.Beasts = "Cat"
	fooBeasts.Conclusion = "'This Beasts is a 'Cat', and he have 4 breeds !'"
	fooBeasts.Venimous = false

	var fooComputer footypes.Computer
	fooComputer.ArchData = 64
	fooComputer.Os = "MacOs X / darwin"
	fooComputer.Azerty = true
	fooComputer.Conclusion = "'Ordinateur Apple de 64Bits en mode azerty'"

	var fooLand footypes.Land
	fooLand.NameOfLand = "France"
	fooLand.LandOfUE = true
	fooLand.LandPopulation = 70000000
	fooLand.Conclusion = "'The French is good'"

	fmt.Printf("————————————————————\n")
	fmt.Println(fooVar)
	fmt.Printf("————————————————————\n")
	fmt.Println(fooBeasts)
	fmt.Printf("————————————————————\n")
	fmt.Println(fooComputer)
	fmt.Printf("————————————————————\n")
	fmt.Println(fooLand)
	fmt.Printf("————————————————————\n")
}
