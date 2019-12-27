package main

import (
    "fmt"
    "time"
    "strconv"
    "strings"
    "net/http"
    "os"
    "log"
)


const version = "8.18:heroku"

const miniscr = 10
const aula_def = "Auletta Libuni"

var anno1_n int
var anno1_s string
var anno2_n int
var anno2_s string

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/altradata", altradata)
    http.HandleFunc("/stat", stat)
    http.HandleFunc("/cal", cal)
    http.HandleFunc("/calnext", calnext)
    http.HandleFunc("/orari", orari)
    http.HandleFunc("/info", info)
    http.HandleFunc("/sys", sys)
// MAGIC!
http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))
    http.HandleFunc("/robots.txt", robots)
}

func mylog(s string) {
f, err := os.OpenFile("lu.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}

logger := log.New(f, "LU: ", log.LstdFlags)

logger.Printf(s)
f.Close()
}

func root(w http.ResponseWriter, r *http.Request) {

if r.URL.Path != "/" {
                        fmt.Fprintf(w, "<html><head><title>Non trovata</title><body><h1>Pagina %s non trovata.</h1><a href=/>Home</a></body></html>", r.URL.Path)
                        return
                }

//FORSEMENTE...
t := time.Now().In(time.FixedZone("UTC+1", 0))

//FAKE ********************************************
// t = time.Date(2018, time.Month(10), 1, 0, 0, 0, 0, time.FixedZone("UTC+1", 0))

y0, m0, d0 := t.Date()
if int(m0) > 5 {
	anno1_n = y0
} else {
	anno1_n = y0 - 1
}

anno2_n = anno1_n + 1
anno1_s = fmt.Sprintf("%d", anno1_n)
anno2_s = fmt.Sprintf("%d", anno2_n)

yoko := "ono oggi"
if r.Method == "POST" {
	anno := r.FormValue("anno")
	if anno < anno1_s || anno > anno2_s {
		errore(w,"Errore Anno non previsto", anno)
		return
	}
	mese := r.FormValue("mese")
	giorno := r.FormValue("giorno")
	aa, err1 := strconv.Atoi(anno)
	if err1 != nil {
		errore(w,"Errore Anno non numerico", anno)
		return
	}
	mm, err2 := strconv.Atoi(mese)
	if err2 != nil {
		errore(w,"Errore Mese non numerico", mese)
		return
	}
	if mm < 1 || mm > 12 {
		errore(w,"Errore Mese", mese)
		return
	}
	gg, err3 := strconv.Atoi(giorno)
	if gg < 1 || gg > 31 {
		errore(w,"Errore Giorno", giorno)
		return
	}
	if err3 != nil {
		errore(w,"Errore Giorno non numerico", giorno)
		return
	}
	if (aa == anno2_n && mm > 6) || (aa == anno1_n && mm < 10) {
		errore(w,"Errore: questa data non esiste nel Calendario Corsi: ", anno + "/" + mese + "/" + giorno)
		return
	}
	t = time.Date(aa, time.Month(mm), gg, 0, 0, 0, 0, time.FixedZone("UTC+1", 0))
	yoko = "ono"
}

y, m, d := t.Date()
x := 10000 * y + 100 * int(m) + d
x0 := 10000 * y0 + 100 * int(m0) + d0
if (y < y0) || (x < x0) {
	errore(w,"Errore: data riferita al passato: ", fmt.Sprintf("%d",y) + "/" + fmt.Sprintf("%d",int(m)) + "/" + fmt.Sprintf("%d",d))
	return
}

fmt.Fprintf(w,mioForm0)
if (m > time.June && y == anno2_n ) || (m < time.October && y == anno1_n ) {
	disp_buone_vacanze(w)
	fmt.Fprintf(w,mioForm3,y,m,d)
	fmt.Fprintf(w,mioForm2,anno1_n, anno2_n, version, verac, verqu)
	return
}
if (m == time.December && d >= 20) || (m == time.January && d <= 6) {
	disp_buone_feste(w)
	t = time.Date(anno2_n, time.January, 6, 0, 0, 0, 0, time.FixedZone("UTC+1", 0))
} else {
	cercacorsi(w, d,m,y, yoko)
}
if yoko != "ono" {
	yoko = "eranno"
}

for qt := t.Add(time.Hour *24); qt != t.Add(time.Hour * 24 * 7); qt = qt.Add(time.Hour * 24) {
y, m, d = qt.Date()
cercacorsi(w, d,m,y, yoko)
}

//xt := t.Add(time.Hour * 24 * 7);
//xy, xm, xd := xt.Date();

zy, zm, zd := t.Date()
zt := t
for {
zt = zt.Add(time.Hour * 24);
zy, zm, zd = zt.Date()
if zt.Weekday().String() == "Monday" {
break
}
}

//fmt.Fprintf(w,mioForm3,xy,xm,xd,zy,zm,zd)
fmt.Fprintf(w,mioForm3,zy,zm,zd)
fmt.Fprintf(w,mioForm2,anno1_n, anno2_n, version, verac, verqu)
}

func disp_buone_vacanze(w http.ResponseWriter) {
fmt.Fprintf(w, "<h2>Libuni augura Buone Vacanze</h2><img src=/images/Buone-Vacanze.png width='30%%'><br clear=all><h2>Le attivit&agrave; riprenderanno ad Ottobre!</h2>")
}

func disp_buone_feste(w http.ResponseWriter) {
fmt.Fprintf(w, "<h2>Libuni augura Buone Feste</h2><img src=/images/Buone-Feste.jpg width='40%%'><br clear=all><h2>Le attivit&agrave; riprenderanno a Gennaio!</h2>")
}

func errore(w http.ResponseWriter, msg string, ddt string) {
fmt.Fprintf(w, errForm, msg, ddt)
fmt.Fprintf(w, mioForm2, anno1_n, anno2_n, version, verac, verqu)
}

func stat(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, statForm)
fmt.Fprintf(w, "<h3>Corsi/Eventi: %d<br>Date: %d<br>Link: %d<br>Doce: %d</h3>", len(corsi), len(quando), len(linktab), len(doctab))
fmt.Fprintf(w, mioForm2, anno1_n, anno2_n, version, verac, verqu)
}

func cercacorsi(w http.ResponseWriter, gg int, mm time.Month, yy int, og string) {
numcorsi := 0
sett := []string {"lunedì","martedì","mercoledì","giovedì","venerdì","sabato","domenica"}
week := []string {"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}
tt := time.Date(yy, mm, gg, 0, 0, 0, 0, time.FixedZone("GMT+1", 0))
settday := "boh"
for j := 0; j < len(week); j++ {
	if week[j] == tt.Weekday().String() {
		settday = sett[j]
		break
	}
}
fmt.Fprintf(w,mioForm1, og, settday, gg, int(mm), yy)
yess := yesegr(int(tt.Weekday()), int(mm), gg)
ycor := yescor(int(mm), gg)
if ycor || (yess != "") {
fmt.Fprintf(w, "<table><tr><th>Corso / Attivit&agrave;</th><th>Docente</th><th>Orario</th><th>Aula / Luogo</th><th>Tutor</th></tr>")
}

if yess != "" {
	fmt.Fprintf(w, "<tr><td class='colsegre'>%s</td><td>%s</td><td>%s</td><td>%s</td><td>&nbsp;---&nbsp;</td></tr>", "Segreteria Aperta", "Segretari", yess, "Stanza Segreteria")
}
for k := 0; k < len(quando); k++ {
	if (quando[k].gio == gg) && (quando[k].mese == int(mm)) {
		ncor := quando[k].ncor
		for j := 0; j < len(corsi); j++ {
			if corsi[j].numco == ncor {
				cls := ""
				n := quando[k].primo
				m := corsi[j].nomeco
				o := corsi[j].orari
				d := corsi[j].dove
				z := corsi[j].doce
				if n > 0 && n < len(cosa) {
					cls = " class='yel' "
					m = m + "<br><span class='it'>" + cosa[n] + "</span>"
					if n == 6 {
						cls = " class='xred'"
						o = "<center><img src=images/attenzione.gif width=25px></center>"
						d = o
					}
				}
				if d == "" {
					d = aula_def
				}
				if n != 3 {
					m = trovalink(ncor, m)
					z = trovadoce(z)
				}
				fmt.Fprintf(w, "<tr><td%s>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", cls, m, z, o, d, corsi[j].tutor)
			numcorsi++
			break
			}
		}
	}
}
if (numcorsi == 0) && (yess == "") {
	fmt.Fprintf(w, "<span class=red>%s non ci sono corsi n&eacute; attivit&agrave;</span><p>", settday)
} else {
	fmt.Fprintf(w, "</table><p>")
}

}

func trovalink(c int, s string) string {
for k := 0; k < len(linktab); k++ {
if linktab[k].numco == c {
	return("<a title='Click per descrizione' target=_blank href=https://libuni.blogspot.com/" + linkpref + linktab[k].link + ">" + s + "</a>")
}
}
return("<span title='Descrizione non disponibile'>" + s + "</span>")
}

func trovadoce(s string) string {
z := strings.Replace(strings.Replace(strings.Replace(strings.ToLower(s), " ", "-", -1), "à", "a", -1), "'", "", -1)
for k:=0; k < len(doctab); k++ {
	if z == doctab[k].link {
		return("<a target=_blank href=https://libuni.blogspot.com/" + docpref + z + ".html>" + s + "</a>")
	}
}
return("<span title='curriculum non trovato'>" + s + "</span")
}

func altradata(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w,mioForm0)
t := time.Now().In(time.FixedZone("UTC+1", 0))
y0, m0, d0 := t.Date()
if int(m0) > 7 {
	anno1_n = y0
} else {
	anno1_n = y0 - 1
}
anno2_n = anno1_n + 1
fmt.Fprintf(w, altra, d0, int(m0), y0, anno2_n, y0)
}

func robots(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "User-agent: *\nDisallow: /\n")
}

func cal(w http.ResponseWriter, r *http.Request) {
var mesi = []string { "", "Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno", "Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre" }
sett := []string {"dom","lun","mar","mer","gio","ven","sab"}

t := time.Now().In(time.FixedZone("UTC+1", 0))
y0, m0, d0 := t.Date()
mes := int(m0)
savedy0 := y0

//FAKE ***********************************************
// mes = 10
// y0 = 2018

mese := r.FormValue("mese")
if mese != "" {
mes, _ = strconv.Atoi(mese)
}

if mes < 1 || mes > 12 {
return
}

//BUGFIX
if mes < int(m0) {
	y0++
}

maxmes := 31
switch mes {
case 4, 6, 9, 11:
	maxmes = 30
	break
case 2:
	maxmes = 28
// gestire max-mes anni bisestili
	if (y0 % 100) > 0 && (y0 % 4) == 0 {
		maxmes = 29
	}
}

//init_ipc()

totcor := 0
extrasegclass := "hid"

fmt.Fprintf(w, calForm0, mesi[mes], mesi[mes], y0)
for d:=1; d <= maxmes; d++ {
tt := time.Date(y0, time.Month(mes), d, 0, 0, 0, 0, time.FixedZone("GMT+1", 0))
wd := tt.Weekday()
if wd == 6 || wd == 0 {
fmt.Fprintf(w, "<th class=red>%d<br>%v</th>", d, sett[wd])
} else {
fmt.Fprintf(w, "<th>%d<br>%v</th>", d, sett[wd])
}
}
fmt.Fprintf(w, "</tr><tr><th>Corso / Docente</th><th>Quando / Dove</th>")
for d:=1; d <= maxmes; d++ {
	tt := time.Date(y0, time.Month(mes), d, 0, 0, 0, 0, time.FixedZone("GMT+1", 0))
	wd := tt.Weekday()
	qs := qsegr(d,mes,int(wd))
	if strings.Contains(qs,"MS") || strings.Contains(qs,"PS") {
		extrasegclass = "sil"
	}
	fmt.Fprintf(w, "<th class='sil'>%s</th>", qs)
}
fmt.Fprintf(w, "</tr>")
s := ""
for j:=0; j < len(corsi); j++ {
u := 0
tit := ""
for k:=0; k < len(quando); k++ {
rarr := "&nbsp;"
cls := ""
cls2 := "blu"
if quando[k].primo > 1 {
cls = "evi"
cls2 = "evi"
}
if quando[k].primo == 6 {
cls = ""
cls2 = "xred"
rarr = "&nbsp;"
}
if quando[k].primo == 4 {
cls = "vi"
cls2 = "vi"
rarr = "<span class=rarr>&rarr;</span>"
}
if corsi[j].numco == quando[k].ncor {
ss := fmt.Sprintf(";%d",corsi[j].numco)
if !strings.Contains(s,ss) && (quando[k].mese == mes) {
totcor++
fmt.Fprintf(w, "<tr><td class=%s>", cls)
if quando[k].primo == 1 {
fmt.Fprintf(w, "<img class='myimg' src='images/new.png'>&nbsp;")
/**************************** FIX per heroku - no DB
ipcfnd := false
for x:=0; x < len(ipctab); x++ {
  if quando[k].ncor == ipctab[x].ncor {
	if ipctab[x].iscr < miniscr {
		fmt.Fprintf(w, "<span class='blu noprint' title='Pochi iscritti'>%02d</span> ", ipctab[x].iscr)
	}
	ipcfnd = true
	break
  }
}
if ipcfnd == false {
	fmt.Fprintf(w, "<span class='xred noprint' title='Nessun iscritto'>00</span> ")
}
*****************************************************************/
}
m := trovalink(corsi[j].numco, corsi[j].nomeco)
fmt.Fprintf(w, "%s", m)
fmt.Fprintf(w, "<br><small>%s</small>", corsi[j].doce)
d := corsi[j].dove
if d == "" {
	d = aula_def
}
fmt.Fprintf(w, "</td><td><small>%s<br>%s</small></td>", corsi[j].orari, d)
s = fmt.Sprintf("%s;%s",s,ss)
tit = strings.Replace(corsi[j].nomeco, "<br>", " ", -1) + " " + corsi[j].orari
} // if already
if quando[k].mese == mes {
for n:=u+1; n < quando[k].gio; n++ {
if mes == int(m0) && n < d0 {
fmt.Fprintf(w, "<td class='gray'>%s</td>", rarr)
} else {
tt := time.Date(y0, time.Month(mes), n, 0, 0, 0, 0, time.FixedZone("GMT+1", 0))
wd := tt.Weekday()
if wd == 6 || wd == 0 {
fmt.Fprintf(w, "<td class=sd>%s</td>", rarr)
} else {
fmt.Fprintf(w, "<td>%s</td>", rarr)
}
}
} // for
fmt.Fprintf(w, "<td class=%s title=\"%s\">%d/%d</td>", cls2, tit, quando[k].gio, quando[k].mese)
u = quando[k].gio
} // mese == 3
} // if found
} //quando
if u > 0 {
for uu := u+1; uu <= maxmes; uu++ {
if mes == int(m0) && uu < d0 {
fmt.Fprintf(w, "<td class='gray'>&nbsp;</td>")
} else {
tt := time.Date(y0, time.Month(mes), uu, 0, 0, 0, 0, time.FixedZone("GMT+1", 0))
wd := tt.Weekday()
if wd == 6 || wd == 0 {
fmt.Fprintf(w, "<td class=sd>&nbsp;</td>")
} else {
fmt.Fprintf(w, "<td>&nbsp;</td>")
}
}
}
fmt.Fprintf(w, "</tr>")
}
} //corsi
if totcor == 0 {
	fmt.Fprintf(w, "<tr><td colspan=2 class=red>Nel mese di %s %d</td><td colspan=10 class=red>Non ci sono corsi n&eacute; attivit&agrave;</td></tr>", mesi[mes], y0)
}

m1 := ((mes) % 12) + 1
m2 := ((mes+1) % 12) + 1
m3 := ((mes+2) % 12) + 1
gpass := "gray"
if mes != int(m0) {
	gpass = "hid"
}
fmt.Fprintf(w, calForm1, gpass, extrasegclass, extrasegclass, d0, mesi[int(m0)], savedy0, m1, mesi[m1], m2, mesi[m2], m3, mesi[m3]) 
fmt.Fprintf(w, mioForm2, anno1_n, anno2_n, version, verac, verqu)
}


const nxt = `
<html>
<head>
   <!-- HTML meta refresh URL redirection -->
   <meta http-equiv="refresh"
   content="0; url=/cal?mese=%d">
</head>
<body>
<h2>Attendere qualche istante...</h2>
</body>
</html>`

func qsegr(g int, m int, wd int) string {
if m < 10 && m > 6 {
	return("&nbsp;")
}
if (m == 12 && g > 20) || (m == 1 && g < 7) {
	return("&nbsp;")
}
for k := 0; k < len(sisegr); k++ {
	if sisegr[k].gio == g && sisegr[k].mese == m {
	if sisegr[k].orario == "" {
		return "&nbsp;"
	}
	if sisegr[k].orario < "15" {
		s := fmt.Sprintf("<span title='%s'>MS</span>", sisegr[k].orario)
		return s
	} else {
		s := fmt.Sprintf("<span title='%s'>PS</span>", sisegr[k].orario)
		return s
	}
	}
}
if wd == int(giorno_segr_default) {
	s := fmt.Sprintf("<span title='%s'>Po</span>", orari_segr_default)
	return s
}
return("&nbsp;")
}

func calnext(w http.ResponseWriter, r *http.Request) {
t := time.Now().In(time.FixedZone("UTC+1", 0))
_, m0, _ := t.Date()
mes := int(m0)
m1 := ((mes) % 12) + 1

fmt.Fprintf(w, nxt, m1)
}

const ipchdr = `
<html>
<head>
<style>
.yel {background-color: yellow; color: red; }
.red {background-color: red; color: white; }
</style>
</head>
<body>
<h1>Numero di Iscritti ai Corsi</h1>
`

const ipcbot = `
<p><i>Dati del: %s</i></p>
<p><a href="/">Home</a>
<p>Libera Università di Citt&agrave; della Pieve APS - Cod.Fisc.: 94056590543
<p>v.%s&nbsp;-&nbsp;<a href="https://libuni.blogspot.com/p/docente-carlo-zappala.html"><span class=cp>&copy; 2018-2019 C. Zappal&agrave;</span></a>
</body>
</html>`


/****************************************
func ipcmese(w http.ResponseWriter, mm int) {
var mesi = []string { "", "Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno", "Luglio", "Agosto", "Settembre", "Ottobre", "Novembre", "Dicembre" }

// ad Aprile, ad Ottobre
d := ""
if mm == 4 || mm == 8 || mm == 10 {
	d = "d"
}
fmt.Fprintf(w, "<h3>Corsi che iniziano a%s %s</h3>", d, mesi[mm])
fmt.Fprintf(w, "<table><tr><th>Corso</th><th>Iscritti</th><th>Giorno</th><th>Mese</th></tr>")
tot := 0
cri := 0
for m :=0; m < len(quando); m++ {
found := 0
if quando[m].mese == mm && quando[m].primo == 1 {
for k:=0; k < len(ipctab); k++ {
 if quando[m].ncor == ipctab[k].ncor {
  for j := 0; j < len(corsi); j++ {
	if corsi[j].numco == ipctab[k].ncor {
		found = ipctab[k].iscr
		cls := ""
		if ipctab[k].iscr < miniscr {
			cls = "yel"
			cri++
		}
  		fmt.Fprintf(w, "<tr><td>%s</td><td class='%s'>%d</td><td>%d</td><td>%d</td></tr>", 
			corsi[j].nomeco, cls, found, quando[m].gio, quando[m].mese)
		tot++
		break
	}
  }
 }
}
if found == 0 {
  for j := 0; j < len(corsi); j++ {
	if corsi[j].numco == quando[m].ncor {
  		fmt.Fprintf(w, "<tr><td>%s</td><td class='red'>%d</td><td>%d</td><td>%d</td></tr>", 
			corsi[j].nomeco, found, quando[m].gio, quando[m].mese)
	}
  }
		tot++
		cri++
}
}
}
fmt.Fprintf(w, "</table><p><b>%s:</b> Num. Corsi: %d - Critici: %d", mesi[mm], tot, cri)
}
**************************/

/************************
func ipc(w http.ResponseWriter, r *http.Request) {

init_ipc()

fmt.Fprintf(w, ipchdr)
mese := r.FormValue("mese")
if mese != "" {
mm, err2 := strconv.Atoi(mese)
if err2 != nil {
	errore(w,"Errore Mese non numerico", mese)
	return
}
if mm < 1 || mm > 12 {
	errore(w,"Errore Mese", mese)
	return
	}
ipcmese(w, mm)
} else {
// ott-dic
for k:=10; k < 13; k++ {
	ipcmese(w, k)
}
// gen-mag
for k:=1; k < 6; k++ {
	ipcmese(w, k)
}
}
fmt.Fprintf(w, ipcbot, veripc, version)
}
********************************************/

func orari(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "<html><body><h2>Controllo orari sovrapposti</h2><table><tr><th>Num.</th><th>Corso</th><th>gg</th><th>mm</th><th>ora</th></tr>")
n := 0
saveor := ""
savenome := ""
savecor := 0
saveday := 0
saveme := 0
for k:=0; k < len(quando); k++ {
for j:=0; j < len(corsi); j++ {
if corsi[j].numco == quando[k].ncor {
if quando[k].gio == saveday && quando[k].mese == saveme {
o1 := strings.Split(saveor, "-")
o2 := strings.Split(corsi[j].orari, "-")
//fmt.Fprintf(w, "<tr><td colspan=7>DBG: %d o2[0]=%s &le; o1[1]=%s and o2[1]=%s &ge; o1[0]=%s</td></tr>", quando[k].ncor, o2[0], o1[1], o2[1], o1[0])
if (o2[0] <= o1[1]) && (o2[1] >= o1[0]) {
fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td>%d</td><td>%d</td><td>%s</td></tr>", savecor, savenome, quando[k].gio, quando[k].mese, saveor)
fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td>%d</td><td>%d</td><td>%s</td></tr><tr><td colspan=5>&nbsp;</td></tr>", quando[k].ncor, corsi[j].nomeco, quando[k].gio, quando[k].mese, corsi[j].orari)
n++
}
}
saveor = corsi[j].orari
savecor = corsi[j].numco
saveday = quando[k].gio
saveme = quando[k].mese
savenome = corsi[j].nomeco
}
}
}
fmt.Fprintf(w, "</table><p>Fine controllo: %d corsi, %d conflitti.<br>", len(corsi), n)
if n == 0 {
	fmt.Fprintf(w, "<b>Non</b> ci sono sovrapposizioni di orari.")
}
fmt.Fprintf(w, "<p><a href='/'>Home</a><p>Libera Università di Citt&agrave; della Pieve APS - Cod.Fisc.: 94056590543<p>v.%s&nbsp;-&nbsp;<a href='https://libuni.blogspot.com/p/docente-carlo-zappala.html'><span class=cp>&copy; 2018-2019 C. Zappal&agrave;</span></a></body></html>", version)
}

func yescor(m int, d int) bool {
for k := 0; k < len(quando); k++ {
	if (quando[k].gio == d) && (quando[k].mese == m) {
		return true
	}
}
return false
}

func yesegr(sd int, m int, d int) string {
// segreteria chiusa fra giugno e settembre

if m > mesi_segr_a && m < mesi_segr_da {
	return ""
}

// controlla tabella extra
for j:=0; j < len(sisegr); j++ {
	if m == sisegr[j].mese && d == sisegr[j].gio {
		if sisegr[j].orario != "" {
			return "<span class='yel'>" + sisegr[j].orario + "</span>"
		} else {
			return ""
		}
	}
}

// se non trovato, apre solo il giorno prefissato
if sd != giorno_segr_default {
	return ""
}
return orari_segr_default
}

const calForm0 = `
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Calendario Corsi Libuni - %s</title>
<link rel="icon" href="/images/favicon.ico" />
<style>
@media screen {
body { background-color:#607d8b; color: white; }
h1 { color: gold; }
h3 { color: #f4ff81; }
}
.red { color: red; font-weight: bold; background-color: #f4ff81; }
.blu { background-color: navy; color: gold; text-align: center; }
.evi { color: navy; background-color: yellow; }
.gray { background-color: gray; }
.sil { background-color: silver; }
.xred { color: white; background-color: red; }
.vi { color: black; background-color: #BACD66; }
.right { text-align: right; }
.rgt { align: right; }
.sd { background-color:	rgb(255, 140, 102); }
.hid {display: none; }
.rarr { text-align: center; font-size: 30px; }
table { border-collapse: collapse; }
table, th, td { border: 1px solid navy; padding: 2px; color: navy; background-color: #f4ff81; }
table.none, td.none, th.none { border: none; !important }
hr { align: left; }
.cp { color: navy; }
a { color: black; text-decoration: none; }
a:hover { background-color: orange; color: black; }
@media print {
h3 {display: none; }
button {visibility: hidden; }
.noprint {display: none; }
}
.myimg {
  border: 2px solid red;
  border-radius: 10px;
  width: 20px;
}
img {
  vertical-align: middle;
}
#rad {
  background-color: #f4ff81;
  color: navy;
  width: 130px;
  border: 2px solid red;
  padding: 10px;
  border-radius: 25px;
}
</style>
</head>
<body>
<h1><a href="https://libuni.blogspot.com/"><img src=/images/logo-300.jpg></a>
 Calendario Corsi del mese di %s %d</h1>
<table><tr><th colspan=2>Giorno</th>
`

const calForm1 = `
</table><p></p>
<table>
<tr><td colspan=10><i><b>Legenda:</b></i></td></tr>
<tr><td><img class='myimg' src='images/new.png'>&nbsp;Nuovi Corsi<br><small>Iniziano questo mese</small></td>
<td class="evi">&nbsp;Conferenze&nbsp;<br>&nbsp;Eventi&nbsp;</td>
<td class="%s">Giorni<br>passati</td>
<td class="blu">Lezioni</td>
<td class="sd">Sabato<br>Domenica</td>
<td class="xred">Lezione<br>Spostata</td>
<td class="vi">Lezione<br>Recuperata</td>
<td class="%s"><b>MS</b>=Segr. Ap. Straord.<br>Mattina</td>
<td class="%s"><b>PS</b>=Segr. Ap. Straord.<br>Pomeriggio</td>
<td class="sil"><b>Po</b>=Segr. Ap. ord.<br>Pomeriggio</td>
</tr>
<tr class=noprint><td colspan=10><b>Cliccando sul nome corso si accede alla sua descrizione se disponibile.</b></td></tr></table>
<h3>Oggi: %d-%s-%d</h3>
<span class="noprint">
<p></p><form action="/cal">Cambia mese: <select name="mese">
  <option value="%d" selected>%s</option>
  <option value="%d">%s</option>
  <option value="%d">%s</option>
</select>
<button type="submit">Vai</button>
</form>
<button onclick="window.print()">Stampa</button>
<p>
</span>
`

const errForm = `
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Oggi in Libuni</title>
<link rel="icon" href="/images/favicon.ico" />
<style>
body { background-color:#607d8b; color: white; }
h1 { color: gold; }
h3 { color: #f4ff81; }
.red { color: red; font-weight: bold; background-color: #f4ff81; }
table { border-collapse: collapse; }
table, th, td { border-bottom: 1px solid navy; padding: 5px; color: navy; background-color: #f4ff81; }
hr { align: left; }
.cp { color: navy; }
a { color: black; text-decoration: none; }
a:hover { background-color: orange; color: black; }
#rad {
  background-color: #f4ff81;
  color: navy;
  width: 130px;
  border: 2px solid red;
  padding: 10px;
  border-radius: 25px;
}
</style>
</head>
<body>
<a href="https://libuni.blogspot.com/"><img src=/images/logo-300.jpg></a>
<h1>Errore</h1>
<span class=red>%s %s</span>
<p>
`

const statForm = `
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Oggi in Libuni</title>
<link rel="icon" href="/images/favicon.ico" />
<style>
body { background-color:#607d8b; color: white; }
h1 { color: gold; }
h3 { color: #f4ff81; }
.red { color: red; font-weight: bold; background-color: #f4ff81; }
table { border-collapse: collapse; }
table, th, td { border-bottom: 1px solid navy; padding: 5px; color: navy; background-color: #f4ff81; }
hr { align: left; }
.cp { color: navy; }
a { color: black; text-decoration: none; }
a:hover { background-color: orange; color: black; }
#rad {
  background-color: #f4ff81;
  color: navy;
  width: 130px;
  border: 2px solid red;
  padding: 10px;
  border-radius: 25px;
}
</style>
</head>
<body>
<a href="https://libuni.blogspot.com/"><img src=/images/logo-300.jpg></a>
<h1>Statistiche</h1>
`

const mioForm0 = `
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Oggi in Libuni</title>
<link rel="icon" href="/images/favicon.ico" />
<style>
@media screen {
body { background-color:#607d8b; color: white; }
h1 { color: gold; }
h3 { color: #f4ff81; }
}
.colsegre { font-weight: bold; color: navy; background-color: #F58223; }
.red { color: red; font-weight: bold; background-color: #f4ff81; }
.xred { color: white; background-color: red; }
.yel { background-color: gold; color: navy; font-weight: bold; }
.it { background-color: gold; color: navy; font-weight: normal; font-style: italic; }
.null { border: none; padding: 5px; color: navy; background-color: #607d8b; }
.rgt { align: right; }
table { border-collapse: collapse; }
table, th, td { border-bottom: 1px solid navy; padding: 5px; color: navy; background-color: #f4ff81; }
hr { align: left; }
table.none, td.none, th.none { border: none; !important }
.cp { color: navy; }
a { color: black; text-decoration: none; }
a:hover { background-color: orange; color: black; }
img {
  vertical-align: middle;
}
#rad {
  background-color: #f4ff81;
  color: navy;
  width: 130px;
  border: 2px solid red;
  padding: 10px;
  border-radius: 25px;
}
@media print {
button {visibility: hidden; }
.noprint {display: none; }
}
</style>
</head>
<body>
<h1><a href="https://libuni.blogspot.com/"><img src=/images/logo-300.jpg></a>
 Calendario 7 giorni</h1>
<span class=noprint><b><i>Cliccando sul nome corso si accede alla sua descrizione se disponibile.
<br>Cliccando sul nome docente si accede al suo curriculum se disponibile.</i></b></span>
`

const mioForm1 = `
<h3>Corsi e Attivit&agrave; che si svolg%s %s %d-%d-%d.</h3>
`

const mioForm2 = `
</tr></table>
<div class="noprint">
<blockquote>
<div id="rad" class="noprint">
<center>
<b>Menu</b></center>
<a href="/"><img class=myicon src="images/home.png" width=20> Libuni Oggi <img class=myicon src="images/home.png" width=20></a><br>
<a target='_blank' href="https://libuni.blogspot.com/p/corsi-aa-%d-%d.html">&nbsp;&nbsp;&nbsp;&nbsp;Schede Corsi
<img class=myicon src='images/right.png' width=20></a><br>
<a href="https://libuni.blogspot.com/"><img class=myicon src='images/left.png' width=20>
Sito Web Libuni</a>
</div>
</blockquote>
</div>
<span class="noprint">
<iframe style="border:none;" src="/info" height=50 width="90%%"></iframe>
<br>___<br><br>
Questo sito non utilizza nessun cookie<br>
e non conserva nessun dato dei visitatori.
</span>
<p>Libera Università di Citt&agrave; della Pieve APS - Cod.Fisc.: 94056590543
<p>v.%s&nbsp;-&nbsp;<span class=cp>&copy; 2018-2019</span>
<br>v.ac=%s v.qu=%s
</body></html>
`

const mioForm3 = `
<table class="null noprint"><tr class="null noprint"><td class="null noprint">
<form action="/" method="post">
<input type="hidden" name="anno" value="%d">
<input type="hidden" name="mese" value="%d">
<input type="hidden" name="giorno" value="%d">
<button type="submit">Prossimo Luned&igrave;</button></form>
</td><td class="null noprint">
<form action="/altradata"><button type="submit">Altra data</button></form>
</td>
</td><td class="null noprint">
<form action="/cal"><button type="submit">Calendario Mensile</button></form>
</td>
`

// ATTENZIONE: CAMBIARE ANNO
const altra = `
<form action="/" method="POST">
<table>
<tr><td>GG</td><td><input name="giorno" type=number min=1 max=31 value=%d autofocus></td></tr>
<tr><td>MM</td><td><input name="mese" type=number min=1 max=12 value=%d></td></tr>
<tr><td>AA</td><td><input name="anno" type=number min=%d max=%d value=%d></td></tr>
<tr><td>&nbsp;</td><td><button type="submit">Ok</button></td></tr>
</table>
</form>
<div id="rad" class="noprint">
<a href="/"><img class=myicon src="images/home.png" width=20> Libuni Oggi <img class=myicon src="images/home.png" width=20></a><br>
</div>
</body>
</html>
`

var cosa = []string { "", "Prima Lezione", "Uscita", "Evento", "Lezione Recuperata", "Conferenza", "Lezione Spostata" }
