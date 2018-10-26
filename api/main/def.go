package main

var (
	MAX_UPLOAD_SIZE int64 = 5 * 1024 * 1024 //5MB
	PIC_DIR               = "./upload//pics/"
)

type FilePath struct {
	Url string `json:"url"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
// Configure the upgrader
var upgrader = websocket.Upgrader{}

// Define our message object
type Message struct {
	Message string `json:"message"`
}
