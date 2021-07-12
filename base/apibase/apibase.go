package apibase

type ApiResponse struct {
	Code int64 `json:"code"`
	Desc interface{} `json:"desc,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
//
//func FormatResponseWithRequest(data interface{}, err error, w http.ResponseWriter, r *http.Request) {
//	httpx.WriteJson(w, http.StatusOK, ApiResponse{
//		Data: data,
//	})
//}
//
//func HttpError(w http.ResponseWriter, httpCode int, appCode int64, desc string, message interface{}) {
//	httpx.WriteJson(w, httpCode, ApiResponse{
//		Code:    appCode,
//		Desc:    desc,
//		Data: message,
//	})
//}
