package usecase

import (
	"context"
	"fss/internal/domain"
)

var _ domain.ITextUsecase = new(text)

type text struct {
	repo domain.ITextRepo
}

func NewTextUsecase(repo domain.ITextRepo) *text {
	return &text{
		repo: repo,
	}
}

func (t *text) QueryTextByFilter(ctx context.Context, request *domain.TextFilter) ([]*domain.Text, error) {
	return t.repo.QueryTextByFilter(ctx, request)
}
