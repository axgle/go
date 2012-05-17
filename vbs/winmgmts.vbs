'c:\>wbemtest.exe

strComputer = "."
Set wbemServices = Getobject("winmgmts:\\" & strComputer)
Set wbemObjectSet = wbemServices.InstancesOf("Win32_Process")

 s=""
 On Error Resume next
For Each wbemObject In wbemObjectSet
    If  InStr(LCase(wbemObject.Name),"chrome") then 
	s =s & "Name: " & wbemObject.Name & vbCrLf & _
	" Handle: " & wbemObject.Handle &  _
	" Process ID: " & wbemObject.ProcessID & vbCrLf
     'wbemObject.Terminate()
	End if
Next

'WScript.Echo s


'http://zhidao.baidu.com/question/134633552.html
Set ser = GetObject("winmgmts:")
d=""
Set rows = ser.execquery("select * from win32_process where commandline like '%tencent%' ")

If rows.count=0 then
  MsgBox "count = 0"
End If

 For each k In rows
   d= d & k.commandline & vbcrlf
 next

WScript.Echo d