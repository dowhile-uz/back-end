package helloworldcontrollerfx

import "context"

type (
	HelloWorldInput  struct{}
	HelloWorldOutput struct {
		Body struct {
			Message string `json:"message" example:"Hello, World!" doc:"Greeting message"`
		}
	}
)

func (h *Controller) HelloWorldHandler(ctx context.Context, input *HelloWorldInput) (*HelloWorldOutput, error) {
	output := &HelloWorldOutput{}

	output.Body.Message = "Hello, World!"

	return output, nil
}
