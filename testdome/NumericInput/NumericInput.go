

package main

import "fmt"


type UserInput interface {
   Add(rune)
   GetValue() string
}

type NumericInput struct {
   input string
}

func main() {
    original := "10a"
    

    
    r := []rune(original[:2])
    
    
    
    
    result := string(r)
    fmt.Println(result)
}
