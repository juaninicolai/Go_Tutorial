package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var score = 99.5

func main() {

	//Go does not allow you to declare strings variables and not used them
	//4 ways to declare variables
	var nameOne string = "Juan"
	var nameTwo = "Coti"
	var nameThree string
	//can't use this way outside of a function
	nameFour := "Karo"

	//fmt.Println(nameOne, nameTwo, nameThree)
	nameTwo = "Snorri"
	nameThree = "Andre"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int = 25
	//can specify the number of bit per itn variable
	var numTwo int8 = -128
	//uint doesn't allow negatives, uint can also get size declaration fx uint8
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	//float ==> need to declare the amount of bits
	var scoreOne float32 = 25.98
	//float64 has a bigger precision and it doesn't take that much memory
	var scoreTwo float64 = 54841546315
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	fmt.Print("hello, ")
	fmt.Print("world! \n")
	//unlike Java, the Println command adds a space by default when a comma is typed
	fmt.Println("My name is", nameOne, "and my age is", ageOne)

	//Printf (formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", ageTwo, nameTwo)
	fmt.Printf("My age is %q and my name is %q \n", ageTwo, nameTwo)
	fmt.Printf("Age is of type %T \n", ageOne)
	fmt.Printf("You scored %0.1f points \n", 255.55)

	//Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", ageTwo, nameTwo)
	fmt.Println("the saved string is: ", str)

	//arrays (like Java Arrays)
	var ages [3]int = [3]int{20, 25, 30}
	var agesShort = [3]int{20, 25, 30}

	names := []string{"mario", "yoshi", "peach", "bouza"}
	fmt.Println(ages, len(ages), agesShort)
	fmt.Println(names, len(names))

	//slices (like Java ArrayLists)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	//Standard Library
	greeting := "hello there friends!"
	//returns boolean
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	//return the index of the given string
	fmt.Println(strings.Index(greeting, "ll"))
	//splits the string into arrays following the provided char / string
	fmt.Println(strings.Split(greeting, " "))

	ages2 := []int{45, 20, 35, 30, 75, 60, 50, 25}
	//sort INT in ascending order within a slice
	sort.Ints(ages2)
	fmt.Println(ages2)

	//find the index of an element in a slice
	index := sort.SearchInts(ages2, 30)
	fmt.Println(index)

	//sort strings in alphabethical order
	sort.Strings(names)
	fmt.Println(names)
	//search and return the index of an element in a slice
	fmt.Println(sort.SearchStrings(names, "bouza"))

	x := 0
	for x < 5 {
		fmt.Println("value of x", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i", i)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	//conditionals
	age := 45
	fmt.Println((age <= 50))
	fmt.Println((age >= 50))
	fmt.Println((age == 50))
	fmt.Println((age != 50))

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

	sayGreeting("Juani")
	sayBye("Coti")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3)

	sayHello("Juan")

	for _, v := range points {
		fmt.Println(v)
	}
	showScore()

	//Maps
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		123456789: "mario",
		987654321: "luigui",
		654123789: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456789])

	//update an item in the map
	phonebook[123456789] = "bowser"
	fmt.Println(phonebook)
}

func sayHello(s string) {
	panic("unimplemented")
}

func showScore() {
	panic("unimplemented")
}

//functions
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

//method signature for a return method
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

//multiple return values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names2 := strings.Split(s, " ")

	var initials []string
	for _, v := range names2 {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
