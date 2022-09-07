package draw

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/fogleman/gg"
)

func WordToPic(msg string) bool {

	// 绘制底图

	wordSize := utf8.RuneCountInString(msg)
	strSize := len(msg)

	wight := strSize*100 - (strSize-wordSize)*2*50

	dc := gg.NewContext(wight, 100)
	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace("IPix.ttf", 90); err != nil {
		fmt.Println(err)
		return false
	}

	floatwight, err := strconv.ParseFloat(strconv.Itoa(wight/2), 64)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(wight)
	fmt.Println(floatwight)

	dc.DrawStringAnchored(msg, floatwight, 50, 0.5, 0.5)
	dc.SavePNG("out.png")

	return true
}
