package skills

import "strings"

func IsAlias(skill string) (string, bool) {
	if alias, exists := aliases[skill]; exists || strings.Contains(alias, skill) {
		return alias, true
	}
	return skill, false
}
