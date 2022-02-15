package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	book := emoji.Sprint("Hello :world_map:!")
	return book
}
