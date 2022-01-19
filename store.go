package tcpserverweb

// Store - добавляет значение tcpId с ключом id, сохраняет json
func (cls *clients) Store(id string, tcpId string) {
	cls.Lock()
	defer cls.Unlock()
	cls.m[id] = tcpId
	go cls.dump()
}
