package main
import "fmt"
import "./sha1"

func main() {
	fmt.Println("Hello Go World!")
	message := []byte{0xFF, 0xFF}
	var result = sha1.Padding(message)
	sha1.OutputBit(result)
}
