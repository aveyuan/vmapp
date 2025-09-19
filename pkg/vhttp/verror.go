package vhttp

type CustomErrorOptions func(*ErrData)

type ErrData struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`	
	Reason   error  `json:"reason"`
	StatusCode *int `json:"statusCode"`
	Metadata interface{} `json:"metadata"`
}

func (t ErrData) Error() string {
	if t.Reason != nil {
		return t.Msg + ":" + t.Reason.Error()
	}
	return t.Msg
}

func NewError(code int, message string, cus ...CustomErrorOptions) error {
	e := ErrData{
		Code: code,
		Msg:  message,
	}
	for _, v := range cus {
		v(&e)
	}

	return e
}

// WithMetadata
func WithMetadata(Metadata interface{}) CustomErrorOptions {
	return func(ErrData *ErrData) {
		ErrData.Metadata = Metadata
	}
}

func WithReason(reason error) CustomErrorOptions {
	return func(ErrData *ErrData) {
		ErrData.Reason = reason
	}
}



func WithStatusCode(statusCode int) CustomErrorOptions {
	return func(ErrData *ErrData) {
		ErrData.StatusCode = &statusCode
	}
}