// строка b образована из строки а путём перемешивания и добавления одной буквы. необходимо вернуть букву
package algs

func ExtraLetter(a, b string) string {
	hashMapB := make(map[rune]int)
	for _, i := range b {
		hashMapB[i]++
	}
	//итерируемся по строке а, на каждое вхождение буквы из строки а в hashMapB уменьшаем счетчик у ключа
	for _, i := range a {
		hashMapB[i]--
	}

	for letter, count := range hashMapB {
		if count > 0 {
			return string(letter)
		}
	}
	return ""
}
