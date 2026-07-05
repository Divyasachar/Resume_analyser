package resume

type Resume struct {
	Name string `json:"name"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Skills []string `json:"skills"`
}