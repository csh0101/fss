package http

import (
	"context"
	v1 "fss/api/product/app/v1"
	"fss/internal/database"
	"fss/internal/domain"
	"fss/internal/server/repo"
	"fss/internal/server/service"
	"fss/internal/server/usecase"
	"fss/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var _ v1.TextHTTPServer = new(Server)

const metricsNamespace = "fss"

var (
	ReqeustFailedTotalVec = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Name:      "reqeust_fail_cnt",
		Help:      "The total number of request_fail_cnt",
	}, []string{"status"})
)

func init() {
	prometheus.MustRegister(ReqeustFailedTotalVec)
}

type Server struct {
	textService *service.Text
}

// todo repleace it with DI ? maybe its not need. the design mind of sprintboot maybe fix in go
// InitServer
func InitServer(uri string) *Server {
	db := database.InitMonger(uri)
	textRepo := repo.NewTextRepo(db)
	textUsecase := usecase.NewTextUsecase(textRepo)
	curlUsecase := usecase.NewCurlUsecase(nil)
	textService := service.NewTextService(textUsecase, curlUsecase)
	return &Server{
		textService: textService,
	}
}

// Run must call after InitServer
func (s *Server) Run(addr string) error {
	app := echo.New()

	// register middleware
	// app.Use(middleware.Logger())
	// app.Use(middleware.Recover())
	// app.Use(middleware.CORS())
	// app.Use(middleware.CSRF())
	// Register Router
	Register(app, s)

	// Start HTTP Server Logic
	srv := &http.Server{
		Addr:    addr,
		Handler: app,
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			log.Println("srv close exit...")
		}
	}()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":9999", nil))
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	//todo replace good log
	log.Println(time.Now(), "fss server start succsffully")
	wg.Wait()
	return nil
}

func (s *Server) QueryTextByFilter(ctx context.Context, reqeust *http.Request) (*v1.HTTPResponse, error) {
	response := &v1.HTTPResponse{}
	filter := &domain.QueryTextReq{}
	var err error
	{
		response.Code = 10001
		query := reqeust.URL.Query()
		if v, ok := query["key_word"]; ok {
			filter.KeyWord = v
		}
		if v, ok := query["start_time"]; ok {
			filter.StartTime, err = utils.Str2Uint64(v[0])
		}
		if v, ok := query["end_time"]; ok {
			filter.EndTime, err = utils.Str2Uint64(v[0])
		}
		if err != nil {
			response.Message = err.Error()
			ReqeustFailedTotalVec.WithLabelValues("failed").Inc()
			goto FINISHED
		}
	}

	{
		response.Code = 10000
		var resp *domain.QueryTextResp
		resp, err = s.textService.QueryTextByFilter(ctx, filter)
		if err != nil {
			response.Code = 10001
			response.Message = err.Error()
			ReqeustFailedTotalVec.WithLabelValues("failed").Inc()
			goto FINISHED
		}
		response.Data = resp
		response.Message = "success"
		ReqeustFailedTotalVec.WithLabelValues("successful").Inc()
	}
FINISHED:
	return response, err
}
