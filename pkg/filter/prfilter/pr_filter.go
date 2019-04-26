package prfilter

import (
	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"
)

// FilterByAssignee filters given pulls according assignee.
func FilterByAssignee(prs []entity.PullRequest, assignee entity.Assignee) []entity.PullRequest {
	resultSet := []entity.PullRequest{}

	for _, pr := range prs {
		if assigneeIndex(pr, assignee) != -1 {
			resultSet = append(resultSet, pr)
		}
	}

	return resultSet
}

// FilterByReviewer filters given pulls according reviewer.
func FilterByReviewer(prs []entity.PullRequest, reviewer entity.Reviewer) []entity.PullRequest {
	resultSet := []entity.PullRequest{}

	for _, pr := range prs {
		if reviewerIndex(pr, reviewer) != -1 {
			resultSet = append(resultSet, pr)
		}
	}

	return resultSet
}

// FilterByParticipation filters given pulls according user participation.
func FilterByParticipation(prs []entity.PullRequest, user entity.User) []entity.PullRequest {
	resultSet := []entity.PullRequest{}

	for _, pr := range prs {
		if checkParticipation(pr, user) {
			resultSet = append(resultSet, pr)
		}
	}

	return resultSet
}

// FilterByState filters given pulls according pull state.
func FilterByState(prs []entity.PullRequest, state string) []entity.PullRequest {
	resultSet := []entity.PullRequest{}

	for _, pr := range prs {
		if pr.State == state {
			resultSet = append(resultSet, pr)
		}
	}

	return resultSet
}

func assigneeIndex(pr entity.PullRequest, assignee entity.Assignee) int {
	for i, a := range pr.Assignees {
		if a == assignee {
			return i
		}
	}

	return -1
}

func reviewerIndex(pr entity.PullRequest, reviewer entity.Reviewer) int {
	for i, r := range pr.RequestedReviewers {
		if r == reviewer {
			return i
		}
	}

	return -1
}

func checkParticipation(pr entity.PullRequest, user entity.User) bool {
	return ((assigneeIndex(pr, entity.Assignee(user)) != -1) || (reviewerIndex(pr, entity.Reviewer(user)) != -1))
}
