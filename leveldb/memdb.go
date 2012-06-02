package main
import(
 
 
 "code.google.com/p/leveldb-go/leveldb/memdb"
 //"code.google.com/p/leveldb-go/leveldb/table"
  //"code.google.com/p/leveldb-go/leveldb/db"
 //"code.google.com/p/leveldb-go/leveldb"
)

func p(r []byte,e error){
  println(string(r))
}
func main(){
 
  db:=memdb.New(nil)
  db.Set([]byte("name"),[]byte("axgle"),nil)
  
  p( db.Get([]byte("name") ,nil) )
  
  
}