package structures

type Course struct {
	Id       int      `json:"id,omitempty"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags,omitempty"`
	Duration string   `json:"duration,omitempty"`
}
