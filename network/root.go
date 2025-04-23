package network

import "github.com/gin-gonic/gin"

// api를 만든다고 했을 때, 라우터 설정을 하는 곳
// 다양한 API에 대한 root를 관리한다.

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engin: gin.New(),
	}

	// network 에 user router 추가
	newUserRouter(r)

	return r
}
