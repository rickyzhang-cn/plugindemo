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
	fMap map[string]ExecFunc
)

type ExecFunc func(*store.GroupStore, int) error

func addHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gidStr := req.Form["gid"][0]
	gid, _ := strconv.Atoi(gidStr)
	fmt.Println(gid)
	fMap["GroupAdd"](gs, gid)
}

func delHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gidStr := req.Form["gid"][0]
	gid, _ := strconv.Atoi(gidStr)
	fmt.Println(gid)
	fMap["GroupDel"](gs, gid)
}

func loadHandler(w http.ResponseWriter, req *http.Request) {

}

func main() {
	gs = store.InitGroupStore()
	InitLogicFunc("")
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/del", delHandler)
	http.HandleFunc("/load", loadHandler)
	http.ListenAndServe(":54321", nil)
}

func InitLogicFunc(filename string) {
	p, err := plugin.Open(filename)
	if err != nil {
		fmt.Println("open plugin err:", err, filename)
		return
	}

	fnames := []string{"GroupAdd", "GroupDel"}

	fMap := make(map[string]ExecFunc)
	for _, fname := range fnames {
		fn, err := p.Lookup(fname)
		if err != nil {
			fmt.Println("not found symbol", fname, err)
			continue
		}
		fMap[fname] = fn.(ExecFunc)
	}

	fmt.Println("loaded plugin successed! file=", filename)
}
