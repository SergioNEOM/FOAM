package common

import (
	"github.com/SergioNEOM/FOAM/templates"

	"github.com/gin-gonic/gin"
)

const MesKeyName = "FOAMmessage"
const MessageQuestion = 1
const MessageWarning = 2
const MessageError = 4

//SetMessage views error/warning message
func SetMessage(c *gin.Context, mesType int, mesText, mesRef string) {
	m := &templates.MessageBox{
		MesQuestion: (mesType == MessageQuestion),
		MesWarning:  (mesType == MessageWarning),
		MesError:    (mesType == MessageError),
		MessageText: mesText,
		MessageHref: mesRef,
	}
	//c.Set(MesKeyName, m)
	c.HTML(200, "messagebox.tmpl", m)
	c.Abort()
	//fmt.Println("[SetMessage] - redirect")
	//c.Redirect(301, "/message")
}
