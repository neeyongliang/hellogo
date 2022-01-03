package type_embbeding

import "fmt"

// Uber go 编程规范：避免在公共结构体中嵌入类型
// 因为这会暴露太多细节

type Widget struct {
	X, Y int
}

type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

// 因为这个接口可以通过 Label 的嵌入带到新的结构体，
// 所以，可以在 Button 中可以重载这个接口方法以
func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}
func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}

func NewButton(x int, y int, label string) Button {
	w := Widget{X: x, Y: y}
	b := Button{Label{Widget: w, Text: label}}
	return b
}

func Example() {
	label := Label{Widget{20, 40}, "teslabel"}
	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}
