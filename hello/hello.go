package hello

const (
	spanish = "ES"
	french = "FR"
	german = "DE"
	englishHelloPrefix = "Hello, "
 	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	germanHelloPrefix = "Hallo, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	case german:
		return germanHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}


