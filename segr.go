package main

const orari_segr_default = "15:30-17:00"
const giorno_segr_default = 1 // 1 = "lunedì"
const mesi_segr_da = 1
const mesi_segr_a = 6

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
//{ 8, 1, "10:30-12:30"},
}
