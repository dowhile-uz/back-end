package githubauthcontrollerfx

import (
	"context"
	"fmt"
)

type (
	CompleteInput struct {
		Back string `query:"back"`
		Code string `query:"code" required:"true"`
	}
	CompleteOutput struct {
		Body struct {
			Message string `json:"message"`
		}
	}
)

func (c *Controller) CompleteHandler(ctx context.Context, input *CompleteInput) (*CompleteOutput, error) {
	o := &CompleteOutput{}
	o.Body.Message = fmt.Sprintf("code: %s", input.Code)
	return o, nil
}
