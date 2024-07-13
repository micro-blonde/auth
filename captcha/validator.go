package captcha

import (
	"context"

	"github.com/ginger-core/errors"
)

type Validator interface {
	Validate(ctx context.Context, secret, code string) errors.Error
}
