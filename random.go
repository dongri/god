package utils

import (
	"math/rand"
	"reflect"
	"time"
)

const (
	number        = "1234567890"
	smallLetter   = "abcdefghijklmnopqrstuvwxyz"
	capitalLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	character     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	hiragana      = "ぁあぃいぅうぇえぉおかがきぎくぐけげこごさざしじすずせぜそぞただちぢっつづてでとどなにぬねのはばぱひびぴふぶぷへべぺほぼぽまみむめもゃやゅゆょよらりるれろゎわゐゑをん"
	katakana      = "ァアィイゥウェエォオカガキギクグケゲコゴサザシジスズセゼソゾタダチヂッツヅテデトドナニヌネノハバパヒビピフブプヘベペホボポマミムメモャヤュユョヨラリルレロヮワヰヱヲンヴヵヶ"
	hangul        = "ㄱㄴㄷㄹㅁㅂㅅㅇㅈㅊㅋㅌㅍㅎㄲㄸㅃㅆㅉㅏㅑㅓㅕㅗㅛㅜㅠㅡㅣㅐㅒㅔㅖㅘㅙㅚㅝㅞㅟㅢ"
	shuzi         = "一二三四五六七八九十零"
)

// Random ...
type Random struct {
	Number        string
	SmallLetter   string
	CapitalLetter string
	Character     string
	Hiragana      string
	Katakana      string
	Hangul        string
	Shuzi         string
}

// UseNumber ...
func (t *Random) UseNumber() {
	t.Number = number
}

// UseSmallLetter ...
func (t *Random) UseSmallLetter() {
	t.SmallLetter = smallLetter
}

// UseCapitalLetter ...
func (t *Random) UseCapitalLetter() {
	t.CapitalLetter = capitalLetter
}

// UseCharacter ...
func (t *Random) UseCharacter() {
	t.Character = character
}

// UseHiragana ...
func (t *Random) UseHiragana() {
	t.Hiragana = hiragana
}

// UseKatakana ...
func (t *Random) UseKatakana() {
	t.Katakana = katakana
}

// UseHangul ...
func (t *Random) UseHangul() {
	t.Hangul = hangul
}

// UseShuzi ...
func (t *Random) UseShuzi() {
	t.Shuzi = shuzi
}

// Random ...
func (t *Random) Random(n int) string {
	letters := []rune(t.generatorData())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[randomInt(len(letters))]
	}
	return string(b)
}

func (t *Random) generatorData() string {
	st := reflect.ValueOf(*t)
	data := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		data += field.String()
	}
	return data
}

func randomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
