package draw

import (
	"unicode/utf8"

	"github.com/fogleman/gg"
)

func WordToPic(msg string) bool {

	// 绘制底图

	wordSize := utf8.RuneCountInString(msg)
	strSize := len(msg)

	dc := gg.NewContext(strSize*100-(strSize-wordSize)*2*50, 100)
	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace("IPix.ttf", 100); err != nil {
		panic(err)
	}

	dc.DrawString(msg, 10, 75)
	dc.SavePNG("out.png")

	return true
}
