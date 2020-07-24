package iteration

func Repeat(caracter string, length int) (repeat string) {
	for i := 0; i < length; i++ {
		repeat += caracter
	}
	return
}
