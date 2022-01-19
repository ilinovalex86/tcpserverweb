package tcpserverweb

import "net/http"

// Valid - Проверяет cookie web клиента + возвращает индекс tcp клиента
func Valid(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("SessionId")
	if err == nil {
		webId := cookie.Value
		if tcpId, ok := Clients.Load(webId); ok {
			return tcpId, true
		}
	}
	return "", false
}
