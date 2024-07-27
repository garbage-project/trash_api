package network

import (
	"net/http"

	. "github.com/garbage-project/trash_api.git/types"
	. "github.com/garbage-project/trash_api.git/types/protocol"
	"github.com/gin-gonic/gin"
)

func (n *Network) StockList(c *gin.Context) {
	var req StockListReq

	if err := c.ShouldBindJSON(&req); err != nil {
		n.Resp(c, FailedToBind, http.StatusBadRequest, err.Error())
	} else if res, err := n.service.StockList(req); err != nil {
		n.Resp(c, Failed, http.StatusInternalServerError, err.Error())
	} else {
		n.Resp(c, Success, http.StatusOK, res)
	}
}

func (n *Network) StockDetail(c *gin.Context) {
	var req StokeIdReq

	if err := c.ShouldBindUri(&req); err != nil {
		n.Resp(c, FailedToBind, http.StatusBadRequest, err.Error())
	} else if res, err := n.service.StockDetail(req.ID); err != nil {
		n.Resp(c, Failed, http.StatusInternalServerError, err.Error())
	} else {
		n.Resp(c, Success, http.StatusOK, res)
	}
}

func (n *Network) Reviews(c *gin.Context) {
	var req StokeIdReq

	if err := c.ShouldBindUri(&req); err != nil {
		n.Resp(c, FailedToBind, http.StatusBadRequest, err.Error())
	} else if res, err := n.service.Reviews(req.ID); err != nil {
		n.Resp(c, Failed, http.StatusInternalServerError, err.Error())
	} else {
		n.Resp(c, Success, http.StatusOK, res)
	}
}
