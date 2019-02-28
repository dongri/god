# :ramen: gou

### Installation

```
$ go get github.com/dongri/gou
```

### Detect Credit Card Type
```go
package main
import (
	"fmt"
	"github.com/dongri/gou"
)
func main() {
	cardType1 := gou.DetectCardType("4111111111111")
	fmt.Println(cardType1) // Visa
	cardType2 := gou.DetectCardType("5555555555554444")
	fmt.Println(cardType2) // Master
}
```

### Date Format
```go
package main
import (
	"fmt"
	"time"
	"github.com/dongri/gou"
)
func main() {
	time, _ := time.Parse("2006-01-02 15:04:05", "2015-03-04 23:29:23")
	formatedDate := gou.DateFormat(time, "%Y年%m月%d日 %H時%M分%S秒")
	fmt.Println(formatedDate) // 2015年03月04日 23時29分23秒
}
```

### Random
```go
package main
import (
	"fmt"
	"github.com/dongri/gou"
)
func main() {
	random := new(gou.Random)
	random.UseNumber()
	random.UseSmallLetter()
	random.UseCapitalLetter()
	r := random.Random(20)
	fmt.Println(r) // 6Zyk2vOQYNIKmbIa6BcH
}
```

### Sha1
```go
package main
import (
	"fmt"
	"github.com/dongri/gou"
)
func main() {
	s := gou.Sha512Sum512("hello"))
	fmt.Println(s)
}
```

### Cipher
```go
package main
import (
	"fmt"
	"github.com/dongri/gou"
)
func main() {
	encryptedText, err := gou.EncryptString("passwordpasswordpasswordpassword", "hoge")
	fmt.Println(err, encryptedText)
}
```

### Beacon
```go
package main
import (
	"fmt"
	"github.com/dongri/gou"
)
func main() {
	major, minor := gou.PackBeacon(1,2,8)
	fmt.Println(major, minor)

	major, minor := gou.UnpackBeacon(0, 258,8)
	fmt.Println(major, minor)
}
```

# License

The MIT License (MIT)

Copyright (c) 2015 Dongri Jin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
