package main
import (
"fmt"
"os"
"bufio"
)
func main () {
	q:=bufio.NewReader(os.Stdin)
	fmt.Println("Введите формулу: ")
	input, _ := q.ReadString('\n')
	fmt.Println("Скобки расставлены верно: ", %d, "открывающиеся", %d "закрывающиеся", )
	fmt.Println("Скобки расставлены неправильно: ", %d, "открывающиеся", %d "закрывающиеся", )
}