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

func MathPlus(a, b int) int {
	return a + b
}

func getFullName() (string, string) {
	return "Azka", "Faridi"
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}