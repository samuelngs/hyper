package server

import (
	"context"

	"github.com/samuelngs/hyper/gql/interfaces"
	"github.com/samuelngs/hyper/router"
)

// FromContext reads router context from context.Context
func FromContext(c context.Context) interfaces.Context {
	return &ctx{c.Value(router.RequestContext).(router.Context)}
}
