package html
import "strconv"
func NumericEntities(s string) string{
 html:=""
 for _,r:=range []rune(s){
	html+="&#"+strconv.Itoa(int(r))+";"	 
 }
 return html 
}