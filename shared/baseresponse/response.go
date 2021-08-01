package baseresponse

import (
	"blackboards/shared/baseerror"
	"blackboards/shared/proto"
	"net/http"
	"github.com/tal-tech/go-zero/core/logx"
)

const codeServiceUnavailable = 10001

var (
	serviceUnavailable = "服务器竟然开小差，一会儿再试试吧"
)

// Deprecated: use FormatResponseWithRequest instead
func FormatResponse(data interface{}, err error, w http.ResponseWriter) {
	if err != nil {
		codeErr, ok := baseerror.FromError(err)
		if ok {
			httpBizError(w, codeErr)
		} else {
			httpServerError(w)
		}
		logx.Errorf("%+v", err)
	} else {
		proto.HttpOk(w, data)
	}
}

func FormatResponseWithRequest(data interface{}, err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		codeErr, ok := baseerror.FromError(err)
		if ok {
			httpBizError(w, codeErr)
		} else {
			httpServerError(w)
		}
		logx.WithContext(r.Context()).Error(err)
	} else {
		proto.HttpOk(w, data)
	}
}

// Deprecated: use HttpParamErrorWithRequest instead
func HttpParamError(w http.ResponseWriter, err error) {
	logx.Error(err)
	proto.HttpParamError(w, err.Error())
}

func HttpParamErrorWithRequest(w http.ResponseWriter, r *http.Request, err error) {
	logx.WithContext(r.Context()).Error(err)
	proto.HttpParamError(w, err.Error())
}

func HttpUnauthorized(w http.ResponseWriter, err error) {
	logx.Error(err)
	proto.HttpUnauthorized(w, err.Error())
}

func httpBizError(w http.ResponseWriter, err *baseerror.CodeError) {
	proto.HttpError(w, http.StatusNotAcceptable, err.Code(), err.Desc(), err.Data())
}

func httpServerError(w http.ResponseWriter) {
	proto.HttpError(w, http.StatusInternalServerError, codeServiceUnavailable, serviceUnavailable, nil)
}
