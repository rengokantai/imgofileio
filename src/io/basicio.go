package main
import (
	"io"
	"strings"
	"fmt"
	"os"
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

func readFromString(){
	data,_ := ReadFrom(strings.NewReader("fromstring"),5)
	fmt.Print(data)
}

func readFromStdin(){
	fmt.Println("enter a string")
	data,_:=ReadFrom(os.Stdin,7)
	fmt.Print(data)
}

func readFromFile(){
	//pwd,_:=os.Getwd()
	file,_ :=os.Open("src\\io\\basicio.go")
	defer file.Close()
	data,_:=ReadFrom(file,7)
	fmt.Println(data)
}

func main(){
	//readFromString()
	//readFromStdin()
	readFromFile()
}