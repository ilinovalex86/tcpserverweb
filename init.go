package tcpserverweb

import (
	"encoding/json"
	ex "github.com/ilinovalex86/explorer"
	tcp "github.com/ilinovalex86/tcpserver"
	"log"
	"sync"
)

var clientsDb = db{file: "webClients.json"}

type db struct {
	sync.Mutex
	file string
}

// Response - Структура для ответа браузеру
type Response struct {
	Menu           []ex.LinkAndName
	Dirs           []ex.LinkAndName
	Files          []ex.LinkAndName
	WebServer      string
	Sep            string
	User           string
	FreeTcpClients []tcp.Available
	Error          string
}

// Clients - Веб клиенты - cookies
var Clients = &clients{m: make(map[string]string)}

type clients struct {
	sync.RWMutex
	m map[string]string
}

//if err != nil -> log.Fatal(err)
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Инициализация cookies
func init() {
	if ex.ExistFile(clientsDb.file) {
		data := ex.ReadFileFull(clientsDb.file)
		err := json.Unmarshal(data, &Clients.m)
		check(err)
	}
}
