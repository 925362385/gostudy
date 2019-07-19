package main

import (
	"fmt"
	"regexp"
)

func main(){
	isok, _ := regexp.Match("[a-zA-z]{3}", []byte("zh1l"))
	is,_ := regexp.MatchString("[a-zA-z]{3}", "zhi")

	fmt.Printf("%v",isok)
	fmt.Printf("%v",is)

}
