package cmd

type Cmd struct {
	// 다른 network, repository, service, types, config 를 구성 할 예정
	// cmd는 main에서 호출된다.

}

func newCmd() *Cmd {
	c := &Cmd{}
	return c
}
