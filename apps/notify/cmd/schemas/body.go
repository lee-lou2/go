package schemas

// PushRequest 모델 생성
type PushRequest struct {
	FcmTokens []string `json:"fmc_tokens"`
	Title     string   `json:"title"`
	Message   string   `json:"message"`
}

// EmailRequest 이메일 요청
type EmailRequest struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

// SMSRequest 문자 요청
type SMSRequest struct {
	Phones  []string `json:"phones"`
	Message string   `json:"message"`
}

// SlackRequest 문자 요청
type SlackRequest struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

// RequestTypes 요청 타입
type RequestTypes interface {
	PushRequest | EmailRequest | SMSRequest | SlackRequest
}

// Scheduler 스케쥴러
type Scheduler struct {
	// SchedulerType 0 : 즉시, 1 : 일회성, 2 : 반복
	SchedulerType int `json:"scheduler_type"`
	// OneTime
	OneTime string `json:"one_time"`
	// Interval
	Interval struct {
		StartTime string `json:"start_time"`
		Count     int    `json:"count"`
		Duration  uint64 `json:"duration"`
	} `json:"interval"`
}

// RequestBody 요청 바디
type RequestBody[T RequestTypes] struct {
	RequestID   int       `json:"request_id"`
	Data        T         `json:"data"`
	CallbackUri string    `json:"callback_uri"`
	Scheduler   Scheduler `json:"scheduler"`
}
