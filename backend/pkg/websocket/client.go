

import{
	"github.com/gorilla/websocket"
}

type Client struct{
	ID string
	Conn *websocket.Conn
	Pool *pool
	mu sync.Mutex
}

type Message struct{
	Type int `json:"type"`
	Body string `json:"body"`

}

func (c *Client) Read(){
	defer func(){
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
	messageType,p,err := c.Conn.ReadMessage()
	if err != nil 
	{
		log.Println(err)
		return 
	}
	message := Message{Type:messageType, Body:string}
	c.Pool.Broadcast <- message
	fmt.Println("Message received:%+v\n",message)
}
}