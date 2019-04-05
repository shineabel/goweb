package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)



var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}


func Ping( c *gin.Context)  {

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}



	defer ws.Close()


	for {
		mt, _, err := ws.ReadMessage()
		if err != nil {
			break
		}


		t := time.Now().String()

		msg2 := []byte(" hello,pong" +  t)


		err = ws.WriteMessage(mt,msg2)
		if err != nil {
			break
		}
	}
}
