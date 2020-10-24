package main

const orari_segr_default = "15:30-18:00"
const giorno_segr_default = 1 // 1 = "lunedì"
const mesi_segr_da = 9
const mesi_segr_a = 12

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
extraseg { 26, 10, ""},
extraseg { 2, 11, ""},
extraseg { 9, 11, ""},
}
