package structures

type Course struct {
	Id       int      `json:"id,omitempty" `
	Name     string   `json:"name" form:"name,omitempty"`
	Tags     []string `json:"tags,omitempty" form:"tags,omitempty"`
	Duration string   `json:"duration,omitempty" form:"duration,omitempty"`
}
