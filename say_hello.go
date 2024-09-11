package go_say_hello

func SayHello() string {
	return "Hello World"
}

func SayName(name string) string{
	return "Hello " + name
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Azka"
	lastName = "Faridi"

	return firstName, middleName, lastName
}

func getFullName() (string, string) {
	return "Azka", "Faridi"
}

func MathPlus(a, b int) int {
	return a + b
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func MathDivide(a, b int) any {
	if b == 0 {
		return "not definition"
	} else {
		return a / b
	}
}