package usecase_test

import (
	"context"
	"fmt"
	"fss/internal/server/usecase"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCurlUsecase(t *testing.T) {
	c := usecase.NewCurlUsecase(nil)
	data, err := c.QueryTextByFilter(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(data))
}

func TestX(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	assert.Equal(t, nil, err)
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	assert.Equal(t, nil, err)
	_, err = io.ReadAll(resp.Body)
	assert.Nil(t, nil, err)

}
