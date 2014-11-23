package porter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestIsConsonant(t *testing.T) {
	if !isConsonant(string("testing"), 0) {
		t.Fail()
		t.Log("t should be a consonant")
	}
	if isConsonant(string("testing"), 1) {
		t.Fail()
		t.Log("e should be a vowel")
	}
	if isConsonant(string("toy"), 2) {
		t.Fail()
		t.Log("y should be a consonant")
	}
	if !isConsonant(string("likely"), 5) {
		t.Fail()
		t.Log("y should be a vowel")
	}
}

func TestWordMeasure(t *testing.T) {
	measure := map[int][]string{
		0: []string{"tr", "ee", "tree", "y", "by"},
		1: []string{"trouble", "oats", "trees", "ivy"},
		2: []string{"troubles", "private", "oaten", "orrery"},
	}

	for k, v := range measure {
		for _, s := range v {
			if m := wordMeasure(string(s)); m != k {
				t.Fail()
				t.Log(fmt.Sprintf("word measure isn't correct: %s, %d != %d", s, m, k))
			}
		}
	}
}

func Test1a(t *testing.T) {
	if result := step1a(string("caresses")); len(result) != 6 {
		t.Log(string(result))
		t.Fail()
	}
}

func Test1b(t *testing.T) {

	testData := map[string]string{
		"feed":      "feed",
		"agreed":    "agree",
		"plastered": "plaster",
		"bled":      "bled",
		"motoring":  "motor",
		"sing":      "sing",
		"conflated": "conflate",
		"troubled":  "trouble",
		"sized":     "size",
		"hopping":   "hop",
		"tanned":    "tan",
		"falling":   "fall",
		"hissing":   "hiss",
		"fizzed":    "fizz",
		"failing":   "fail",
		"filing":    "file",
	}
	for k, v := range testData {
		if result := step1b(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func Test1c(t *testing.T) {

	testData := map[string]string{
		"sky":   "sky",
		"happy": "happi",
	}
	for k, v := range testData {
		if result := step1c(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func Test2(t *testing.T) {

	testData := map[string]string{
		"analogousli":    "analogous",
		"vietnamization": "vietnamize",
		"predication":    "predicate",
		"operator":       "operate",
		"feudalism":      "feudal",
		"decisiveness":   "decisive",
		"hopefulness":    "hopeful",
		"callousness":    "callous",
		"formaliti":      "formal",
		"sensitiviti":    "sensitive",
		"sensibiliti":    "sensible",
	}
	for k, v := range testData {
		if result := step2(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}
func Test3(t *testing.T) {

	testData := map[string]string{
		"triplicate":  "triplic",
		"formative":   "form",
		"formalize":   "formal",
		"electriciti": "electric",
		"hopeful":     "hope",
		"goodness":    "good",
		//"electrical":  "electric",
	}
	for k, v := range testData {
		if result := step3(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}
func Test4(t *testing.T) {

	testData := map[string]string{
		"revival":     "reviv",
		"allowance":   "allow",
		"inference":   "infer",
		"airliner":    "airlin",
		"gyroscopic":  "gyroscop",
		"adjustable":  "adjust",
		"defensible":  "defens",
		"irritant":    "irrit",
		"replacement": "replac",
		"adjustment":  "adjust",
		"dependent":   "depend",
		"homologou":   "homolog",
		"communism":   "commun",
		"activate":    "activ",
		"angulariti":  "angular",
		"homologous":  "homolog",
		"effective":   "effect",
		"bowdlerize":  "bowdler",
	}
	for k, v := range testData {
		if result := step4(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func Test5a(t *testing.T) {

	testData := map[string]string{
		"probate": "probat",
		"rate":    "rate",
		"cease":   "ceas",
	}
	for k, v := range testData {
		if result := step5a(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func Test5b(t *testing.T) {

	testData := map[string]string{
		"controll": "control",
		"roll":     "roll",
	}
	for k, v := range testData {
		if result := step5b(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func TestFullFlowTest(t *testing.T) {
	file, _ := os.Open("dictionary.txt")
	r := bufio.NewReader(file)
	defer file.Close()
	for {
		if str, err := r.ReadString('\n'); err == nil {
			str = strings.TrimSuffix(str, "\n")
			newStr := step5b(step5a(step4(step3(step2(step1c(step1b(step1a(str))))))))
			fmt.Println(newStr)
		} else {
			break
		}
	}
}
