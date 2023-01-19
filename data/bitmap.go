package data

import "fmt"

const (
	bitSize = 8
)

var bitmask = []byte{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7}

type BitMap struct {
	bits     []byte
	counter  uint64
	capacity uint64
}

// 创建工厂函数
func NewBitmap(maxnum uint64) *BitMap {
	return &BitMap{bits: make([]byte, (maxnum+7)/bitSize), counter: 0, capacity: 8 * ((maxnum + 7) / bitSize)}
}

// 填入数字
func (this *BitMap) Set(num uint64) {
	byteIndex, bitPos := this.offset(num)
	// 1 左移 bitPos 位 进行 按位或 (置为 1)
	this.bits[byteIndex] |= bitmask[bitPos]
	this.counter++
}

// 清除填入的数字
func (this *BitMap) Reset(num uint64) {
	byteIndex, bitPos := this.offset(num)
	// 重置为空位 (重置为 0)
	this.bits[byteIndex] &= ^bitmask[bitPos]
	this.counter--
}

// 数字是否在位图中
func (this *BitMap) Test(num uint64) bool {
	byteIndex := num / bitSize
	if byteIndex >= uint64(len(this.bits)) {
		return false
	}
	bitPos := num % bitSize
	// 右移 bitPos 位 和 1 进行 按位与
	return !(this.bits[byteIndex]&bitmask[bitPos] == 0)
}

func (this *BitMap) offset(num uint64) (byteIndex uint64, bitPos byte) {
	byteIndex = num / bitSize // 字节索引
	if byteIndex >= uint64(len(this.bits)) {
		panic(fmt.Sprintf(" runtime error: index value %d out of range", byteIndex))
		return
	}
	bitPos = byte(num % bitSize) // bit位置
	return byteIndex, bitPos
}
