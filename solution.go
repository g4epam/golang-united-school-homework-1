package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	worldMap := emoji.Sprint("Hello :world_map:!")
	return worldMap
}
