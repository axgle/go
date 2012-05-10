package main

import (
     "net/http"
	 "net/url"
	 "fmt"
	 "io/ioutil"
	
)

 

func main() {

 
 //?mod=forumdisplay&fid=42
params:=url.Values{  "mod":{"forumdisplay"},  "fid":{"42"}}

//r,_:=http.Get("http://dashanxue.com/forum.php?mod=forumdisplay&fid=42")
//Get for get,PostForm for post with params
r,_:=http.PostForm("http://dashanxue.com/forum.php",params)

//defer r.Body.Close()

b,_:=ioutil.ReadAll(r.Body)
 r.Body.Close()
fmt.Printf(string(b))
 

}
