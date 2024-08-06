package helloDeps8

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps8.PrintPhrase")
	return phrase
}
