package network

import "github.com/gin-gonic/gin"

// register 유틸 함수들

func (r *Network) ServerStart(port string) error {
	return r.engin.Run(port)
}

func (n *Network) requestGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.GET(path, handler...)
}

func (n *Network) requestPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.POST(path, handler...)
}

func (n *Network) requestPUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.PUT(path, handler...)
}

func (n *Network) requestDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.DELETE(path, handler...)
}

// Response 형태를 맞추기 위한 유틸 함수

func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *Network) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
