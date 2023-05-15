package restapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	accountRoutes "github.com/lee-lou2/go/apps/account/api/restapi/routes"
	locationRoutes "github.com/lee-lou2/go/apps/location/api/restapi/routes"
	notifyRoutes "github.com/lee-lou2/go/apps/notify/api/restapi/routes"
	utilRoutes "github.com/lee-lou2/go/apps/util/api/restapi/routes"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"os"
)

// Run API 서버 실행
func Run() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	// 모니터링
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	// 미들웨어
	DefaultMiddleware(r)

	// 라우터 설정
	r.GET("/ping", HealthCheckHandler)
	v1 := r.Group("/v1")
	{
		// 테스트
		test := v1.Group("/test")
		{
			test.GET("/sleep", TestSleepHandler)
		}
		// 알림
		notifyRoutes.V1Routes(v1)
		// 지역
		locationRoutes.V1Routes(v1)
		// 계정
		accountRoutes.V1Routes(v1)
		// 유틸
		utilRoutes.V1Routes(v1)
	}

	// 서버 포트 지정
	port := fmt.Sprintf(":%s", os.Getenv("HTTP_API_PORT"))
	_ = r.Run(port)
}
