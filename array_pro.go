package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	arr1[0]=1
	arr1[2]=10
	arr1[3] = 42
	fmt.Println("---Array print---")
	fmt.Println(arr1)
	fmt.Println("---Array length----")
	fmt.Println(len(arr1))
	names := []string{"laxmi","anjali","roopal","komal"}
	array_slice_compositeliteral(names)
	array_slice_range(names)
	array_slicing_slice(names)
	array_slice_append(names)
	array_slice_delete(names)
	array_make_function()
	multidimensional_arr()
	underlying_arr()
}

func array_slice_compositeliteral(names[]string){
	fmt.Println("---Slice array using Composite Literal--")
	fmt.Println(names)
}

func array_slice_range(names[]string){
	 fmt.Println("---Slice array using index--")
         for index := range names {
         fmt.Println(names[index]);
        }


	fmt.Println("---Slice array using Range--")
        for index,value := range names {
         fmt.Println(index , value);
        }

	fmt.Println("---Slice array using Range optional index--")
        for _,value := range names {
	 fmt.Println(value);
	}
}


func array_slicing_slice(names[]string){
	fmt.Println("---Slicing a Slice using colon----")
	fmt.Println("---Slice first element----")
	fmt.Println(names[1:])
	fmt.Println("---Slice first and thrid element----")
	fmt.Println(names[1:3])
	fmt.Println("---Slice first,second,third element----")
        fmt.Println(names[2:3])
}


func array_slice_append(names[]string){
	fmt.Println("---append slice to a slice---")
	names = append(names,"first","second","third","fourth")
	fmt.Println(names)

        fmt.Println("---create new slice using composite literal----" )
	new_names := []string{"fifth","sixth","seventh","eight","ninth"}

	fmt.Println("----append new array in existing array----")
	names=append(names,new_names...)
	fmt.Println(names)

}

func array_slice_delete(names[]string){
       fmt.Println("---delete element from Slice ----")
       names = append(names[2:],names[1:]...)
       fmt.Println(names)
}


func array_make_function(){
    fmt.Println("----array using make func----")
    arr1 := make([]int, 10, 12)
    arr1 =append(arr1,11,12)
    fmt.Println(cap(arr1),"---its capacity is 12 if we try to slice more then 12 we get error panic: runtime error: slice bounds out of range")
    arr1=arr1[:12]
    fmt.Println(arr1)
}


func multidimensional_arr(){
    arr1 :=[]string{"laxmi","roopal","anjali"}
    arr2 :=[]string{"sharma","aggarwal","singh"}

    arr3 :=[][]string{arr1,arr2}

    fmt.Println("---print multidimensional arr---")
    fmt.Println(arr3)
}


func underlying_arr(){
     fmt.Println("----underlying arr---");
     arr1 := []int{1,2,3,4,5,6,7,8,9,10}
     fmt.Println(arr1)
     arr2 := append(arr1[:2],arr1[5:]...)
     fmt.Println(arr1)
     fmt.Println(arr2)

     fmt.Println("----make arr---");
     arr3 := make([]int,5,10)
     fmt.Println(cap(arr3))
     arr3 =append(arr3[:2],arr3[:10]...)
     fmt.Println(arr3)
}




