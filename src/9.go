package main
import "fmt"
func main() {
	name := "Alice"
	age := 30
	if name == "Alice" && age > 25 {
		fmt.Println("Hello, ", name)
	} else if age < 18 {
		fmt.Println("Sorry, you are not old enough to enter")
	} else {
		fmt.Println("You cannot come in")
	}
}