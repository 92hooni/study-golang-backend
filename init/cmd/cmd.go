package cmd

import (
	"fmt"
	"web-study/config"
)

type Cmd struct {
	// 다른 network, repository, service, types, config 를 구성 할 예정
	// cmd는 main에서 호출된다.

	config *config.Config // config 를 선언
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	fmt.Println(c.config.Server.Port)

	return c
}
