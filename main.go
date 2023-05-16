package main

import (
	"github.com/lee-lou2/go/api/restapi"
	"github.com/lee-lou2/go/apps/socket"
	"github.com/lee-lou2/go/configs"
)

func main() {
	// 환경 변수 조회
	configs.LoadEnvironments()

	// App 실행
	{
		// 소켓 모드 추가
		socket.Run()
	}
	restapi.Run()
}
