package domain

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
	Label    string `json:"label"`
	Priority string `json:"priority"`
}
