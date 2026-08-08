package twofer

import "strings"

const placeholderString = "__PLACEHOLDER__"

func ShareWith(name string) string {
	baseString := "One for __PLACEHOLDER__, one for me."
	finalString := ""
	if len(name) != 0 {
		finalString = strings.ReplaceAll(baseString, placeholderString, name)
	} else {
		finalString = strings.ReplaceAll(baseString, placeholderString, "you")
	}

	return finalString
}
