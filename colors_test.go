package colors

import "testing"

func Test_testy(t *testing.T) {
	phrase := "holala, chouette alors"
	print(Faint)
	print(phrase)
	println(Reset)

	Println("whatever dude", Red, BLUE)

	sentence := "holala, merde alors"

	print(Blue, Blink)
	defer Reseting()
	println(sentence)
	println(phrase)
}
