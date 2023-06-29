package main

import (
	"github.com/lee-lou2/go/api/restapi"
	"github.com/lee-lou2/go/configs"
)

func main() {
	// 환경 변수 조회
	configs.LoadEnvironments()

	// App 실행
	{

	}
	restapi.Run()
}
