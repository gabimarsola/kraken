package structure

type PRD struct {
	Title                string
	Introduction         string
	Objectives           []string
	UserStories          []UserStory
	FunctionalReqs       []FunctionalRequirement
	OutOfScope           []string
	DesignConsiderations []string
	TechConsiderations   []string
	SuccessMetrics       []string
	OpenQuestions        []string
}

type UserStory struct {
	ID                 string
	Title              string
	Description        string
	AcceptanceCriteria []string
}

type FunctionalRequirement struct {
	ID          string
	Title       string
	Description string
	Priority    string // high, medium, low
}

type PRDData struct {
	PRD         PRD
	ProjectInfo *ProjectInfo
}
