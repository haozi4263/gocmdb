package response

type JSONResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewJSONResponse(code int, msg string, resulet interface{}) *JSONResponse  {
	return &JSONResponse{
		Code: code,
		Msg: msg,
		Result: resulet,
	}
}
var (
	Unauthorization = NewJSONResponse(401, "Unauthorization", nil)
	Ok = NewJSONResponse(200, "ok", "success")
	BadRequest = NewJSONResponse(400, "bad request", nil)
)