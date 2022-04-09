package logic

import "pair/service/chat/api/internal/svc"

type Manager struct {
	chatroom   map[string]*Hub
	register   chan *Hub
	unregister chan *Hub
	svcCtx     *svc.ServiceContext
}

func NewManager(ctx *svc.ServiceContext) *Manager {
	return &Manager{
		chatroom:   make(map[string]*Hub),
		register:   make(chan *Hub),
		unregister: make(chan *Hub),
		svcCtx:     ctx,
	}
}

func (m *Manager) NewHub(name string, uid int64, toUid int64) *Hub {
	var hub *Hub
	if _, ok := m.chatroom[name]; ok == false {
		hub = NewHub(name, uid, toUid, m)
		go hub.Run()

		m.register <- hub

		return hub
	}

	return m.chatroom[name]
}

func (m *Manager) Run() {
	for {
		select {
		case hub := <-m.register:
			m.chatroom[hub.name] = hub
		case hub := <-m.unregister:
			if _, ok := m.chatroom[hub.name]; ok {
				delete(m.chatroom, hub.name)
				close(hub.broadcast)
			}
		}
	}
}
