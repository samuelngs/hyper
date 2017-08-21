package hyper

import (
	"context"

	"github.com/samuelngs/hyper/router"
)

// Context reads router context from context.Context
func Context(c context.Context) router.Context {
	return c.Value(router.RequestContext).(router.Context)
}
