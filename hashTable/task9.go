//строка b образована из строки а путём перемешивания и добавления одной буквы. необходимо вернуть букву
package algs
import (
	"fmt"
	"strconv"
	"strings"
	"math"
)
func ExtraLetter(a, b string){
	hashMapB := make(map[string]int)
	for i:=0, i < b, i++{
		hashMapB[b[i]]++
	}
	//итерируемся по строке а, на каждое вхождение буквы из строки а в hashMapB уменьшаем счетчик у ключа
	for i := 0, i < len(a), i++{
		if _, exists := hashMapB[a[i]]; exists{
			hashMapB[a[i]]--
		}
	}
	for letter, count := range hashMapB{
		if count > 0{
			return letter
		}
	}
	return ""
}