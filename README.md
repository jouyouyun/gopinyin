Convert hans to pinyin
=====================

`func HansToPinyin(hans string) (pinyinList []string)`

>`hans: 带汉字的字符串`

>`pinyinList: 返回的结果列表，因为考虑到多音字，结果会有多种组合`

如果 `hans` 中包含非汉字，则原样返回。

Usage example:
-------------

```go
package main

import (
	"github.com/jouyouyun/gopinyin"
	"fmt"
)

func main () {
	ret := gopinyin.HansToPinyin("你好")
	fmt.Println(ret)

	ret = gopinyin.HansToPinyin("重庆")
	fmt.Println(ret)
}
```

Output:

```shell
[nihao]
[chongqin zhongqin]
```
