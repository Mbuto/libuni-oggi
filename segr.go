package main

const orari_segr_default = "16:00-18:00"
const giorno_segr_default = 9 // 1 = "lunedì"
const mesi_segr_da = 9
const mesi_segr_a = 12

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
extraseg { 1, 9, "10.30-12.30"},
extraseg { 3, 9, "10.30-12.30"},
extraseg { 5, 9, "10.30-12.30"},
extraseg { 7, 9, "16.30-18.30"},
extraseg { 9, 9, "16.30-18.30"},
extraseg { 11, 9, "16.30-18.30"},
}
