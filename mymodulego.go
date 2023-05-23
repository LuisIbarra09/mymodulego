package mymodulego

func SayHello() string {
	return "Hola todo mundo!"
}

func Sum(numbers ...int) (total int) {
	for _, n := range numbers {
		total += n
	}
	return
}
