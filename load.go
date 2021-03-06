package tcpserverweb

// Load - возвращает значение и статус по id
func (cls *clients) Load(id string) (string, bool) {
	cls.RLock()
	defer cls.RUnlock()
	tcpId, ok := cls.m[id]
	return tcpId, ok
}
