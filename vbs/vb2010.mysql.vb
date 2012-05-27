'mysql connector net http://dev.mysql.com/downloads/connector/net/
Imports MySql.Data.MySqlClient
Public Class Form1
    Dim conn As MySqlConnection

    Dim cmd As MySqlCommand
    Dim read As MySqlDataReader

    Private Sub Form1_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
        Me.Text = "google"
        conn = New MySqlConnection()
        conn.ConnectionString = "uid=root;pwd=root;database=wordpress;server=localhost"
        conn.Open()
        cmd = New MySqlCommand("select ID,post_title,post_content from wp_posts limit 100", conn)
        read = cmd.ExecuteReader()
    End Sub

    Private Sub Label1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Label1.Click

        read.Read()
        TextBox1.Text = CStr(read.GetString(0))

        TextBox2.Text = CStr(read.GetString(1) + read.GetString(2))

    End Sub

    Private Sub Button2_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button2.Click
        Close()

    End Sub

    Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click

        Label1_Click(sender, e)

    End Sub

    Private Sub Form1_FormClosed(ByVal sender As System.Object, ByVal e As System.Windows.Forms.FormClosedEventArgs) Handles MyBase.FormClosed
        conn.Close()
        conn.Dispose()

    End Sub
End Class
