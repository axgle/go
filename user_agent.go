package main
import "net/http" 

func main(){
 
client:=&http.Client{}
req,_:=http.NewRequest("GET","http://dhv.cc",nil)
req.Header.Add("user-agent","go~~")
client.Do(req)
 
}