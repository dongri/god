package god

import "strings"

func ReplaceWaveDash(source string) string {
	dash := "~"
	waveDashReplacer := strings.NewReplacer(
		// 波ダッシュ
		"〜", dash,
		// 半角チルダ
		"~", dash,
		// SMALL TILDE
		"˜", dash,
		// ギリシャ語
		"῀", dash,
		// SWUNG DASH
		"⁓", dash,
		// 数学記号
		"∼", dash,
		// SINE WAVE
		"∿", dash,
		// WAVY DASH
		"〰", dash,
		// 全角チルダ
		"～", dash,
	)
	return waveDashReplacer.Replace(source)
}
