// based on https://github.com/nsf/gocode/blob/master/client.go func try_run_server

package main

import "os"

//import "time"

func main() {
	/**
	cwd, _ := os.Getwd()
	procattr := os.ProcAttr{
		Dir:   cwd,
		Env:   os.Environ(),
		Files: []*os.File{nil, nil, nil},
	}	
	*/

	procattr := os.ProcAttr{
		//Dir:"d:/www",
		Files: []*os.File{nil, nil, nil},
	}
	p, _ := os.StartProcess(`d:\phpstudy\php5\php.exe`,
		[]string{"php", "-S", "localhost:9999"}, &procattr,
	)

	p.Release()

}
