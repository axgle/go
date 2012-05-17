'http://demon.tw/programming/generate-guid-in-vbs.html
Set TypeLib = CreateObject("Scriptlet.TypeLib")
strGUID = Left(TypeLib.Guid,38)
WScript.Echo strGUID