'http://demon.tw/programming/vbs-uac-elevation.html
For Each objOS in GetObject("winmgmts:").InstancesOf("Win32_OperatingSystem")
	If InStr(objOS.Caption,"XP") = 0 Then
		If WScript.Arguments.length = 0 Then
			Dim objShell
			Set objShell = CreateObject("Shell.Application")
			objShell.ShellExecute "wscript.exe", Chr(34) & _
			WScript.ScriptFullName & Chr(34) & " uac", "", "runas", 1
		Else
			Call Main()
		End If
	Else
		Call Main()
	End If
Next

Sub main()

Set sh  = CreateObject("wscript.shell")
MsgBox sh.exec("cmd /c reg query HKCU\environment").stdout.readall

End sub

 