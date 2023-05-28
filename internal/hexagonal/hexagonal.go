package hexagonal

import (
	"context"

	"github.com/go-git/go-git/v5"
)

type Hexagonal struct {
	gitUrl string
	path   string
}

func NewHexagonal(gitUrl, path string) Hexagonal {
	return Hexagonal{gitUrl: gitUrl, path: path}
}

func (h *Hexagonal) Generate(ctx context.Context) error {
	cloneOptions := new(git.CloneOptions)
	cloneOptions.URL = h.gitUrl
	_, err := git.PlainCloneContext(ctx, h.path, false, cloneOptions)
	if err != nil {
		return err
	}
	return nil
}
