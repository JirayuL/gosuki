package main

import "fmt"

func init() {
	println("init func()")
}

func main() {
	// fmt.Println("Hi!")
	// var name string
	// fmt.Print("Input your name: ")
	// fmt.Scanf("%s", &name)
	// age := 21
	// fmt.Printf("Hello %s. I'm %d years old.\n", name, age)

	// var j int64
	// var na string

	// for loop
	// for i := 0; i < age; i++ {
	// 	fmt.Println(i)
	// }

	// array (cant change value)
	// var arr []string{
	// 	"a",
	// 	"b"
	// }

	// slice (can change value)
	// var slice make([]string)
	// slice := make([]string)
	// slice := []string{}

	// slice = append(slice, "aa")

	// map key:string value:string
	// mapVal := map[string]string{
	// 	"key": "value",
	// }

	// map key:string value:any
	// mapVal := map[string]interface{}{
	// 	"key": "value",
	// 	"number": 5,
	// }

	// mapVal := map[string]interface{}{
	// 	"key": "value",
	// 	"number": 5,
	// 	"nestedMap": map[string]interface{}{
	// 		"childKey": "childVal",
	// 	},
	// }
	// fmt.Printf("%+v", mapVal)

	// Hello("Gear", "GG", 21)
	// Prinln(0, "a", "b", "c")

	s, s2 := Hell()
	fmt.Println(s + s2)
}

// Hell func
func Hell() (string, string) {
	return "Wow", "Voooo"
}

// CoditionFunc func
func CoditionFunc(i int) int {
	if b := i + 5; b < 5 {
		fmt.Println(false)
		return 0
	}
	return 1
}

// Hello func
func Hello(name, surname string, age int) {
	fmt.Println("Hello", name, surname, age)
}

// Prinln func
func Prinln(num int, a ...string) {
	print(num)
	for _, e := range a {
		print(e)
	}
}
