package placement

type Gender string

const (
	MALE         Gender = "male"
	FEMALE       Gender = "female"
	OTHER_GENDER Gender = "other"
)

type Criteria struct {
	Xth           float32  `json:"xth,omitempty"`
	XIIth         float32  `json:"xiith,omitempty"`
	CGPA          float32  `json:"cgpa,omitempty"`
	Branch        []string `json:"branch,omitempty"`
	Course        []string `json:"course,omitempty"`
	ActiveBacklog int      `json:"active_backlog,omitempty"`
	TotalBacklog  int      `json:"total_backlog,omitempty"`
	Gender        Gender   `json:"gender,omitempty"`
	Disability    bool     `json:"disability,omitempty"`
}
