package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")
	fixeWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixeWords)
}
