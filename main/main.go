package main

import "fmt"

/**
* make s := make([]int,10,20)
* 10 为length
* 20 为内容中的容量
*
* map map[keyType]valueType
* ss := make(map[string]int)
**/
func ss() {
	xx := [10]int{1, 2, 2, 3, 4, 2, 5, 5, 6}

	y := xx[1:3]
	y = append(y,5,6,7)

	//yy := make([]int, 3, 5)
	sum:=0

	fmt.Printf("%v", len(y))
	fmt.Println()
	fmt.Printf("%v", cap(y))
	fmt.Println()
	var x int
	for x = 1; x < len(xx); x++ {
		fmt.Println(xx[x])
		sum +=xx[x]
	}
	fmt.Println(sum)

	mapss := make(map[string]int)
	mapss["ss"] = 70
	for i,v:=range mapss {
		fmt.Println(i,v)
	}
	fmt.Println(swap(1,2))


}


func swap(a int,b int)(int,int){
	return b,a
}