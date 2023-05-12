package msggateway

import "sync"

type UserMap struct {
	m sync.Map
}

func newUserMap() *UserMap {
	return &UserMap{}
}
func (u *UserMap) GetAll(key string) ([]*Client, bool) {
	allClients, ok := u.m.Load(key)
	if ok {
		return allClients.([]*Client), ok
	}
	return nil, ok
}
func (u *UserMap) Get(key string, platformID int) ([]*Client, bool, bool) {
	allClients, userExisted := u.m.Load(key)
	if userExisted {
		var clients []*Client
		for _, client := range allClients.([]*Client) {
			if client.platformID == platformID {
				clients = append(clients, client)
			}
		}
		if len(clients) > 0 {
			return clients, userExisted, true

		}
		return clients, userExisted, false
	}
	return nil, userExisted, false
}
func (u *UserMap) Set(key string, v *Client) {
	allClients, existed := u.m.Load(key)
	if existed {
		oldClients := allClients.([]*Client)
		oldClients = append(oldClients, v)
		u.m.Store(key, oldClients)
	} else {
		var clients []*Client
		clients = append(clients, v)
		u.m.Store(key, clients)
	}
}
func (u *UserMap) delete(key string, platformID int) (isDeleteUser bool) {
	allClients, existed := u.m.Load(key)
	if existed {
		oldClients := allClients.([]*Client)
		var a []*Client
		for _, client := range oldClients {
			if client.platformID != platformID {
				a = append(a, client)
			}
		}
		if len(a) == 0 {
			u.m.Delete(key)
			return true
		} else {
			u.m.Store(key, a)
			return false
		}
	}
	return existed
}
func (u *UserMap) DeleteAll(key string) {
	u.m.Delete(key)
}