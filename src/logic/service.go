package main

import (
	"fmt"
	"store"
)

func GroupAdd(gs *store.GroupStore, gid int) error {
	fmt.Println("group add, gid", gid)
	gs.Lock()
	defer gs.Unlock()
	g := new(store.Group)
	gs.Groups[gid] = g
	g.Id = gid
	g.Players = []string{"wangzhizhi", "yaoming"}
	g.KV = make(map[int]string)
	g.KV[1] = "yaoming"
	return nil
}

func GroupDel(gs *store.GroupStore, gid int) error {
	fmt.Println("group del, gid", gid)
	gs.Lock()
	delete(gs.Groups, gid)
	return nil
}
