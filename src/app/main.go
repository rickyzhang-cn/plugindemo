package main

import (
	"fmt"
	"net/http"
	"plugin"
	"store"
	"strconv"
)

var (
	gs   *store.GroupStore
	fMap map[string]func(*store.GroupStore, int) error
)

func addHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gidStr := req.Form["gid"][0]
	gid, _ := strconv.Atoi(gidStr)
	fMap["GroupAdd"](gs, gid)
}

func delHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gidStr := req.Form["gid"][0]
	gid, _ := strconv.Atoi(gidStr)
	fmt.Println(gid)
	fMap["GroupDel"](gs, gid)
}

func showHandler(w http.ResponseWriter, req *http.Request) {
	for _, g := range gs.Groups {
		fmt.Fprintln(w, g.Id, g.Players, g.KV)
	}
}

func loadHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fname := req.Form["filename"][0]
	InitLogicFunc(fname)

}

func main() {
	gs = store.InitGroupStore()
	InitLogicFunc("logic_plugin1.so")
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/del", delHandler)
	http.HandleFunc("/load", loadHandler)
	http.HandleFunc("/show", showHandler)
	http.ListenAndServe(":54321", nil)
}

func InitLogicFunc(filename string) {
	p, err := plugin.Open(filename)
	if err != nil {
		panic(fmt.Sprintln("open plugin err:", err, filename))
	}

	fnames := []string{"GroupAdd", "GroupDel"}

	fMap = make(map[string]func(*store.GroupStore, int) error)
	for _, fname := range fnames {
		fn, err := p.Lookup(fname)
		if err != nil {
			panic(fmt.Sprintln("not found symbol", fname, err))
		}
		fMap[fname] = fn.(func(*store.GroupStore, int) error)
	}
	fmt.Println("func map", fMap)
	fmt.Println("loaded plugin successed! file=", filename)
}
