package god

import "strings"

type HyphenKind string

const (
	HANKAKU_HYPHEN HyphenKind = "-"
	ZENKAKU_HYPHEN HyphenKind = "ー"
)

func ReplaceHyphen(source string, kind HyphenKind) string {
	hyphen := string(kind)
	hyphenReplacer := strings.NewReplacer(
		// 半角ハイフン
		string(HANKAKU_HYPHEN), hyphen,
		// 全角ハイフン
		string(ZENKAKU_HYPHEN), hyphen,
		//全角ハイフン/マイナス
		"－", hyphen,
		// 小さいハイフン/マイナス
		"﹣", hyphen,
		// 全角マイナス
		"−", hyphen,
		// ハイフン
		"‐", hyphen,
		// 短いハイフン
		"⁃", hyphen,
		// 右端で折り返さないハイフン
		"‑", hyphen,
		// フィギュアダッシュ
		"‒", hyphen,
		// "n"の幅のダッシュ
		"–", hyphen,
		// "m"の幅のダッシュ
		"—", hyphen,
		// 小さい"m"の幅のダッシュ
		"﹘", hyphen,
		// 水平バー
		"―", hyphen,
		// 水平線
		"⎯", hyphen,
		// 直線
		"⏤", hyphen,
		// 飾り付きマイナス
		"˗", hyphen,
		// 太いマイナス
		"➖", hyphen,
		// 下付きマイナス
		"₋", hyphen,
		// 長音記号
		"ー", hyphen,
		// 半角長音記号
		"ｰ", hyphen,
		// アンダーライン
		"_", hyphen,
		// 全角アンダーライン
		"＿", hyphen,
		// オーバーライン
		"‾", hyphen,
		// 長音符号
		"¯", hyphen,
		// 飾り付き長音記号
		"ˉ", hyphen,
		// 全角長音記号
		"￣", hyphen,
		// 飾り付き下付き長音記号
		"ˍ", hyphen,
		// 横細罫線
		"─", hyphen,
		// 横太罫線
		"━", hyphen,
		// 細罫線用（左）
		"╴", hyphen,
		// 細罫線用（右）
		"╶", hyphen,
		// 太罫線用（左）
		"╸", hyphen,
		// 太罫線用（右）
		"╺", hyphen,
		// 下から1/8
		"▁", hyphen,
		// 下から1/4
		"▂", hyphen,
		// 下から3/8
		"▃", hyphen,
		// 下から1/2
		"▄", hyphen,
		// 下から5/8
		"▅", hyphen,
		// 下から3/4
		"▆", hyphen,
		// 下から7/8
		"▇", hyphen,
		// 下から8/8
		"█", hyphen,
		// 上から1/8
		"▔", hyphen,
		// 黒い長方形
		"▬", hyphen,
		// 漢字の「一」
		"一", hyphen,
		// 漢字の部首の「⼀」
		"⼀", hyphen,
		// 漢文の返り点の「㆒」
		"㆒", hyphen,
	)
	return hyphenReplacer.Replace(source)
}
