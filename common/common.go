package common

import (
	"os"
	"runtime"
	"strconv"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
)

// AdjustGoMaxProcs 设置 CPU 处理的最大线程数量
func AdjustGoMaxProcs() {

	inner.Mlogger.Debug("Adjusting bot runtime performance..")
	maxProcsEnv := os.Getenv("GOMAXPROCS")
	maxProcs := runtime.NumCPU()
	inner.Mlogger.Debugf("Number of CPU's detected:%d", maxProcs)

	if maxProcsEnv != "" {
		inner.Mlogger.Debugf("GOMAXPROCS env =%s", maxProcsEnv)
		env, err := strconv.Atoi(maxProcsEnv)
		if err != nil {
			inner.Mlogger.Debugf("Unable to convert GOMAXPROCS to int, using %d", maxProcs)
		} else {
			maxProcs = env
		}
	}
	if i := runtime.GOMAXPROCS(maxProcs); i != maxProcs {
		inner.Mlogger.Error("Go Max Procs were not set correctly.")
	}
	inner.Mlogger.Debugf("Set GOMAXPROCS to:%d", maxProcs)
}
