Set oShell = WScript.CreateObject("WScript.Shell")

oShell.Run "regsvr32  /s AutoItX3_x64.dll"
WScript.Sleep 100

Set oAutoIt = WScript.CreateObject("AutoItX3.Control")
oAutoIt.Run("notepad.exe")
oAutoIt.Run("cmd /c echo 我是中文看看效果如何这样也是可以的|clip")
oAutoIt.WinWaitActive "无标题", ""

oAutoIt.Send("^v")