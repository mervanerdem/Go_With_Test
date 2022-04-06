package main

const Turkish = "Turkish"
const English = "English"
const Spanish = "Spanish"

func Hello(s, lang string) string {
	if s == "" {
		s = "World"
	}
	return selectlang(lang) + s
}

func selectlang(lang string) (prefix string) {
	switch lang {
	case Turkish:
		prefix = "Merhaba "
	case Spanish:
		prefix = "Hola "
	default:
		prefix = "Hello "
	}
	return prefix
}
func main() {

}
