package html
import "strconv"
func Entities(s string) string{
 html:=""
 for _,r:=range []rune(s){
	i:=int(r)
	if i < 128 {
		html+=string(r)
	}else{
		html+="&#"+strconv.Itoa(i)+";"
	}
 }
 return html 
}