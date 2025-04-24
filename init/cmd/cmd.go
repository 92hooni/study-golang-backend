package cmd

import (
	"web-study/config"
	"web-study/network"
	"web-study/repository"
	"web-study/service"
)

type Cmd struct {
	// 다른 network, repository, service, types, config 를 구성 할 예정
	// cmd는 main에서 호출된다.

	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
