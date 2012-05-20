Sub download(url,target)
	On Error Resume next
	Const adTypeBinary = 1
	Const adSaveCreateOverWrite = 2
	Dim http,ado
	Set http = CreateObject("Msxml2.XMLHTTP")
	http.open "GET",url,False
	http.send
	Set ado = createobject("Adodb.Stream")
	ado.Type = adTypeBinary
	ado.Open
	ado.Write http.responseBody
	ado.SaveToFile target
	ado.Close
End Sub

download "http://demon.tw/programming/vbs-download-file.html","demon.tw.html"