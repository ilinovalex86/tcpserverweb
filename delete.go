package tcpserverweb

// Удаляет web клиента, сохраняет json
func (cls *clients) delete(id string) {
	cls.Lock()
	defer cls.Unlock()
	delete(cls.m, id)
	go cls.dump()
}
