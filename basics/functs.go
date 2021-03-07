// Объявление функции
func main() {

}

// Возвращает данные
func get() string {
	return "qq"
}

// Множественные возвращаемые значения
func get2() (int, int) {
	return 1, 2
}

// Получение множественных значений

var a, b int
a, b = get2()

// Испольщование параметров функции
func set(x, y int16) int16 {
	return x * y
}

// Отсрочка разрешения
// Тело после defer выполняется при завершении выполнения команд в данной области
func testDefer() {
	defer get()	// Ключевое слово defer
	a := 10
	a += 20
}