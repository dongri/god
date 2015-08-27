# GoUtils

```
package main

import "fmt"
import "time"
import "github.com/dongri/goutils"

func main() {

	// Detect Credit Card Type
  cardType := goutils.DetectCardType("4111111111111")
	fmt.Println(cardType)

	// Generator Random
  random := new(goutils.Random)
	random.UseNumber()
	random.UseSmallLetter()
	random.UseCapitalLetter()
	r := random.Random(20)
	fmt.Println(r)

  // Date Format
  time, _ := time.Parse("2006-01-02 15:04:05", "2015-03-04 23:29:23")
	formatedDate := goutils.DateFormat(time, "%Y年%m月%d日 %H時%M分%S秒")
  fmt.Println(formatedDate)	
}

```
