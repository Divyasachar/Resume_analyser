package skills

import "regexp"

var emailRegex = regexp.MustCompile(
	`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`,
)

var phoneRegex = regexp.MustCompile(
	`(\+?\d[\d\s-]{8,}\d)`,
)
var knownSkills = []string{
	"Go",
	"Java",
	"Python",
	"Docker",
	"Kubernetes",
	"AWS",
	"Terraform",
	"Kafka",
	"Redis",
	"MySQL",
	"PostgreSQL",
	"React",
	"Node.js",
	"Git",
	"Linux",
}
var aliases = map[string]string{
	"golang": "go",
	"postgres": "postgresql",
	"k8s": "kubernetes",
	"js": "javascript",
}
