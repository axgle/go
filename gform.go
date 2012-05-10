package main

import (
     "net/http"
	"net/url"
	 "fmt"
	 "io/ioutil"
	
)

func main() {
 
q:=url.Values{  "mod":{"forumdisplay"},  "fid":{"42"}}

r,_:=http.PostForm("http://dashanxue.com/forum.php",q)

defer r.Body.Close()

b,_:=ioutil.ReadAll(r.Body)
 
fmt.Printf(string(b))


}
