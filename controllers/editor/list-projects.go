package editorControllerFx

import "context"

type (
	ListProjectsInput  struct{}
	ListProjectsOutput struct{}
)

func (c *Controller) ListProjects(ctx context.Context, input *ListProjectsInput) {}
