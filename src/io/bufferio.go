package main
import (
	"strings"
	"bufio"
	"fmt"
	"os"
)



func main(){
	strReader :=strings.NewReader("string2 3")
	bufReader:=bufio.NewReader(strReader)
	data,_:=bufReader.Peek(5)
	fmt.Println(data, bufReader.Buffered())
	str, _:=bufReader.ReadString(' ')
	fmt.Println(str,bufReader.Buffered())

	w:=bufio.NewWriter(os.Stdout)
	fmt.Fprint(w,"Firstline")
	w.Flush()
}