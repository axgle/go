Set WshShell=CreateObject("WScript.Shell")
strDesKtop=WshShell.SpecialFolders("DesKtop")

Set oShellLink=WshShell.CreateShortcut(strDesKtop&"\画图.lnk")

oShellLink.TargetPath="mspaint.exe"

oShellLink.WindowStyle=1

oShellLink.Hotkey="CTRL+SHIFT+P"
 
oShellLink.IconLocation="mspaint.exe,0"
oShellLink.Description="用标准VBS建立的画笔快捷方式"

oShellLink.WorkingDirectory=strDesKtop
 
oShellLink.Save
