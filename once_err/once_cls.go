package onceerr

import "sync"

type SingleStop struct {
	once sync.Once
	stop chan struct{}
}

func (s *SingleStop) Stop() {
	s.once.Do(func() {
		close(s.stop)
	})
}
