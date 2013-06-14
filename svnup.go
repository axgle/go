//svn up with golang online
package main

import (
	"flag"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	p := flag.String("p", "5200", "port")
	flag.Parse()
	http.HandleFunc("/svnup", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("svn", "up")

		d, e := cmd.Output()
		if e != nil {
			println(e)
		}

		//println(string(d))
		io.WriteString(w, string(d))

	})
	http.ListenAndServe("0.0.0.0:"+*p, nil)

}
