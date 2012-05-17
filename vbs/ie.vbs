'http://demon.tw/my-work/vbs-discuz-auto-reply.html

Set ie = CreateObject("internetexplorer.application")

url = "http://dashanxue.com/forum.php?mod=viewthread&tid=338&extra=page%3D1"
ie.navigate2 url

 

ie.visible = True
WScript.Sleep 3000


IE.Document.GetElementById("fastpostmessage").innerHTML = "Hello world"
IE.Document.getElementById("fastpostsubmit").click
 