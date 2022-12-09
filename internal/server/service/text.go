package service

import (
	"context"
	"errors"
	"fss/internal/domain"
)

type Text struct {
	textUsecase domain.ITextUsecase
	curlUsecase domain.ITextUsecase
}

func NewTextService(usecase domain.ITextUsecase, curlUsecase domain.ITextUsecase) *Text {
	return &Text{
		textUsecase: usecase,
		curlUsecase: curlUsecase,
	}
}

func (t *Text) QueryTextByFilter(ctx context.Context, request *domain.QueryTextReq) (*domain.QueryTextResp, error) {

	// input vaild check
	if request.StartTime > request.EndTime {
		return nil, errors.New("unvaild error startTime > endTime")
	}

	filter := &domain.TextFilter{
		KeyWord:   request.KeyWord,
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
	}
	if _, err := t.textUsecase.QueryTextByFilter(ctx, filter); err != nil {
		return nil, err
	}
	texts, err := t.curlUsecase.QueryTextByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &domain.QueryTextResp{
		Data: texts,
	}, nil
}
