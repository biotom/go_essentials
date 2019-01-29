package main
import(
	"fmt"
	"strings"
)
func main(){
	//wordCount := make(map[string]int)
	text := "Needles and pins " +
		"Needles and pins " +
		"Sew me a sail " +
		"To catch me the wind"
	textFields := strings.Fields(text)
	wordCount := map[string]int{}

	for _, value := range(textFields){
		wordCount[strings.ToLower(value)]++
		/*
		original answer

		lowerWord := strings.ToLower(value)
		i, ok := wordCount[lowerWord]
		if !ok{
			wordCount[lowerWord] = 1
		}else{
			wordCount[lowerWord] = i + 1
		}
		*/
	}
	for key,value := range wordCount{
		fmt.Println("Key:", key, "Value:", value)
	}


	}

