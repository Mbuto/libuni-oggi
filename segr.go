package main

const orari_segr_default = "15:30-17:00"
const giorno_segr_default = 10 // 1 = "lunedì"
const mesi_segr_da = 9
const mesi_segr_a = 12

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
{ 16, 12, ""},
}
