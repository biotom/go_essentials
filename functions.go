//package Exercise_Files
package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error){
	resp, err := http.Get(url)
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	contType := resp.Header.Get("Content-Type")
	if contType == ""{
		return "", fmt.Errorf("no content-type header")
	}
	return contType,nil
}
func main(){
	resp, err := contentType("http://www.google.com")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(resp)
	}

}