package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	// 라우터는 한번만 실행되야 하므로 sync.Once 설정 추가
	// network Root에서 2번 호출되어도 1번만 실행된다.
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		// Go 기본 모듈의 http.handleFunc("/", ) 과 동일한 액션
		/* 아래처럼 선언해도 되지만, network 에 함수로 분리해 호출을 간소화 할 수 있다.
		userRouterInstance.router.engin.POST("/", userRouterInstance.create)
		userRouterInstance.router.engin.GET("/", userRouterInstance.get)
		userRouterInstance.router.engin.PUT("/", userRouterInstance.update)
		userRouterInstance.router.engin.DELETE("/", userRouterInstance.delete)*/

		router.requestPOST("/", userRouterInstance.create)
		router.requestGET("/", userRouterInstance.get)
		router.requestPUT("/", userRouterInstance.update)
		router.requestDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

/*
TIP: Go 에서의 접근제어자는 함수명을 보고 판단한다.
(함수명의 시작이 대/소문자 인지를 보고 판단)
*/
func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create requested !")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get requested !")

	u.router.okResponse(c, "테스트 입니다.")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update requested !")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete requested !")
}
