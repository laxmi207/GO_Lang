package main

import "fmt"

func main() {
	map_arr := map[string]int{
		"apple":           100,
		"Grapes": 120,
	}
	fmt.Println(map_arr)

	index := map_arr["Grapes"]
	fmt.Println(index)

	fmt.Println("----if key exists in map---")
	if value,ok :=map_arr["apple"]; ok{
	  fmt.Println(value)
	}

	fmt.Println("---if key not exists in map---")

	if value,ok :=map_arr["aaa"]; ok{
	  fmt.Println(value)
	}

        fmt.Println("----Add element in map---")
        map_arr["banana"]=50
	fmt.Println(map_arr)

	fmt.Println("---Delete element in map----")
	delete(map_arr,"banana")
	fmt.Println(map_arr)
}
