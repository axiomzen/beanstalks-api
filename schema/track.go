package schema

type Track struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Tags     []string  `json:"tags"`
	Feedback *Feedback `json:"feedback"`
	Levels   []*Levels `json:"levels"`
}
