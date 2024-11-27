package requestlog

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinfeijie/pangu/pkg/middleware/log"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
	"unsafe"
)

type ResponseWriterWrapper struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w ResponseWriterWrapper) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriterWrapper) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func RequestLog(ctx *gin.Context) {

	var (
		requestBody []byte
		err         error
		s           = time.Now()
	)
	if requestBody, err = ctx.GetRawData(); err != nil {
		log.Info(ctx.Request.Context(), "[middleware][RequestLog]获取请求数据异常",
			zap.String("err", err.Error()))
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	reqData := ""
	if ctx.Request.Method == http.MethodPost {
		reqData = string(requestBody)
	} else if ctx.Request.Method == http.MethodGet {
		reqData = ctx.Request.URL.Query().Encode()
	}

	request := map[string]interface{}{
		"RequestHeader":   ctx.Request.Header,
		"RequestMethod":   ctx.Request.Method,
		"RequestClientIP": ctx.ClientIP(),
		"RequestUrl":      ctx.Request.URL.Path,
		"RequestQuery":    ctx.Request.URL.Query().Encode(),
		"RequestBody":     reqData,
		"RequestProto":    ctx.Request.Proto,
		"RequestHost":     ctx.Request.Host,
		"RequestDataLen":  ctx.Request.ContentLength,
		"RequestCookies":  ctx.Request.Cookies(),
	}

	blw := &ResponseWriterWrapper{Body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	ctx.Next()

	log.Info(ctx.Request.Context(), "[middleware][RequestLog]请求日志",
		zap.String("request", ToStr(request)),
		zap.String("resp", blw.Body.String()),
		zap.Int64("costMs", time.Now().Sub(s).Milliseconds()),
	)
}

func ToStr(obj interface{}) string {
	js, _ := json.Marshal(obj)
	return Bytes2str(js)
}

func Bytes2str(bt []byte) string {
	return *(*string)(unsafe.Pointer(&bt))
}
