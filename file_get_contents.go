package main
import "fmt"
import "net/http" 
 import "io/ioutil" 
 
 func file_get_contents(url string,call func(string)){
    
	r,_:=http.Get(url)
	defer r.Body.Close()
	d,_:=ioutil.ReadAll(r.Body)
	if(call!=nil){
		call(string(d))
	}
 }
func main(){
 
	file_get_contents("http://baidu.com",func(data string){	
		fmt.Printf(data)
	} )

}