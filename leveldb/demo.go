package main
import(
 "fmt"
 "os"
 "code.google.com/p/leveldb-go/leveldb/table"
  //"code.google.com/p/leveldb-go/leveldb/db"
 //"code.google.com/p/leveldb-go/leveldb"
)
func main(){
  
  f,_:=os.Create("a.db")
  t:=table.NewWriter(f,nil)
 t.Set([]byte("google"),[]byte("yes!"),nil)
 t.Close()
 
 f,_ =os.Open("a.db")
 r:=table.NewReader(f,nil)
 d,_:=r.Get([]byte("google"),nil)
 
 //leveldb.Set([]byte("google"),[]byte("yes!"),nil)
 fmt.Printf("%s",d)

}
 