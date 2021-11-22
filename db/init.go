package db

import (
	"fmt"
	"sake/tools/logging"
)

// 这里是有不同的log系统
// 用于不同的声明
var (
	log logging.Logger
	// Views is a map to store views created automatically.
	// It will be processed by the views package and added to the views registry.
	// Views map[*Model][]string
)

func init() {

	// 这个地方初始化
	// 这里是不是要初始化一个log
	logging.Initialize()

	log = logging.GetLogger("models")

	fmt.Println("123123123")

	// DB drivers
	adapters = make(map[string]Dialect)
	RegisterDialect("pgx", new(postgres))
}
