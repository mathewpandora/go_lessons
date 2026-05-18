package main

//в интерфейсах никогда не должно быть лишнего иначе дробим
import "fmt"

type Writer interface {
	Write(string)
}
type DB interface { //
	Read() string
	Writer
}

type MyDB struct {
}

func (d MyDB) Write(string) {
	fmt.Println("Записали")
}

func (d MyDB) Write1(string) {
	fmt.Println("Записали")
}

func (d MyDB) Read() string {
	return "my string"
}

type Logger struct {
}

func (d Logger) Write(string) {
	fmt.Println("Записали")
}

// solid - I - Interface agregastion plincaple (говорит о том что интерфейсы должны содержать только необходимые методы )
func main() {

	db := MyDB{}
	logger := Logger{}
	run(db, logger)
}

func run(db DB, logger Writer) { //вот так грамотно!
	db.Write("1")
	//db.Write1("1") -не может использовать методы которых нет в интерфейсе
	fmt.Println(db.Read())
	logger.Write("FVfdv")
}

//Интерфейсы не могут иметь методы с одинаковыми именами, если они имеют разные возвращаемые значения. Это приведет к конфликту, и компилятор выдаст ошибку.
