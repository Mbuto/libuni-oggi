package main

import (
    "log"
    "time"
    "fmt"
//    "database/sql"
//    _ "github.com/mattn/go-sqlite3"
)

type myipc struct {
ncor int
iscr int
}

const mincorso = "130"

var ipctab = []myipc {
}
var veripc = ""

func init_ipc() {

ipctab = nil
t := time.Now().In(time.FixedZone("UTC+1", 0))
c := fmt.Sprintf("%v", t)
veripc = c[:19]

/**************************************************************
database, err1 := sql.Open("sqlite3", "../libuni.db")
if err1 != nil {
log.Fatalf("ERR DB Open %v\n", err1)
return
}

qry := "select id_corso, count(id_corso) 'iscritti' from iscritti_corsi where id_corso >= " + mincorso + " group by id_corso"
rows, err := database.Query(qry)

if err != nil {
log.Fatalf("ERR DB Query %v\n", err)
return
}

for rows.Next() {
	corso := 0
	iscritti := 0
	rows.Scan(&corso, &iscritti)
	new := myipc{corso, iscritti}
	ipctab = append(ipctab, new)
}
rows.Close()
************************************************************/
}
