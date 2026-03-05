package executor

import (
	"gameSrv/pkg/gopool"

	"github.com/spf13/viper"
)

var AsyncNetMsgExecutor *gopool.WorkerPool

var AsyncDBExecutor *gopool.WorkerPool

func InitExecutorWithConf() {
	AsyncNetMsgExecutor = gopool.StartNewWorkerPool(viper.GetInt("executor.net.workCount"), viper.GetInt("executor.net.queSize"))
	AsyncDBExecutor = gopool.StartNewWorkerPool(viper.GetInt("executor.db.workCount"), viper.GetInt("executor.db.queSize"))
}
