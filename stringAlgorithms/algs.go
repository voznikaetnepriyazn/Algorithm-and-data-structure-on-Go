package algs

//метод прямого перебора

//алгоритм бойера - мура- хорспула
//начать сравнение с последнего элемента подстроки и соответствующего по номеру элемента строки, при несовпадении - сдвигаем по таблице
//выставляем индекс каждому символу строки, потом заполняем значение смещения для каждого символа -
// - равное индексу этого символа, если символ уже встречался, то берется самое правое(меньшее) значение, для нулевого индекса - значение, равное значению строки
//смещение для символов, не встречающихся в шаблоне - длина строки
//если несовпадение произошло не на первом символе, а на последующем, берем смещение для последнего символа подстроки
func BoyerMureHorsepule(text, pattern string) any { //O(n)

	len_text := len(text)
	len_pattern := len(pattern)
	if len_pattern > len_text {
		return -1
	}
	if len_text || len_pattern == nil {
		return -1
	}

	offset := make(map[byte]int) //таблица смещений
	for _, i := range len_pattern - 1 {
		offset[pattern[i]] = len_pattern - 1 - i
	}

	skipVal := len_pattern - 1 //индекс последнего символа, с которого нужно начать проверку
	for skipVal < len_text {
		match := true
		var lastIndex int
		var currentLetter byte
		for i := 0; i < len_pattern; i++ {
			lastIndex = i
			currentLetter = text[skipVal-i]
			if pattern[len_pattern-i-1] != text[skipVal-i] {
				match = false
				break
			}
		}
		if match {
			return skipVal - len_pattern + 1
		}
		if lastIndex > 0 {
			skipVal += offset[pattern[len_pattern-1]]
			continue
		}

		skip, ok := offset[currentLetter]
		if !ok {
			skip = len_pattern
		}
		skipVal += skip

	}
	return -1
}

//метод рабина - карпа
//вычисляем хэш для подстроки
//для каждого возможного смещения в искомой строке рассчитывается хэш отрезка, равного длине подстроки
//сравниваются хэши отрезка строки и подстроки. при совпадении - сравнение символов
//при несовпадении смещаем отрезок поиска и продолжаем до конца строки

//скользящая хэш - функция - переиспользование вычисленного хэша
func Hash(l string) any {
	result := rune(l[0])
	x := 31
	q := 2147483647
	for _, i := range len(l) - 1 {
		result = result*x + rune(l[i+1])
	}
	return result % q
}
func RabinaKarpa(text, pattern string) any {
	x := 31
	q := 2147483647
	patternHash := Hash(pattern)
	m := len(pattern)
	n := len(text)

	if m > n {
		return -1
	}
	patternHash = simpleHash(pattern, x, q)
	currentHash := simpleHash(text[:m], x, q)

	highPow := 1
	for i := 0; i < m-1; i++ {
		highPow = (highPow * x) % q
	}
	for i := 0; i < n-m; i++ {
		if patternHash == currentHash {
			if text[i:i+m] == pattern {
				return i
			}
		}
		if i < n-m {
			currentHash = (currentHash - int(text[i])*highPow) * x
			currentHash = (currentHash + int(text[i+m])) % q
			if currentHash < 0 {
				currentHash += q
			}
		}
	}
	return -1
}
