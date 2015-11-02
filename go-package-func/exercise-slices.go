/*
实现 Pic 。它返回一个 slice 的长度 dy ，和 slice 中每个元素的长度的 8 位无符号整数 dx 。当执行这个程序，它会将整数转换为灰度（好吧，蓝度）图片进行展示。

计算每个像素的灰度值的方法由你决定；几个有意思的选择包括 (x+y)/2、x*y 和 x^y 。

（需要使用循环来分配 [][]uint8 中的每个 []uint8 。）

（使用 uint8(intValue) 在类型之间进行转换。）

*/

package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			a[i][j] = (uint8)(i &^ j)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
