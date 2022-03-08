package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseInfo 通用返回
type ResponseInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, err error, code int, data interface{}) {
	res := &ResponseInfo{
		Code: Success,
		Data: data,
	}
	if err != nil {
		res.Code = code
		res.Msg = err.Error()
	}
	//ret, _ := json.Marshal(res)
	//c.Writer.Write(ret)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, res)
	return
}

func SimpleResponse(c *gin.Context, res *ResponseInfo) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, res)
	return
}
