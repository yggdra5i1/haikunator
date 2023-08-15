package haikunator

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	adjectives = []string{
		"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
		"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
		"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
		"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
		"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
		"sparkling", "thrumming", "shy", "wandering", "withered", "wild", "black",
		"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
		"restless", "divine", "polished", "ancient", "purple", "lively", "nameless",
	}

	nouns = []string{
		"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
		"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
		"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
		"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
		"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
		"sound", "sky", "shape", "surf", "thunder", "violet",
	}
)

func Build(tokenRange int, delimiter string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	tkn := token(tokenRange)
	sections := []string{
		adjectives[rand.Intn(len(adjectives))],
		nouns[rand.Intn(len(nouns))],
	}
	if len(tkn) > 0 {
		sections = append(sections, tkn)
	}
	return join(sections, delimiter)
}

func BuildDefault() string {
	return Build(9999, "-")
}

func BuildWithTokenRange(tokenRange int) string {
	return Build(tokenRange, "-")
}

func BuildWithDelimiter(delimiter string) string {
	return Build(9999, delimiter)
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
