package ai

import (
	"context"

	"github.com/Southclaws/storyden/internal/config"
)

type Result struct {
	Answer string
}

type Prompter interface {
	Prompt(ctx context.Context, input string) (*Result, error)
}

func New(cfg config.Config) (Prompter, error) {
	switch cfg.LanguageModelProvider {
	case "openai":
		return newOpenAI(cfg)

	default:
		return &Disabled{}, nil
	}
}

type Disabled struct{}

func (d *Disabled) Prompt(ctx context.Context, input string) (*Result, error) {
	return nil, nil
}
