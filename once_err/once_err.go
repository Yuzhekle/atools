package onceerr

import "sync"

// onceError is an object that will only store an error once.
type OnceError struct {
	sync.Mutex // guards following
	err        error
}

func (a *OnceError) Store(err error) {
	a.Lock()
	defer a.Unlock()
	if a.err != nil {
		return
	}
	a.err = err
}
func (a *OnceError) Load() error {
	a.Lock()
	defer a.Unlock()
	return a.err
}
