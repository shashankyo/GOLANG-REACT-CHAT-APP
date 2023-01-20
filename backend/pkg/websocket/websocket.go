package websocket

import(
	"log"
	"net/http"
	"github.com/gorilla/websockets"
)

var upgrader = websocket.Upgrade{
	ReadBufferSize : 1024,
	WriteBufferSize : 1024,

}



func Upgrade(w http.ResponseWriter, r * http.Request) (*websocket.Conn, error){
	upgrader.CheckOrigin = func(r *http.Rewquest) bool {return true}
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}