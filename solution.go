package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	book := emoji.Sprint("Hello :book:!")
	return book
}
