
 Set sh = CreateObject("WScript.Shell")
' http://www.hackbase.com/tech/2008-12-04/42556.html
 ' sh.popup "auto close after 3 second",30
 
 'http://hi.baidu.com/yiivon/blog/item/2b9ba598c727960e6e068c8a.html
 sh.sendKeys "^+{esc}"
' sh.appactivate "windows"
 wscript.sleep 500
  sh.sendKeys "%fn"
 wscript.sleep 500
  
 sh.sendKeys "chrome http://dashanxue.com"

  sh.sendKeys "{enter}"

 