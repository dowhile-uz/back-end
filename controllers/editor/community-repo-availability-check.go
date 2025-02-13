package editorControllerFx

import "context"

type (
	CommunityRepoAvailabilityCheckInput struct {
		Body struct {
			RepoName string `json:"repo-name"`
		}
	}
	CommunityRepoAvailabilityCheckOutput struct {
		Body struct {
			IsAvailable bool `json:"is-available"`
		}
	}
)

func (c *Controller) CommunityRepoAvailabilityCheckHandler(ctx context.Context, input *CommunityRepoAvailabilityCheckInput) (*CommunityRepoAvailabilityCheckOutput, error) {
	output := &CommunityRepoAvailabilityCheckOutput{}

	isAvailable, err := c.service.CommunityRepoAvailabilityCheck(ctx, input.Body.RepoName)
	if err != nil {
		return nil, err
	}

	output.Body.IsAvailable = isAvailable

	return output, nil
}
