package shared

//func FormatResponseWithRequest(data interface{}, err error, w http.ResponseWriter, r *http.Request) {
//	if err != nil {
//		codeErr, ok := baseerror.FromError(err)
//		if ok {
//			httpBizError(w, codeErr)
//		} else {
//			httpServerError(w)
//		}
//		logx.WithContext(r.Context()).Error(err)
//	} else {
//		proto.HttpOk(w, data)
//	}
//}

