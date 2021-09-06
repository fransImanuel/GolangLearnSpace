package main

import "fmt"

type struct2 struct {
	variabledidalamstruct1 string
}
func (s struct2) test() {
	fmt.Println("MASUK")
}

type struct1 struct {
	struct2
}


func main() {
	var contoh []struct1
	// &contoh = "Test"
	fmt.Println(contoh)
}