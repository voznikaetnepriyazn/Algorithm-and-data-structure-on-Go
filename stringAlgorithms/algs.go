package algs

//метод прямого перебора

//алгоритм бойера - мура- хорспула
//начать сравнение с последнего элемента подстроки и соответствующего по номеру элемента строки, при несовпадении - сдвигаем по таблице
//выставляем индекс каждому символу строки, потом заполняем значение смещения для каждого символа -
// - равное индексу этого символа, если символ уже встречался, то берется самое правое(меньшее) значение, для нулевого индекса - значение, равное значению строки
//смещение для символов, не встречающихся в шаблоне - длина строки
//если несовпадение произошло не на первом символе, а на последующем, берем смещение для последнего символа подстроки

func BoyerMooreHorspool(text, pattern string) int {
	lenText := len(text)
	lenPattern := len(pattern)

	// Проверка граничных случаев
	if lenPattern == 0 || lenPattern > lenText {
		return -1
	}

	// Построение таблицы смещений (эвристика плохого символа)
	// Для символов, не встречающихся в шаблоне, смещение = длина шаблона
	shift := make(map[byte]int)
	for i := 0; i < 256; i++ {
		shift[byte(i)] = lenPattern
	}

	// Для символов шаблона (кроме последнего) смещение = расстояние до конца
	for i := 0; i < lenPattern-1; i++ {
		shift[pattern[i]] = lenPattern - 1 - i
	}

	// Поиск подстроки
	i := lenPattern - 1 // начинаем с последнего символа шаблона
	for i < lenText {
		k := 0
		// Сравниваем символы справа налево
		for k < lenPattern && pattern[lenPattern-1-k] == text[i-k] {
			k++
		}

		if k == lenPattern {
			// Найдено совпадение
			return i - lenPattern + 1
		}

		// Сдвигаем шаблон на основе таблицы смещений
		// Используем символ текста, который выровнен с последним символом шаблона
		i += shift[text[i]]
	}

	return -1
}

//метод рабина - карпа
//вычисляем хэш для подстроки
//для каждого возможного смещения в искомой строке рассчитывается хэш отрезка, равного длине подстроки
//сравниваются хэши отрезка строки и подстроки. при совпадении - сравнение символов
//при несовпадении смещаем отрезок поиска и продолжаем до конца строки

//скользящая хэш-функция - переиспользование вычисленного хэша
func simpleHash(s string, base, mod int) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = (hash*base + int(s[i])) % mod
	}
	return hash
}

func RabinKarp(text, pattern string) int {
	const base = 256       // размер алфавита (расширенный ASCII)
	const mod = 1000000007 // большое простое число для модуля

	lenText := len(text)
	lenPattern := len(pattern)

	// Проверка граничных случаев
	if lenPattern == 0 || lenPattern > lenText {
		return -1
	}

	// Вычисляем хэш шаблона и первого окна текста
	patternHash := simpleHash(pattern, base, mod)
	windowHash := simpleHash(text[:lenPattern], base, mod)

	// Предвычисляем значение base^(lenPattern-1) % mod для скользящего хэша
	h := 1
	for i := 0; i < lenPattern-1; i++ {
		h = (h * base) % mod
	}

	// Поиск подстроки
	for i := 0; i <= lenText-lenPattern; i++ {
		// Если хэши совпадают — проверяем полное совпадение строк
		if patternHash == windowHash {
			if text[i:i+lenPattern] == pattern {
				return i
			}
		}

		// Обновляем хэш для следующего окна (скользящий хэш)
		// Удаляем старый символ слева и добавляем новый справа
		if i < lenText-lenPattern {
			windowHash = (windowHash - int(text[i])*h) % mod
			windowHash = (windowHash*base + int(text[i+lenPattern])) % mod

			// Обработка отрицательных значений
			if windowHash < 0 {
				windowHash += mod
			}
		}
	}

	return -1
}
