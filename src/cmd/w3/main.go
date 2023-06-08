package main

import (
	"log"
	"github.com/surdeus/gosrv/src/httpx"
	"github.com/surdeus/gosrv/src/tmplx"
)

type TableEntry struct {
	Time string
	Problems string
	Nodes []int
}

var (
	AddrStr = ":8080"
	templates tmplx.Templates
	CableProblems = "Повреждение кабеля или проблема с затуханием его оптического кабеля"
)


func handler(c *httpx.Context) {
	tables := []TableEntry{
		TableEntry{
			"2023-05-28 15:42:41",
			CableProblems,
			[]int{1901, 3041},
		},
		TableEntry{
			"2023-04-26 7:42:41",
			CableProblems,
			[]int{1508, 2706},
		},
		TableEntry{
			"2023-05-07 16:33:51",
			CableProblems,
			[]int{1476, 2804},
		},
	}
	templates.Exec(c.W, "default", "table", tables)
}

func main() {
	var err error
	funcMap := tmplx.StdFuncMap()
	funcMap["sum"] = tmplx.Sum
	funcMap["neg"] = tmplx.Neg
	tcfg := tmplx.ParsingConfig{
		Component: "tmpl/c/",
		View: "tmpl/v/",
		Template: "tmpl/t/",
		FuncMap: funcMap,
	}
	templates, err = tmplx.Parse(tcfg)
	if err != nil {
		panic(err)
	}

	mux := httpx.NewMux()
	mux.Define(
		httpx.Def("/").Re("").
			Method("get", handler),
	)

	srv := httpx.Server {
		Addr: AddrStr,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

