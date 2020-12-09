package conf

import "sync"

var (
	C    *Config
	once sync.Once
)

func init() {
	C = &Config{
		ConnUndoQueueSize:   100,
		ConnWriteQueueSize:  100,
		FirstPackageTimeout: 5,
		ConnReadTimeout:     15,
		ConnWriteTimeout:    5,
		MaxDataPackageSize:  2048,
		MaxHeaderLen:        1024,
	}
}

// please set this before startup server
func SetConf(cfg *Config) {
	once.Do(func() {
		if cfg.ConnUndoQueueSize < 100 {
			cfg.ConnUndoQueueSize = 100
		}
		if cfg.ConnWriteQueueSize < 100 {
			cfg.ConnWriteQueueSize = 100
		}
		C = cfg
	})
}

type Config struct {
	// 单个连接未处理消息包缓存队列大小 [超过这个大小，包将丢弃，视为当前系统无法处理，最小100]
	ConnUndoQueueSize int
	// 单个连接未写入消息包队列大小 [超过这个大小，包将丢弃，视为当前系统无法处理，最小100]
	ConnWriteQueueSize int
	// 第一个包等待超市时间 (s) [默认5秒，连接上来未读到正确包，断开连接]
	FirstPackageTimeout int
	// 连接读取超时(s) [默认15秒, 超时等待时间内，请发送任何数据包，如心跳包]
	ConnReadTimeout int
	// 连接写超时(s) [默认5秒, 超时等待时间内，请发送任何数据包，如心跳包]
	ConnWriteTimeout int
	// 数据包最大限制，[默认2048]
	MaxDataPackageSize int
	// ws 最大header，[默认1024]
	MaxHeaderLen int
}