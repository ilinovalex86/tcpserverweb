package tcpserverweb

// Проверяет есть ли клиент с таким id
func (cls *clients) exist(id string) bool {
	cls.RLock()
	defer cls.RUnlock()
	for k, _ := range cls.m {
		if k == id {
			return true
		}
	}
	return false
}
