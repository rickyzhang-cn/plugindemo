package store

import "sync"

type Group struct {
	Id      int
	Players []string
	KV      map[int]string
}

type GroupStore struct {
	sync.RWMutex
	Groups map[int]*Group
}

func InitGroupStore() *GroupStore {
	gs := new(GroupStore)
	gs.Groups = make(map[int]*Group)
	return gs
}
