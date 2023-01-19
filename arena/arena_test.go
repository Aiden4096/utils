package arena

//import (
//	"fmt"
//	"testing"
//)
//
//func TestArena(t *testing.T) {
//	i0 := Arena().newPage(1024)
//	Arena().newPage(1024)
//	i2 := Arena().newPage(1024)
//	Arena().Close(i0)
//	fmt.Println(Arena().free)
//	Arena().newPage(1024)
//	i4 := Arena().newPage(1024)
//	Arena().Close(i4)
//	fmt.Println(Arena().free)
//	i5 := Arena().newPage(1024)
//	i6 := Arena().newPage(1024)
//	Arena().newPage(1024)
//	Arena().Close(i6)
//	Arena().Close(i2)
//	Arena().Close(i5)
//	fmt.Println(Arena().free)
//	i8 := Arena().newPage(1024)
//	i9 := Arena().newPage(1024)
//	fmt.Println(Arena().free)
//	fmt.Println(i8, i9)
//	data1 := []byte("小垃圾")
//	data2 := []byte("偶是反对")
//	Arena().Put(data1, i8, 0, uint32(len(data1)))
//	Arena().Put(data2, i9, 0, uint32(len(data2)))
//	b1 := Arena().Get(i8, 0, uint32(len(data1)))
//	b2 := Arena().Get(i9, 0, uint32(len(data2)))
//	fmt.Println(string(b1))
//	fmt.Println(string(b2))
//}
