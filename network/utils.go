package network

import (
	"net/http"

	"github.com/garbage-project/trash_api.git/types/protocol"
	"github.com/gin-gonic/gin"
)

func (n *Network) registerRouter() {
	GET := func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
		return n.engine.GET(relativePath, handlers...)
	}

	POST := func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
		return n.engine.POST(relativePath, handlers...)
	}

	GET("/health-check", nil)

	POST("/stock-list", n.StockList)
	GET("/stock-detail/:id", n.StockDetail)
	GET("/reviews")
}

func (r *Network) Resp(c *gin.Context, code protocol.ResultCode, status int, resp interface{}) {
	type res struct {
		*protocol.RespHeader
		Data interface{} `json:"data"`
	}

	if status != http.StatusOK {
		c.JSON(status, &res{
			RespHeader: protocol.NewRespHeader(code, resp.(string)),
		})
	} else {
		c.JSON(status, &res{
			RespHeader: protocol.NewRespHeader(code),
			Data:       resp,
		})
	}
}
