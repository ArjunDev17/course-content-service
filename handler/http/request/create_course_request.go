package request

type CreateCourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Instructor  string `json:"instructor"`
}