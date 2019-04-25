package entity

type Assignee User
type Reviewer User

type PullRequest struct {
	HTMLURL            string     `json:"html_url"`
	State              string     `json:"state"`
	Title              string     `json:"title"`
	User               User       `json:"user"`
	Assignees          []Assignee `json:"assignees"`
	RequestedReviewers []Reviewer `json:"requested_reviewers"`
	Head               head       `json:"head"`
}

type head struct {
	Repo Repository `json:"repo"`
}
