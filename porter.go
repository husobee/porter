package porter

import "strings"

func isConsonant(word string, index int) bool {
	switch word[index] {
	case 'a', 'e', 'i', 'o', 'u':
		return false
	case 'y':
		if index == 0 {
			return false
		} else if !isConsonant(word, index-1) {
			return false
		}
		return true
	}
	return true
}

// [C](VC){m}[V] <- solve for m
func wordMeasure(word string) int {
	var foundVStart = false
	var mCount int = 0
	for idx, _ := range word {
		// find first vowel sequence
		if !isConsonant(word, idx) {
			if !foundVStart {
				foundVStart = true
			}
			if idx == len(word)-1 {
				return mCount
			}
			continue
		}
		if foundVStart && isConsonant(word, idx) {
			if len(word)-1 == idx {
				mCount++
				break
			}
			if !isConsonant(word, idx+1) {
				mCount++
			}
		}
	}
	return mCount
}

func containsVowel(word string) bool {
	for idx, _ := range word {
		if !isConsonant(word, idx) {
			return true
		}
	}
	return false
}

func performReplacement(word string, replacementMapping [][]string) string {
	var newWord = word
	for _, mapping := range replacementMapping {
		if strings.HasSuffix(newWord, mapping[0]) {
			newWord = strings.TrimSuffix(newWord, mapping[0]) + mapping[1]
			break
		}
	}
	return newWord
}

func step1a(word string) string {
	var suffixMapping = [][]string{
		[]string{"sses", "ss"},
		[]string{"ies", "i"},
		[]string{"ss", "ss"},
		[]string{"s", ""},
	}
	return performReplacement(word, suffixMapping)
}

func step1b(word string) string {
	var newWord = string(word)
	var furtherProcessing = false
	if strings.HasSuffix(string(word), "eed") {
		stem := strings.TrimSuffix(newWord, "eed") + "ee"
		if wordMeasure(string(stem)) > 0 {
			return string(stem)
		}
		return string(newWord)
	}

	if strings.HasSuffix(newWord, "ed") {
		stem := strings.TrimSuffix(newWord, "ed")
		if containsVowel(string(stem)) {
			furtherProcessing = true
			newWord = stem
		}
	}
	if strings.HasSuffix(newWord, "ing") {
		stem := strings.TrimSuffix(newWord, "ing")
		if containsVowel(string(stem)) {
			furtherProcessing = true
			newWord = stem
		}
	}
	if furtherProcessing {
		var suffixMapping = [][]string{
			[]string{"at", "ate"},
			[]string{"bl", "ble"},
			[]string{"iz", "ize"},
		}
		newWord = performReplacement(newWord, suffixMapping)
		if len(newWord) < 2 {
			return newWord
		}
		if newWord[len(newWord)-1] != 'l' && newWord[len(newWord)-1] != 's' && newWord[len(newWord)-1] != 'z' {
			if newWord[len(newWord)-1] == newWord[len(newWord)-2] && isConsonant(newWord, len(newWord)-1) {
				newWord = newWord[0 : len(newWord)-1]
				return newWord
			}
		}
		if newWord[len(newWord)-1] != 'w' && newWord[len(newWord)-1] != 'x' && newWord[len(newWord)-1] != 'y' && len(newWord) > 2 {
			if wordMeasure(newWord) == 1 && isConsonant(newWord, len(newWord)-1) && !isConsonant(newWord, len(newWord)-2) && isConsonant(newWord, len(newWord)-3) {
				return newWord + "e"
			}

		}
	}
	return newWord
}
func step1c(word string) string {
	if word[len(word)-1] == 'y' && containsVowel(word[:len(word)-1]) {
		return word[:len(word)-1] + "i"
	}
	return word
}

func step2(word string) string {
	var suffixMapping = [][]string{
		[]string{"ational", "ate"},
		[]string{"tional", "tion"},
		[]string{"enci", "ence"},
		[]string{"anci", "ance"},
		[]string{"izer", "ize"},
		[]string{"abli", "able"},
		[]string{"alli", "al"},
		[]string{"entli", "ent"},
		[]string{"eli", "e"},
		[]string{"ousli", "ous"},
		[]string{"ization", "ize"},
		[]string{"ation", "ate"},
		[]string{"ator", "ate"},
		[]string{"alism", "al"},
		[]string{"iveness", "ive"},
		[]string{"fulness", "ful"},
		[]string{"ousness", "ous"},
		[]string{"aliti", "al"},
		[]string{"iviti", "ive"},
		[]string{"biliti", "ble"},
	}
	if wordMeasure(word) > 0 {
		return performReplacement(word, suffixMapping)
	}
	return word
}

func step3(word string) string {
	var suffixMapping = [][]string{
		[]string{"icate", "ic"},
		[]string{"ative", ""},
		[]string{"alize", "al"},
		[]string{"iciti", "ic"},
		[]string{"ical ", "ic"},
		[]string{"ful", ""},
		[]string{"ness", ""},
	}
	if wordMeasure(word) > 0 {
		return performReplacement(word, suffixMapping)
	}
	return word
}

func step4(word string) string {
	var suffixMapping = [][]string{
		[]string{"al", ""},
		[]string{"ance", ""},
		[]string{"ence", ""},
		[]string{"er", ""},
		[]string{"ic", ""},
		[]string{"able", ""},
		[]string{"ible", ""},
		[]string{"ant", ""},
		[]string{"ement", ""},
		[]string{"ment", ""},
		[]string{"ent", ""},
		[]string{"ou", ""},
		[]string{"ism", ""},
		[]string{"ate", ""},
		[]string{"iti", ""},
		[]string{"ous", ""},
		[]string{"ive", ""},
		[]string{"ize", ""},
	}
	//(m>1 and (*S or *T)) ION ->     adoption       ->  adopt

	if wordMeasure(word) > 1 {
		if word[len(word)-1] == 's' || word[len(word)-1] == 't' {
			suffixMapping = append(suffixMapping, []string{"TION", ""})
		}
		return performReplacement(word, suffixMapping)
	}
	return word
}

func step5a(word string) string {
	var suffixMapping = [][]string{
		[]string{"e", ""},
	}

	nw := strings.TrimSuffix(word, "e")
	if wordMeasure(nw) > 1 {
		return performReplacement(word, suffixMapping)
	}
	if len(word) < 3 {
		return word
	}
	if (nw[len(nw)-1] == 'w' || nw[len(nw)-1] == 'x' || nw[len(nw)-1] != 'y') && len(nw) > 2 {
		if wordMeasure(nw) == 1 && !(isConsonant(nw, len(nw)-1) && !isConsonant(nw, len(nw)-2) && isConsonant(nw, len(nw)-3)) {
			return performReplacement(word, suffixMapping)
		}
	}
	return word
}

func step5b(word string) string {
	var suffixMapping = [][]string{
		[]string{"l", ""},
	}
	if wordMeasure(word) > 1 && word[len(word)-2] == word[len(word)-1] && word[len(word)-2] == 'l' {
		return performReplacement(word, suffixMapping)
	}
	return word
}
