package captcha

import (
	"context"

	"github.com/ginger-core/errors"
)

type Generator interface {
	New(ctx context.Context) (*Captcha, errors.Error)
}
