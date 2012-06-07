package html
import "testing"
func TestEntities(t *testing.T){
	testStrings:=map[string]string{
		" ":"&#32;",//space
		"<":"&#60;",
		"=":"&#61;",
		">": "&#62;",	
		"÷":"&#247;",
		"©":"&#169;",//&copy; is Name format like php's mb_convert_encoding($string, "HTML-ENTITIES", "UTF-8"); 
		"中文":"&#20013;&#25991;",
		"编码": "&#32534;&#30721;",
	}
	
	for k,v:=range testStrings{
		result:=NumericEntities(k)
		if result!=v{
			t.Errorf("%s != %s",result,v)
		}
	}
}