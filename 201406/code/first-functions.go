package main

import "fmt"
// START OMIT
type op struct {
    fn   func(...bool) bool
}
type boolFunc func(...bool) bool
func main() {
	var fns = map[string]boolFunc{
 	   "and": func(x ...bool) bool { 
			for _, res := range x {
				if res == false { return false }
			}
			return true 
		},
		"or": func(x ...bool) bool { 
			for _, res := range x {
				if res == true { return true }
			}
			return false 
		},
	}
	for fn, result := range fns {
		fmt.Printf("Logic: \"%v\", Result: %v \n", fn, result(false, true, false))
	}
}
// END OMIT