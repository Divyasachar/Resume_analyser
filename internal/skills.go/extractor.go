package skills

import "strings"

func ExtractSkills(text string) []string {

	var result []string

	lower, _ := IsAlias(strings.ToLower(text))
	for _, skill := range knownSkills {

		if strings.Contains(
			lower,
			strings.ToLower(skill),
		) {

			result = append(result, skill)
		}
	}
	return result
}
func ExtractEmail(text string) string {

	email := emailRegex.FindString(text)

	return email
}

func ExtractPhone(text string) string {

	phone := phoneRegex.FindString(text)

	return phone
}
func ExtractName(text string) string {

	lines := strings.Split(text, "\n")

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line != "" {

			return line
		}
	}

	return ""
}
func ExtractTitle(text string) string {

	lines := strings.Split(text, "\n")

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line != "" {

			return line
		}
	}

	return ""
}
func SplitSections(text string) (string, string) {

	lower := strings.ToLower(text)

	required := lower

	preferred := ""

	if strings.Contains(lower, "preferred") {

		parts := strings.Split(lower, "preferred")

		required = parts[0]

		if len(parts) > 1 {
			preferred = parts[1]
		}
	}

	return required, preferred
}
