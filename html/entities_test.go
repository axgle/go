package html
import "testing"
func TestEntities(t *testing.T){
	testStrings:=map[string]string{
		"©":"&#169;",//&copy; //todo: complete like php's mb_convert_encoding($string, "HTML-ENTITIES", "UTF-8"); 
		"中文":"&#20013;&#25991;",
		"golang unicode编码": "golang unicode&#32534;&#30721;",
	}
	
	for k,v:=range testStrings{
		result:=Entities(k)
		if result!=v{
			t.Errorf("%s != %s",v,result)
		}
	}
}