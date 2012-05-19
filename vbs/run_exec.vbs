'http://demon.tw/programming/vbs-run-and-exec.html
Set ws = CreateObject("WScript.Shell")
host = WScript.FullName
 
If LCase(Right(host, 11)) = "wscript.exe" Then
    ws.run "cscript """ & WScript.ScriptFullName & chr(34), 0
    WScript.Quit
End If

Set oexec = ws.Exec( "ipconfig")
Msgbox oExec.StdOut.ReadAll, , "ipconfig"
'此时不要用WScript.Echo，因为当前是在控制台运行
'WScript.Echo的结果会在控制台输出，不会弹出对话框。