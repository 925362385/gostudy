# JSON and Md5

---
## JSON    go语言内置的encoding/json标准库

**性能更优的ffJson**:

github.com/pquerna/ffjson/ffjson

## go语言内置了一个MD5的包 ‘crypto/md5’标准库

```go
package main

import (
	"crypto/md5"
	"fmt")
func main(){
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("wang"))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n",Result)
}
```
