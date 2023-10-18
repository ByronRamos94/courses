package calculator

type Terms struct {
	FirstNum  int
	SecondNum int
	CallBack  func()
}

func (t Terms) Sum() int {
	t.CallBack()
	return (t.FirstNum + t.SecondNum)
}
