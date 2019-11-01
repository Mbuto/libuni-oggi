package main

import (
    "net/http"
    "os"
    "os/exec"
    "fmt"
    "strings"
    "bytes"
)

func info(w http.ResponseWriter, r *http.Request) {

const stage = "stage"
const cmd_linux = "./aggiorlu"
const cmd_dos = "aggiorlu.bat"

s, e := os.Executable()
if e != nil {
	fmt.Fprintf(w, "error: %v\n", e)
	return
}

k := strings.LastIndex(s, "\\")
slash := "\\"
cmd := cmd_dos

if k < 0 {
	k = strings.LastIndex(s, "/")
	slash = "/"
	cmd = cmd_linux
}
k++

x := s[k:]
p := s[:k]
nw := p + stage + slash + x
//fmt.Fprintf(w, "exec: %s %s %s %s\n", s, x, p, nw)
fo, _ := os.Stat(s)
fn, en := os.Stat(nw)
if en != nil {
//	fmt.Fprintf(w, "%s non esiste\n", nw)
} else {
	mto := fo.ModTime()
	mtn := fn.ModTime()
//	fmt.Fprintf(w, "old: %v new: %v\n", mto, mtn)
	if mtn.After(mto) {
		mylog("info: agg. sw " + cmd + "\n")
		fmt.Fprintf(w, "<html><body><b>Ci sono aggiornamenti del software.</b>&nbsp;<button onclick=window.location.assign('/sys?c=%s')>Aggiorna</button></body></html>", cmd)
	} else {
		fmt.Fprintf(w, "")
	}
}
}

func sys(w http.ResponseWriter, r *http.Request) {
com := r.FormValue("c")
if com == "" {
	return
}
mylog("sys: " + com + "\n")
cmd := exec.Command(com)
	var out bytes.Buffer
	var serr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &serr
	err := cmd.Run()
	if err != nil {
		mylog("sys: err\n")
		fmt.Fprintf(w, "sys: err %v\n", err)
		return
	}
	mylog("sys: ok\n")
	fmt.Fprintf(w, "sys: OK \n")
//	fmt.Fprintf(w, "stdout: .%s. \n", out.String())
//	fmt.Fprintf(w, "stderr: .%s. \n", serr.String())
}
