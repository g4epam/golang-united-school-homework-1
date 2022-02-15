package solution

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	book := emoji.Sprint("Hello :book:!")
	fmt.Println(book)
	return book
}
