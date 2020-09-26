package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"strconv"
	"time"
)

/*andlabs这个gui库缺点太大，最不能容忍的是组件缺少，窗口属性缺少，就连颜色什么的也没有，可以说是相当鸡肋了，可以弃用*/

func main() {
	err := ui.Main(func() {
		box := ui.NewVerticalBox() // 生成：垂直容器
		weekday := time.Now().Weekday()
		greet := ui.NewLabel("阖家团圆")
		box.Append(greet, false)
		content := "今天是" + weekday.String() + "，要开心哦~"
		greet.SetText(content)
		box.Append(ui.NewHorizontalSeparator(), false) // 分割线
		// 日期时间
		box.Append(ui.NewLabel("起始时间"), false)
		t1 := ui.NewDateTimePicker()
		box.Append(t1, false)
		box.Append(ui.NewLabel("终止时间"), false)
		t2 := ui.NewDateTimePicker()
		box.Append(t2, false)

		countButton := ui.NewButton("计算时差")
		result := ui.NewEntry()
		result.SetReadOnly(true)
		box.Append(countButton, false)
		box.Append(result, false)
		countButton.OnClicked(func(*ui.Button) {
			sumM := t2.Time().Sub(t1.Time()).Minutes()
			res := strconv.FormatFloat(sumM, 'f', -1, 64)
			res = res + string("分钟")
			result.SetText(res)
		})
		// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
		window := ui.NewWindow(`Time Count Tools`, 250, 150, true)
		window.SetChild(box)                     // 窗口容器绑定
		window.OnClosing(func(*ui.Window) bool { // 设置：窗口关闭时
			ui.Quit() // 窗体关闭
			return true
		})
		// 窗体显示
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
