package main

import "fmt"

func main(){
	eeCount := 0
	for i:=1000; i <= 9999;i++{
		for j:=i;i <= 9999;j++{
			strMul := fmt.Sprintf("%d", i*j)
			if strMul[0] == strMul[len(strMul)-1]{
				eeCount += 1
			}
		}
	}
	fmt.Println(eeCount)
}
