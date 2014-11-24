package porter

import "testing"

func Test1a(t *testing.T) {
	if result, _ := step1a(string("caresses")); len(result) != 6 {
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
		if result, _ := step1b(k); result != v {
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
		if result, _ := step1c(k); result != v {
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
		if result, _ := step2(k); result != v {
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
		"electrical":  "electric",
	}
	for k, v := range testData {
		if result, _ := step3(k); result != v {
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
		if result, _ := step4(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func Test5a(t *testing.T) {

	testData := map[string]string{
		"probate": "probat",
		"rate":    "rate",
		"cease":   "cease",
	}
	for k, v := range testData {
		if result, _ := step5a(k); result != v {
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
		if result, _ := step5b(k); result != v {
			t.Log(k, result)
			t.Fail()
		}
	}
}

func TestStem(t *testing.T) {
	words := []string{"ramdomize", "something", "forth", "the", "machine", "to", "crunchy"}
	t.Log(Stem(words))
}
