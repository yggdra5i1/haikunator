package main

import (
	"fmt"
	"math/rand"
	"time"
)

var adjectives = []string{
	"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
	"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
	"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
	"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
	"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
	"sparkling", "thrumming", "shy", "wandering", "withered", "wild", "black",
	"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
	"restless", "divine", "polished", "ancient", "purple", "lively", "nameless",
}

var nouns = []string{
	"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
	"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
	"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
	"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
	"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
	"sound", "sky", "shape", "surf", "thunder", "violet",
}

func Haikunate(tokenRange int, delimiter string) string {
	rand.Seed(time.Now().UnixNano())
	sections := []string{
		adjectives[rand.Intn(len(adjectives))],
		nouns[rand.Intn(len(nouns))],
		token(tokenRange),
	}
	return join(sections, delimiter)
}

func token(rangeNum int) string {
	if rangeNum > 0 {
		return fmt.Sprintf("%d", rand.Intn(rangeNum))
	}
	return ""
}

func join(strs []string, delimiter string) string {
	result := ""
	for i, str := range strs {
		if i > 0 {
			result += delimiter
		}
		result += str
	}
	return result
}
