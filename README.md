# Go API Project

이 프로젝트는 Go를 이용해 제작되었습니다.

gin 프레임워크와 gRPC, Consumer 등으로 호출되는 API를 제공합니다.

**개인 프로젝트로 제작되어 직접적인 프로젝트 실행이 불가능합니다.**


## Description

이 프로젝트는 다음과 같은 API 및 기능을 제공합니다:

| API        | Method | Description |
|------------|--------|-------------|
| Notification | Email  | 이메일 전송      |
|            | SMS    | 문자 전송       |
|            | Push   | 푸시 전송       |
|            | Slack  | 슬랙 메시지 전송  |
| Account | Auth   | 인증 API       |
|            | User   | 사용자 API      |
| Location   | IP     | IP 주소 조회    |
|            | Geo    | Geo 정보 조회   |
| Socket     | Websocket | 웹소켓 통신 |
|            | SlackSocketMode | 슬랙 소켓 통신 |
| Content | Streaming | 비디오 스트리밍 |
| AI | LLM | ChatGPT API |
| 기타         | SSE    | SSE 통신       |


## Installation

프로젝트 다운로드 및 패키지를 설치하셔야 사용이 가능합니다.

```bash
git clone https://github.com/lee-lou2/go.git
cd go
go mod tidy
```

## Usage

프로젝트는 아래 명령어로 실행할 수 있습니다.

```bash
# .env 파일 생성
# .env.sample 정보를 입력
go run main.go
```
