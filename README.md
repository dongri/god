# go-beer
:beer: :beer: :beer: :beer: :beer: :beer: :beer: :beer: :beer: :beer: 

:beers: :beers: :beers: :beers: :beers: :beers: :beers: :beers: :beers: :beers:

# go-timeformatter

```
package main

import (
	"fmt"
	"time"

	"github.com/dongri/go-timeformatter"
)

func main() {
	time := TimeFormatter.Format(time.Now(), "%Y年%m月%d日 %H時%M分%S秒")
	fmt.Println(time) // 2015年03月04日 22時56分23秒
}
```
