// ui2walk && go build && dialog.exe
package main
 
import (
	"github.com/lxn/walk"
)
type Dialog struct{
  *walk.Dialog
}
func main(){
 // w,_:=walk.NewMainWindow()
  d:=&Dialog{}
 
  d.init(nil)
  d.Show()
  d.Run()
}