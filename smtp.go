//based on http://code.google.com/p/go-wiki/wiki/SendingMail
//run: "godoc net/smtp" see more detail
package main

import (
        "log"
        "net/smtp"
)

func main() {
        addr:="smtp.gmail.com"
		sender:= "yourname@gmail.com"
		password:="yourpassword"
		
		to:="who@domain.com"
		// \n\n for split subject and body,the"subject:" is required because mail protocol
		msg:="subject:this is title\n\nThis is the email body.sent by Go(net/smtp)"
		
        // Set up authentication information.
        auth := smtp.PlainAuth(
                "go agent",//can be empty
               sender,
			   password ,
                addr,
        )
        // Connect to the server, authenticate, set the sender and recipient,
        // and send the email all in one step.
        err := smtp.SendMail(
                addr+":25",
                auth,
                sender,
                []string{to},
                []byte(msg),
        )
        if err != nil {
                log.Fatal(err)
        }
}