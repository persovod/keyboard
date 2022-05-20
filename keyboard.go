package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"

)
const Exzam float64=60 // создаем констану 
func GetFloat()(float64,error){   // функция для чтения с клавиатуры
	reader := bufio.NewReader(os.Stdin)
input,err := reader.ReadString('\n')
if err != nil {  // если при чтении ввода будет ошибка, она будет возвращена функцией
	return 0, err
}
input=strings.TrimSpace(input)
number,err := strconv.ParseFloat(input,64)
if err != nil{    // также возвращаются ошибки при преобразованиии строки в float64
	return 0,err
}
return number,nil
}
