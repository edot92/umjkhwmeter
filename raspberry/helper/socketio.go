package helper

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/googollee/go-socket.io"
)

var Serversocketio *socketio.Server

// var clients = make(map[*websocket.Conn]bool) // connected clients
// var broadcast = make(chan Message)           // broadcast channel
// // Define our message object
// type Message struct {
// 	Email    string `json:"email"`
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// }

// // Configure the upgrader
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func handleConnections(ctx *context.Context) {
// 	w := ctx.ResponseWriter
// 	r := ctx.Request
// 	// Upgrade initial GET request to a websocket
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Make sure we close the connection when the function returns
// 	defer ws.Close()

// 	// Register our new client
// 	clients[ws] = true

// 	for {
// 		var msg Message
// 		// Read in a new message as JSON and map it to a Message object
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			log.Printf("error: %v", err)
// 			delete(clients, ws)
// 			break
// 		}
// 		// Send the newly received message to the broadcast channel
// 		broadcast <- msg
// 	}
// }

// func handleMessages() {
// 	for {
// 		// Grab the next message from the broadcast channel
// 		msg := <-broadcast
// 		// Send it out to every client that is currently connected
// 		for client := range clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }
func InitSocketio() {
	// Configure websocket route
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		// helper.SocketIO = so
		so.BroadcastTo("chat", "datapm", "tes")
		// so.On("chat message", func(msg string) {
		// 	log.Println("emit:", so.Emit("chat message", msg))
		// 	so.BroadcastTo("chat", "chat message", msg)
		// })
		so.On("disconnection", func() {
			// log.Println("on disconnect")
		})

	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})
	Serversocketio = server

	beego.Handler("/socket.io/", server)
	// stop := interval.Start(func() {
	// fmt.Println("Hello, World !")
	// a := server.GetMaxConnection()
	// server.BroadcastTo("chat", "chat message", "kikuk")
	// fmt.Println(a)
	// }, 1*time.Second)
	// _ = stop
	// Start listening for incoming chat messages
	// go handleMessages()

	//setup http server
	// serveMux := http.NewServeMux()
	// serveMux.Handle("/socket.io/", server)
}
