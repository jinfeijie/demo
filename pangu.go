package pangu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goxtools/watcher"
	"github.com/jinfeijie/pangu/internal/constant"
	"github.com/jinfeijie/pangu/internal/infra"
	handler2 "github.com/jinfeijie/pangu/internal/infra/handler"
	"github.com/jinfeijie/pangu/internal/infra/handler/configcenter"
	"github.com/jinfeijie/pangu/pkg"
	"github.com/jinfeijie/pangu/pkg/middleware/trace"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type PanGu struct {
	httpServer *gin.Engine
}

func New() *PanGu {
	return &PanGu{
		httpServer: gin.New(),
	}
}

func (pg *PanGu) HttpRoute(f func(eng *gin.Engine)) {
	pg.httpServer.
		Use(gin.Recovery()).
		Use(trace.TraceId()) // 提早注入通用中间件
	f(pg.httpServer)
}

func (pg *PanGu) GrpcRoute() {

}

func (pg *PanGu) Dubbo3Route() {

}

func (pg *PanGu) JobRoute() {

}

func (pg *PanGu) Start() {
	// 这里注意夯住，等待下线
	// 可以切换成监听下线的代码
	infra.Handlers(constant.GroupInit)
	infra.Handlers(constant.GroupBefore)
	go watcher.NewWatcher(3, time.Second*10, time.Second*5).On(func(args ...interface{}) {
		pg.httpServer.Run(fmt.Sprintf(":%d", configcenter.GetConfig().Server.Port))
	})
	infra.Handlers(constant.GroupAfter)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("Program Exit...", s)
			//GracefullExit() // 优雅退出

			// 如果有优雅关机的毁掉，shutdown方法会优雅一些
			infra.Handlers(constant.GroupShutDown)

			time.Sleep(time.Second * 10)
			os.Exit(1)

		case syscall.SIGUSR1:
			fmt.Println("usr1 signal", s)
		case syscall.SIGUSR2:
			fmt.Println("usr2 signal", s)
		default:
			fmt.Println("other signal", s)
		}
	}
}

/* 以下提供注册事件 开始 */

func (pg *PanGu) HandlerInit(handler ...pkg.Handler) {
	for _, h := range handler {
		handler2.Register(h.Name(), h, constant.GroupInit)
	}
}

func (pg *PanGu) HandlerBefore(handler ...pkg.Handler) {
	for _, h := range handler {
		handler2.Register(h.Name(), h, constant.GroupBefore)
	}
}

func (pg *PanGu) HandlerAfter(handler ...pkg.Handler) {
	for _, h := range handler {
		handler2.Register(h.Name(), h, constant.GroupAfter)
	}
}

func (pg *PanGu) HandlerShutDown(handler ...pkg.Handler) {
	for _, h := range handler {
		handler2.Register(h.Name(), h, constant.GroupShutDown)
	}
}

/* 以上提供注册事件 结束 */
