package main

import "fmt"

const french = "francês"
const italian = "italiano"
const spanish = "espanhol"
const prefixHelloFrench = "Bonjour, "
const prefixHelloItalian = "Ciao, "
const prefixHelloPortuguese = "Olá, "
const prefixHelloSpanish = "Hola, "

func prefixGreetingByIdiom(idiom string) (prefix string) {
	switch idiom {
	case french:
		prefix = prefixHelloFrench
	case italian:
		prefix = prefixHelloItalian
	case spanish:
		prefix = prefixHelloSpanish
	default:
		prefix = prefixHelloPortuguese
	}
	return
}

func Hello(name string, idiom string) string {
	if name == "" {
		name = "mundo"
	}
	return prefixGreetingByIdiom(idiom) + name

}

func main() {
	fmt.Println(Hello("mundo", "português"))
}
