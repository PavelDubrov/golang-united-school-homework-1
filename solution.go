package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	emj := emoji.Sprint("Hello :world_map:!")
	return emj
}
