'http://os.baiup.com/vbs/2335.html
set ws=wscript.createobject("Wscript.shell")  
ws.logevent 0 ,"write log success"
ws.logevent 1 ,"write log error"
ws.logevent 2 ,"write log warning"
ws.logevent 4 ,"write log info"
ws.logevent 8 ,"write log sucess ..."
ws.logevent 16,"write log ... ..."

ws.run "eventvwr"

wscript.Sleep 10000

ws.run "cmd /k wmic nteventlog where filename=""application"" call cleareventlog"
ws.run "cmd /k wmic nteventlog where filename=""system"" call cleareventlog"