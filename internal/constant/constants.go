package constant

type SortType int

const (
	GroupInit     = SortType(iota) // 服务初始化
	GroupBefore                    // 服务启动前
	GroupAfter                     // 服务启动后
	GroupShutDown                  // 关机信号
)
