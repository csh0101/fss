package domain

import "context"

type Text struct {
	CreateTime uint64 `bson:"create_time"`
	Content    string `bson:"content"`
	Name       string `bson:"name"`
}

type TextFilter struct {
	KeyWord   []string
	StartTime uint64
	EndTime   uint64
}

type QueryTextResp struct {
	Data []*Text
}

type QueryTextReq struct {
	KeyWord   []string `json:"key_words"`
	StartTime uint64   `json:"start_time"`
	EndTime   uint64   `json:"end_time"`
}

type TextHTTPService interface {
	QueryTextByFilter(ctx context.Context, reqeust *QueryTextReq) (*QueryTextResp, error)
}

type ITextUsecase interface {
	QueryTextByFilter(ctx context.Context, filter *TextFilter) ([]*Text, error)
}

type ITextRepo interface {
	QueryTextByFilter(ctx context.Context, filter *TextFilter) ([]*Text, error)
}
