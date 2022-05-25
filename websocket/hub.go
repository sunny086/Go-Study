package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	sendMessage chan *SendMessage
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
	// Register requests from the clients.
	clientMap map[string]*Client
}

type SendMessage struct {
	User string
	Msg  []byte
}

func newHub() *Hub {
	return &Hub{
		broadcast:   make(chan []byte),
		sendMessage: make(chan *SendMessage),
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		clientMap:   make(map[string]*Client),
	}
}

var h = newHub()

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true // client注册，
			h.clientMap[client.Data.User] = client
		case client := <-h.unregister:
			// client注销，将该client从map中删除，并关闭send通道
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				delete(h.clientMap, client.Data.User)
				close(client.send)
			}

		case message := <-h.broadcast:
			// 将每一个user发送的到broadcast通道的信息，通过hub发送到所有client的通道。
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.sendMessage:
			client := h.clientMap[message.User]
			select {
			case client.send <- message.Msg:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}
