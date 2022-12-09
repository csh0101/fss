package usecase

import (
	"context"
	"fss/internal/domain"
	"io"
	"net/http"
	"sync"
	"time"
)

var _ domain.ITextUsecase = new(curl)

var urls = []string{"http://www.baidu.com", "http://www.bilibili.com"}

type curl struct {
	processer URLProcess
}

func NewCurlUsecase(processer URLProcess) *curl {
	if processer == nil {
		return &curl{
			defaultWebClient{},
		}
	}
	return &curl{
		processer: processer,
	}
}
func (c *curl) QueryTextByFilter(ctx context.Context, request *domain.TextFilter) ([]*domain.Text, error) {
	// the number can
	ch := make(chan string, 5)
	wg := &sync.WaitGroup{}
	res := &result{}
	workerCtx, cancel := context.WithCancel(context.Background())
	once := &sync.Once{}
	cancelCtx := func() {
		once.Do(func() {
			cancel()
		})
	}
	defer cancelCtx()
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go worker(workerCtx, cancelCtx, wg, ch, c.processer, res)
	}
	for _, v := range urls {
		ch <- v
	}
	close(ch)
	wg.Wait()
	if res.err != nil {
		return nil, res.err
	}
	texts := make([]*domain.Text, 0)
	res.data.Range(func(key, value interface{}) bool {
		if val, ok := value.(string); ok {
			texts = append(texts, &domain.Text{
				Content: val,
			})
		}
		return true
	})
	return texts, nil
}

func worker(ctx context.Context, cancelCtx func(), wg *sync.WaitGroup, ch <-chan string, processer URLProcess, res *result) {
	defer wg.Done()
	var url string
LOOP:
	for {
		select {
		case url = <-ch:
			data, err := processer.DoRequest(url)
			if err != nil {
				res.err = err
				cancelCtx()
				break LOOP
			}
			res.data.Store(url, string(data))
			return
		case <-ctx.Done():
			break LOOP
		}
	}
}

type result struct {
	data sync.Map
	err  error
}

type URLProcess interface {
	DoRequest(url string) ([]byte, error)
}

func TestDo() {
}

type defaultWebClient struct {
}

func (c defaultWebClient) DoRequest(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
