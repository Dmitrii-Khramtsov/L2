// Что выведет программа?
// Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
package main

import (
  "fmt"
  "os"
)

func Foo() error {
  var err *os.PathError = nil // в данном примере мы присваеваем nil полю data, 
	// но поле tab - сохраняет тип
  return err
}

// var err *os.PathError = nil

// iface{
// 	tab  = itab{_type=*os.PathError},
// 	data = nil
// }

func main() {
  err := Foo()
  fmt.Println(err) // fmt.Println смотрит на data (содержимое), а оно равно nil - напечатает <nil>
  fmt.Println(err == nil) // интерфейс не nil, потому что tab != nil - напечатает false
}
