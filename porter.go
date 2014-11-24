package porter

import (
	"errors"
	"strings"
)

var IndexOutBoundsError = errors.New("Index out of Bounds")

func isConsonant(word string, index int) (bool, error) {
	if index < 0 || index > len(word) {
		return false, IndexOutBoundsError
	}
	switch word[index] {
	case 'a', 'e', 'i', 'o', 'u':
		return false, nil
	case 'y':
		consonant, err := isConsonant(word, index-1)
		if err == IndexOutBoundsError || !consonant {
			return false, nil
		}
		return true, nil
	}
	return true, nil
}

func lastPosition(word string) int {
	return len(word) - 1
}
func secondToLastPosition(word string) int {
	return len(word) - 2
}
func thirdToLastPosition(word string) int {
	return len(word) - 3
}

// [C](VC){m}[V] <- solve for m
func wordMeasure(word string) (int, error) {
	var foundVStart = false
	var mCount int = 0
	for idx, _ := range word {
		consonant, err := isConsonant(word, idx)
		if err != nil {
			return 0, err
		}
		// find first vowel sequence
		if !consonant {
			if !foundVStart {
				foundVStart = true
			}
			if idx == lastPosition(word) {
				return mCount, nil
			}
			continue
		}
		if foundVStart && consonant {
			if len(word)-1 == idx {
				mCount++
				break
			}
			next_consonant, err := isConsonant(word, idx+1)
			if err != nil && err == IndexOutBoundsError {
				break
			}
			if !next_consonant {
				mCount++
			}
		}
	}
	return mCount, nil
}

func containsVowel(word string) bool {
	for idx, _ := range word {
		if consonant, err := isConsonant(word, idx); err == nil && !consonant {
			return true
		}
	}
	return false
}

func performReplacement(word string, replacementMapping []func(string) (string, bool)) (string, bool) {
	for _, f := range replacementMapping {
		if newWord, modified := f(word); modified {
			return newWord, modified
		}
	}
	return word, false
}

func directReplace(word, suffix, replacement string) (string, bool) {
	if strings.HasSuffix(word, suffix) {
		return word[:len(word)-len(suffix)] + replacement, true
	}
	return word, false
}

func step1a(word string) (string, bool) {
	var transformMapping = []func(string) (string, bool){
		func(s string) (string, bool) { return directReplace(word, "sses", "ss") },
		func(s string) (string, bool) { return directReplace(word, "ies", "i") },
		func(s string) (string, bool) { return directReplace(word, "ss", "ss") },
		func(s string) (string, bool) { return directReplace(word, "s", "") },
	}
	return performReplacement(word, transformMapping)
}

func step1b(word string) (string, bool) {

	var furtherProcessingFuncs = []func(string) (string, bool){
		func(s string) (string, bool) { return directReplace(s, "at", "ate") },
		func(s string) (string, bool) { return directReplace(s, "bl", "ble") },
		func(s string) (string, bool) { return directReplace(s, "iz", "ize") },
	}
	var furtherProcessing = func(s string) (string, bool) {
		stem, _ := performReplacement(s, furtherProcessingFuncs)
		if len(stem) < 2 {
			return stem, true
		}
		if stem[lastPosition(stem)] != 'l' && stem[lastPosition(stem)] != 's' && stem[lastPosition(stem)] != 'z' {
			consonant, err := isConsonant(stem, lastPosition(stem))
			if err != nil {
				return stem, true
			}
			if stem[lastPosition(stem)] == stem[secondToLastPosition(stem)] && consonant {
				stem = stem[0:lastPosition(stem)]
				return stem, true
			}
		}
		if (stem[lastPosition(stem)] != 'w' && stem[lastPosition(stem)] != 'x' && stem[lastPosition(stem)] != 'y') && len(stem) > 2 {
			last_is_consonant, _ := isConsonant(stem, lastPosition(stem))
			second_last_is_consonant, _ := isConsonant(stem, secondToLastPosition(stem))
			third_last_is_consonant, _ := isConsonant(stem, thirdToLastPosition(stem))
			m, _ := wordMeasure(stem)
			if m == 1 && last_is_consonant && !second_last_is_consonant && third_last_is_consonant {
				return stem + "e", true
			}
		}
		return stem, true
	}
	var transformMapping = []func(string) (string, bool){
		func(s string) (string, bool) {
			stem, modified := directReplace(s, "eed", "ee")
			if m, err := wordMeasure(stem); err == nil && m > 0 {
				return stem, modified
			}
			return s, true
		},
		func(s string) (string, bool) {
			if stem, modified := directReplace(s, "ed", ""); modified {
				if containsVowel(stem) {
					//further processing
					return furtherProcessing(stem)
				}
			}
			return s, false
		},
		func(s string) (string, bool) {
			if stem, modified := directReplace(s, "ing", ""); modified {
				if containsVowel(stem) {
					//further processing
					return furtherProcessing(stem)
				}
			}
			return s, false
		},
	}
	return performReplacement(word, transformMapping)
}

func step1c(word string) (string, bool) {
	if word[len(word)-1] == 'y' && containsVowel(word[:len(word)-1]) {
		return word[:len(word)-1] + "i", true
	}
	return word, false
}

func mGreaterThan(m int, orig string, stemFunc func(string, string, string) (string, bool), suffix, to string) (string, bool) {
	stem, modified := stemFunc(orig, suffix, to)
	if measured, _ := wordMeasure(orig); measured > m {
		return stem, modified
	} else {
		return orig, false
	}
}

func step2(word string) (string, bool) {
	var transformMapping = []func(string) (string, bool){
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ational", "ate") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "tional", "tion") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "enci", "ence") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "anci", "ance") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "izer", "ize") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "abli", "able") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "alli", "al") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "entli", "ent") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "eli", "e") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ousli", "ous") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ization", "ize") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ation", "ate") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ator", "ate") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "alism", "al") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "iveness", "ive") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "fulness", "ful") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ousness", "ous") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "aliti", "al") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "iviti", "ive") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "biliti", "ble") },
	}
	return performReplacement(word, transformMapping)
}

func step3(word string) (string, bool) {
	var transformMapping = []func(string) (string, bool){
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "icate", "ic") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ative", "") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "alize", "al") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "iciti", "ic") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ical", "ic") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ful", "") },
		func(s string) (string, bool) { return mGreaterThan(0, s, directReplace, "ness", "") },
	}
	return performReplacement(word, transformMapping)
}

func step4(word string) (string, bool) {

	var transformMapping = []func(string) (string, bool){
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "al", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ance", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ence", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "er", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ic", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "able", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ible", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ant", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ement", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ment", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ent", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ou", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ism", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ate", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "iti", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ous", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ive", "") },
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "ize", "") },
		//(m>1 and (*S or *T)) ION ->     adoption       ->  adopt
		func(s string) (string, bool) {
			if s[lastPosition(s)] == 's' || s[lastPosition(s)] == 't' {
				return mGreaterThan(1, s, directReplace, "tion", "")
			}
			return s, false
		},
	}
	return performReplacement(word, transformMapping)
}

func step5a(word string) (string, bool) {
	var transformMapping = []func(string) (string, bool){
		func(s string) (string, bool) { return mGreaterThan(1, s, directReplace, "e", "") },
		func(s string) (string, bool) {
			if (s[lastPosition(s)] == 'w' || s[lastPosition(s)] == 'x' || s[lastPosition(s)] != 'y') && len(s) > 2 {
				s = strings.TrimSuffix(s, "e")
				last_is_consonant, err := isConsonant(s, lastPosition(s))
				second_last_is_consonant, err := isConsonant(s, secondToLastPosition(s))
				third_last_is_consonant, err := isConsonant(s, thirdToLastPosition(s))
				if err != nil {
					return s, false
				}
				m, _ := wordMeasure(s)
				if m == 1 && !last_is_consonant && second_last_is_consonant && !third_last_is_consonant {
					return directReplace(s, "e", "")
				}
			}
			return s, false
		},
	}
	return performReplacement(word, transformMapping)
}

func step5b(word string) (string, bool) {
	m, err := wordMeasure(word)
	if err != nil {
		return word, false
	}

	if m > 1 && word[lastPosition(word)] == word[secondToLastPosition(word)] && word[secondToLastPosition(word)] == 'l' {
		return strings.TrimSuffix(word, "l"), true
	}
	return word, false
}

func Stem(words []string) []string {
	var w []string
	for _, word := range words {
		word, _ = step1a(word)
		word, _ = step1b(word)
		word, _ = step1c(word)
		word, _ = step2(word)
		word, _ = step3(word)
		word, _ = step4(word)
		word, _ = step5a(word)
		word, _ = step5b(word)
		w = append(w, word)
	}
	return w
}
