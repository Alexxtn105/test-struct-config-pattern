package main

import (
	"fmt"
)

//region Необходимые структуры

// OptFunc функция-опция. Такого типа аргументы необходимо будет передавать в конструктор в качестве параметров
type OptFunc func(*Opts)

// defaultOpts Функция возвращающая значения по умолчанию
func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

// Opts Структура с конфигурацией
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

// withTLS Функция-опция. Устанавливает tls в true
func withTLS(opts *Opts) {
	opts.tls = true
}

// withMaxConn Функция-опция с параметром максимального количества соединений.
// В этом случае возвращаемый результат - это OptFunc. Внутри используется замыкание.
func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

// withID Функция-опция с параметром идентификатора
func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

// Server структура сервера
type Server struct {
	Opts
}

//endregion

// newServer Конструктор. В качестве аргументов принимается слайс функций OptFunc
// Использование такого механизма позволит нам не указывать опции в конструкторе.
// В этом случае будут взяты опции по умолчанию.
// Возможно указать только необходимые опции.
func newServer(opts ...OptFunc) *Server {

	// Сперва берем значения по умолчанию
	o := defaultOpts()

	// бежим по указанным в аргументе опциям
	for _, fn := range opts {
		// вызываем функцию для опции по умолчанию, модифицируя ее
		fn(&o)
	}

	// возвращаем объект
	return &Server{
		Opts: o,
	}
}

func main() {
	// вариант запуска без опций
	//s := newServer()

	// вариант запуска с опциями (можно комбинировать, не использовать вовсе)
	s := newServer(
		withTLS,
		withMaxConn(99),
		withID("identity"),
	)

	fmt.Printf("%+v\n", s)
}
