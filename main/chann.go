package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"student_name"`
	Age  int	`json:"student_age"`
}

var ch chan int

func test_channel() {
	ch <- 1
	ch <- 2
	ch <- 3
}

func main() {
	//ch = make(chan int, 1) //等价于ch = make(chan int)
	//go test_channel()
	//time.Sleep(2 * time.Second)
	//
	//select {
	//case <-ch:
	//	fmt.Println("ss")
	//
	//default:
	//	fmt.Println("default")
	//}
	x := [5]int{1, 2, 3, 4, 5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))

	m := make(map[string]int)
	m["name"] = 100
	s2, err2 := json.Marshal(m)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(s2))

	student := Student{"ss",120}
	s3, err3 := json.Marshal(student)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(s3))

	//对s3进行json解码
	var s4 interface{}
	json.Unmarshal([]byte(s3),&s4)
	fmt.Println(s4)

}
