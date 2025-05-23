package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"web-study/service"
	"web-study/types"
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

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
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
	var req types.CreateRequest

	// body 값에 대한 바인딩 체크
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(-1, "바인딩 오류입니다.", err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(-1, "서비스 오류입니다.", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(1, "생성 성공!", nil),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get requested !")

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse(1, "성공입니다.", nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update requested !")

	var req types.UpdateRequest

	// body 값에 대한 바인딩 체크
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(-1, "바인딩 오류입니다.", err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(-1, "서비스 오류입니다.", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(1, "업데이트 성공!", nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteRequest

	// body 값에 대한 바인딩 체크
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(-1, "바인딩 오류입니다.", err.Error()),
		})
	} else if err = u.userService.Delete(req.Name); err != nil {
		u.router.failedResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(-1, "서비스 오류입니다.", err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CommonResponse{
			ApiResponse: types.NewApiResponse(1, "삭제 성공!", nil),
		})
	}
}
