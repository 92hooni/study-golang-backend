package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// binding 할 수 있는 방법은 여러개가 있으며, gin framework git을 참조
// 링크 :
type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type UpdateRequest struct {
	Name       string `json:"name" binding:"required"`
	UpdatedAge int64  `json:"updatedAge" binding:"required"`
}
type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}
type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"userList"`
}
type CommonResponse struct {
	*ApiResponse
}
