package main

const orari_segr_default = "16:00-18:00"
const giorno_segr_default = 1 // 1 = "lunedì"
const mesi_segr_da = 10
const mesi_segr_a = 5

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
}
