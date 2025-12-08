package main

func main() {
	println(Hello(""))
}

const englishHelloPrefix = "Hello, "
const exclamationPoint = "!"

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name + exclamationPoint
}
