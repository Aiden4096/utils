package arena

import (
	"sync"
	"utils/data"
)

var (
	lock     = sync.Mutex{}
	instance *arena
)

type (
	Page []byte
)
type arena struct {
	free data.BitMap
	buf  []Page
}

func Arena() *arena {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &arena{
				//free: make(data.Bitmap, 1),
				buf: make([]Page, 1),
			}
			lock.Unlock()
		}
	}
	return instance
}
func (a *arena) newPage(cap uint32) int {
	return 0
}
func (a *arena) Put(data []byte, bucket int, offset, sz uint32) {
	buf := a.buf[bucket][int(offset) : int(offset)+int(sz)]
	copy(buf, data)
}
func (a *arena) Get(bucket int, offset, sz uint32) []byte {
	data := a.buf[bucket][offset : offset+sz]
	return data
}
func (a *arena) Close(bucket int) {

}
func (a *arena) Codec(func(buf []byte, index ...int) any) {

}
