package profile

import (
	"context"

	"github.com/wendao2000/couply/pkg/errs"
)

func (s *service) CreateBasicProfile(ctx context.Context, req *BasicProfile) (err *errs.Error) {
	// TODO: push to MQ for notification
	return nil
}
