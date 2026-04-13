package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . sample.txt result.txt")
	}
	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	data, err := os.ReadFile(sampleFile)
	if err != nil {
		fmt.Println("error", err)
	}

	text := string(data)
	result := CompileFunc(text)

	err = os.WriteFile(resultFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error writing file", err)
	}

	
}


func CompileFunc(text string) string{

	text = ModifyCases(text)
	text = ModifyArticle(text)
	text = ConvertToDecimal(text)
	text = FixPunctuation(text)
	text = FixQuotes(text)

	return text + "\n"
}