package Add

import "fmt"

func Sum(a ...int) {
	sum := 0
	for _, v := range a {
		sum = v + sum
	}
	fmt.Println(sum)

}

// func main() {

// 	Sum(11, 11, 11, 11)
// }
