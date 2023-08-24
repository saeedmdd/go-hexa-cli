package hexagonal

import (
	"context"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Hexagonal struct {
	gitUrl string
	path   string
	branch string
}

func NewHexagonal(gitUrl, path, branch string) Hexagonal {
	return Hexagonal{gitUrl: gitUrl, path: path, branch: branch}
}

func (h *Hexagonal) Generate(ctx context.Context) error {
	cloneOptions := new(git.CloneOptions)
	cloneOptions.URL = h.gitUrl
	cloneOptions.ReferenceName = plumbing.ReferenceName(h.branch)
	_, err := git.PlainCloneContext(ctx, h.path, false, cloneOptions)
	if err != nil {
		return err
	}
	return nil
}
