package matcher

import "strings"

func toSet(skills []string) map[string]bool {

	set := make(map[string]bool)

	for _, skill := range skills {

		set[strings.ToLower(skill)] = true
	}

	return set
}