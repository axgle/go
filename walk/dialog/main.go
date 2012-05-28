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


/*
// work with ui
package main
 
import (
	"github.com/lxn/walk"
)
type Dialog struct{
  *walk.Dialog
  ui *dialogUI
}
func main(){
 
  ui:=&dialogUI{}
  
  d:=&Dialog{ui:ui}
 
  d.init(nil)
  d.Show()
  d.Run()
}

*/