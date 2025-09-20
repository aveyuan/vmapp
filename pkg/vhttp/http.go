package vhttp

type Data struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Page *Page       `json:"meta,omitempty"`
}

type DataOptions func(*Data)

func WithPage(page *Page) DataOptions {
	return func(data *Data) {
		data.Page = page
	}
}

func WithMsg(msg string) DataOptions {
	return func(data *Data) {
		data.Msg = msg
	}
}

func WithCode(code int) DataOptions {
	return func(data *Data) {
		data.Code = code
	}
}
