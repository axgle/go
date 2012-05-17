'http://demon.tw/programming/vbs-picture-resolution.html
Set img = CreateObject("wia.imagefile")

img.loadfile "demo.gif"
wscript.echo img.width 
wscript.echo img.height 



'http://demon.tw/my-work/vbs-tcscconverter-class.html

Set TCSC = New TCSCConverter
WScript.Echo TCSC.SC2TC("我下面给你吃")
WScript.Echo TCSC.TC2SC("我下麵給你吃")
WScript.Echo TCSC.SC2TC("我干妹妹的故事")
WScript.Echo TCSC.TC2SC("我乾妹妹的故事")


Class TCSCConverter
    Private Word, Doc

    'Author: Demon
    'Date: 2011/12/13
    'Website: http://demon.tw

    Private Sub Class_Initialize()
        Set Word = CreateObject("wps.Application") ' or word.Application
    End Sub

    Private Sub Class_Terminate()
        Word.Quit
        Set Word = Nothing
    End Sub

    'Traditional Chinese To Simplified Chinese
    Public Function TC2SC(str)
        Set Doc = Word.Documents.Add
        Word.Selection.TypeText str
        Doc.Range.TCSCConverter 1, True
        TC2SC = Replace(Doc.Range.Text, vbCr, vbCrLf)
        TC2SC = Left(TC2SC, Len(TC2SC) - 2)
        Doc.Saved = True
        Doc.Close
        Set Doc = Nothing
    End Function

    'Simplified Chinese To Traditional Chinese
    Public Function SC2TC(str)
        Set Doc = Word.Documents.Add
        Word.Selection.TypeText str
        Doc.Range.TCSCConverter 0, True, True
        SC2TC = Replace(Doc.Range.Text, vbCr, vbCrLf)
        SC2TC = Left(SC2TC, Len(SC2TC) - 2)
        Doc.Saved = True
        Doc.Close
        Set Doc = Nothing
    End Function

End Class