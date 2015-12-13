package main
import (
	"io"
	"strings"
	"fmt"
)
//output n bytes of input stream
func ReadFrom(reader io.Reader, num int)([]byte,error){
	p:=make([]byte, num)
	n,err:=reader.Read(p)
	if n>0{
		return p[:n],nil
	}
	return  p,err
}

func sampleReadFromString(){
	data,_ := ReadFrom(strings.NewReader("fromstring"),5)
	fmt.Print(data)
}

func main(){
	sampleReadFromString()
}