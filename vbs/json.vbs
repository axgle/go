'http://demon.tw/my-work/vbs-json.html
Function ParseJson(strJson)
    Set html = CreateObject("htmlfile")
    Set window = html.parentWindow
    window.execScript "var json = " & strJson, "JScript"
    Set ParseJson = window.json
End Function

strJson = CreateObject("scripting.filesystemobject").openTextFile("json.txt").readall

Set j = ParseJson(strJson)
wscript.echo j.name & ":" & j.version & j.more.name  