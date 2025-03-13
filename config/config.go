package config

import (
	"github.com/naoina/toml"
	"os"
)

// 서버의 기본적인 설정
// config를 사용하기 위해 내장 패키지를 import 해준다. (환경 파일로 toml을 사용 할 예정)
// $ go get github.com/naoina/toml
type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config) // c := &Config{} 혹은 var c *Config 와 동일함
	// c := &Config{} 형태로 선언하면 간헐적으로 go에서 타입 추론이 불가능하다는 경고창이 뜰 때가 있음.

	if file, err := os.Open(filePath); err != nil { // filePath 에 있는 파일을 open 한다.
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil { // toml 패키지로 Decode 한다.
		panic(err)
	} else {
		return c
	}
}
