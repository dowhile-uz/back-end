package editorControllerFx

import (
	"context"
)

type (
	CreateProjectInput struct {
		Body struct {
			IsCommunityProject bool   `json:"is-community"`
			RepoName           string `json:"repo-name"`
		}
	}
	CreateProjectOutput struct {
		Body struct {
			IsAvailable bool `json:"is-available"`
		}
	}
)

func (c *Controller) CreateProjectHandler(ctx context.Context, input *CreateProjectInput) (*CreateProjectOutput, error) {
	output := &CreateProjectOutput{}

	if input.Body.IsCommunityProject {
		isAvailable, err := c.service.CommunityRepoAvailabilityCheck(ctx, input.Body.RepoName)
		if err != nil {
			return nil, err
		}

		output.Body.IsAvailable = isAvailable
	}

	return output, nil
}
