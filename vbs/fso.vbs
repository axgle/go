Set fso = CreateObject("scripting.filesystemobject") 

Set f = fso.opentextfile(wscript.scriptfullname,1)

MsgBox f.readall
