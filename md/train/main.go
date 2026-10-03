package train

import (
	"fmt"
	"train"
)

func main() {
	fmt.Println(train.IsCapitalized("Hello! How are you?"))
	fmt.Println(train.IsCapitalized("Hello How Are You"))
	fmt.Println(train.IsCapitalized("Whats 4this 100K?"))
	fmt.Println(train.IsCapitalized("Whatsthis4"))
	fmt.Println(train.IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(train.IsCapitalized(""))
}
