package queue

import (
	"github.com/zeromicro/go-queue/kq"
)

func ListenKq(handlers []kq.ConsumeHandle, conf kq.KqConf) {
	for _, handler := range handlers {
		go func() {
			q := kq.MustNewQueue(conf, kq.WithHandle(handler))

			defer q.Stop()
			q.Start()
		}()
	}
}
