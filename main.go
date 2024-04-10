package main

type lamp struct{ status bool }

func (l *lamp) TurnOn() {
	l.status = true
}

func (l *lamp) TurnOff() {
	l.status = false
}

func SumTwoNumber(a int, b int) int {
	return a + b
}

func HelloWorld(name string) string {
	return "hello" + name
}

func main() {

}
