package main

import (
	"demo/plugindemo/store"
	"fmt"
)

func GroupAdd(gs *store.GroupStore, gid int) error {
	fmt.Println("group add, gid", gid)
	return nil
}

func GroupDel(gs *store.GroupStore, gid int) error {
	fmt.Println("group del, gid", gid)
	return nil
}
