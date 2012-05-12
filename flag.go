package main

import ("fmt"
"flag"
"os"
)


func main() {
var dir = flag.String("dir", ".", "Directory to publish")

   flag.Parse()
   fmt.Printf("program's name:%s  -dir=%s \n",os.Args[0], *dir)
   fmt.Printf("len(os.Args):%d  flag.NArg():%d (this line:before=after+1,because os include program's name,flag package not include)\n",len(os.Args), flag.NArg())
   fmt.Printf("flag.NFlag():%d \n",flag.NFlag())
   
   
}