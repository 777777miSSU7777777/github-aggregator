package entity

// Assignee alias to user entity.
type Assignee User

// Reviewer alias to user entity.
type Reviewer User

// PullRequest is an entity which represents github pull request.
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
