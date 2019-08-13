package main

const englishHelloPrefix = "Hello, "
const bgHelloPrefix = "Здрасти, "
const frenchHelloPrefix = "Bonjour, "
const bulgarian = "bulgarian"
const french = "french"

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name

}

func greetingPrefix(lang string) string {
	prefix := englishHelloPrefix
	switch lang {
	case bulgarian:
		prefix = bgHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix
}

func main() {
}
