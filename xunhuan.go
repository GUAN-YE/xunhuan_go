package main

import "fmt"

func main(){
	numbers :=[6]int{1,2,3,4,5}
	for i,x :=range numbers{
		fmt.Println("第%d位的值是：%d",i,x)
	}
}