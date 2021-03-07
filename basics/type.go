// Оператор type позволяет определять именованный тип. Именнованный тип основывается на уже существующем типе
type mile int

// Структуры
// Представляют тип данных, определяемый разработчиком и служащий для представления каких-либо объектов
type person struct{
    name string
    age int
}
// Является новым типом данных
var tom = person {name: "Tom", age: 24}

// Обращеине к полям структуры
tom.age

// Структуры могут быть вложенными
type contact struct{
    email string
    phone string
}
 
type person struct{
    name string
    age int
    contactInfo contact
}

// При наличии в структуре ссылки на структуру того же типа нужно представлять указатель на структуру 
type node struct{
    value int
    next *node
}

// Метод
// Это функция, связанная с определенным типом
type library []string
func (l library) print(){
    for _, val := range l{
        fmt.Println(val)
    }
}

// Вызывается как
lib = []string {"1", "2", "3"}
lib.print()

// Также можно указывать методы структур
// Если метод изменяет состояние структуры, то нужно использовать указатели на структуры
type Human struct{
	name string
    age int
} 

func (h *Human) addAge(newAge int) {
	(*h).age = newAge
}

// Интерфейсы
// Интерфейсы представляют абстракцию поведения других типов.
type vehicle interface{
    move()
}

type Car struct{ }
 
func (c Car) move(){
    fmt.Println("Автомобиль едет")
}

// Интерфейсы реализуются неявно
var tesla Vehicle = Car{}
