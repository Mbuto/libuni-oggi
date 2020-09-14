package main

const orari_segr_default = ""
const giorno_segr_default = 1 // 1 = "lunedì"
const mesi_segr_da = 9
const mesi_segr_a = 12

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
extraseg { 7, 9, ""},
extraseg { 14, 9, ""},
}
