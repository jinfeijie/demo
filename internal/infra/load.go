package infra

import (
	"github.com/jinfeijie/pangu/internal/constant"
	"github.com/jinfeijie/pangu/internal/infra/handler"
	"log"
	"time"
)

// Handlers 加载阶段函数
func Handlers(groupStatus constant.SortType) {
	if handlers, exist := handler.Chain[groupStatus]; exist {
		for _, handler := range handlers {
			execCostStart := time.Now()
			handler.Serve()
			log.Printf("配置 【%s】已初始化完成，耗时【%s】\n", handler.Name(), time.Since(execCostStart))
		}
	}
}
