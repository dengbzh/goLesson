/*
一个常见模式是 io.Reader 包裹另一个 io.Reader，然后通过某种形式修改数据流。

例如，gzip.NewReader 函数接受 io.Reader（压缩的数据流）并且返回同样实现了 io.Reader 的 *gzip.Reader（解压缩后的数据流）。

编写一个实现了 io.Reader 的 rot13Reader， 并从一个 io.Reader 读取， 利用 rot13 代换密码对数据流进行修改。

已经帮你构造了 rot13Reader 类型。 通过实现 Read 方法使其匹配 io.Reader。
*/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (v rot13Reader) Read(b []byte) (int, error) {

	a := make([]byte, 50)

	i, err := v.r.Read(a)

	numa := byte('a')
	numz := byte('z')
	numA := byte('A')
	numZ := byte('Z')
	var tmp byte
	for j := 0; j < i; j++ {
		if a[j] >= numa && a[j] <= numz {
			tmp = a[j] - numa + 13
			if tmp >= 26 {
				tmp -= 26
			}
			tmp += numa
			b[j] = tmp
			fmt.Println(a[j]-numa, b[j]-numa)
		} else if a[j] >= numA && a[j] <= numZ {
			tmp = a[j] - numA + 13
			if tmp >= 26 {
				tmp -= 26
			}
			tmp += numA
			b[j] = tmp
			fmt.Println(a[j]-numA, b[j]-numA)
		} else {
			b[j] = a[j]
		}

	}
	return i, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
