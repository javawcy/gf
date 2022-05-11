package ghttp

//CEBaseResponse baseResp for rest
type CEBaseResponse struct {
	Ret      int    `json:"ret" dc:"Error code"`
	Msg      string `json:"msg" dc:"Error message"`
	DataHash string `json:"dataHash" dc:"data cache key"`
}

//CEMiddlewareHandlerResponse custom res middleware
func CEMiddlewareHandlerResponse(r *Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	// var (
	// 	msg  string
	// 	ctx  = r.Context()
	// 	err  = r.GetError()
	// 	res  = r.GetHandlerResponse()
	// 	code = gerror.Code(err)
	// )

}
