package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

const (
	alpahbe = JsonContent{}.Alpahbet
)

// JsonContent ...
type JsonContent struct {
	Alphabet []string `json:alphabet` // aaa
	Number   []string `json:number`
	JA       []string `json:ja`
	KO       []string `json:ko`
}

var jsonData = JsonContent{}

func init() {
	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	json.Unmarshal(file, &jsonData)
}

func seedAndReturnRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// Returns a random part of a slice
func randomFrom(source []string) string {
	return source[seedAndReturnRandom(len(source))]
}

// Number ...
func Number(a int) int {
	// nr := 0
	// rand.Seed(time.Now().UnixNano())
	// if len(numberRange) > 1 {
	// 	nr = 1
	// 	nr = seedAndReturnRandom(numberRange[1]-numberRange[0]) + numberRange[0]
	// } else {
	// 	nr = seedAndReturnRandom(numberRange[0])
	// }
	// return nr
	return 0
}

func main() {
	fmt.Println(reflect.TypeOf(JsonContent{}))

}
