package middleware

import (
	"errors"
	"mime"
	"net/http"
	"vmapp/pkg/vhttp"

	"github.com/gogf/gf/v2/net/ghttp"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code     int         `json:"code"    dc:"Error code"`
	Msg      string      `json:"msg"  dc:"Error message"`
	Data     interface{} `json:"data"    dc:"Result data"`
	Metadata interface{} `json:"metadata"    dc:"Result metadata for certain request according API definition"`
}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

var (
	// streamContentType is the content types for stream response.
	streamContentType = []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace}
)

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0 {
		return
	}

	// It does not output common response content if it is stream response.
	mediaType, _, _ := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))
	for _, ct := range streamContentType {
		if mediaType == ct {
			return
		}
	}

	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	if err != nil {
		var errData vhttp.ErrData
		d := &DefaultHandlerResponse{}
		if errors.As(err, &errData) {
			d.Code = errData.Code
			d.Msg = errData.Msg
			d.Metadata = errData.Metadata
		} else {
			d.Code = http.StatusInternalServerError
			d.Msg = err.Error()
		}
		r.Response.WriteJson(d)
		return
	}

	r.Response.WriteJson(res)

}
