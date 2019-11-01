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
extraseg { 1,10,"" },
extraseg { 7,10,"15.00-17.00" },
extraseg { 8,10,"15.00-17.00" },
extraseg { 10,10,"15.00-17.00" },
extraseg { 9,10,"09.30-12.00" },
extraseg { 11,10,"09.30-12.00" },
extraseg { 12,10,"09.30-12.00" },
extraseg { 14,10,"15.00-17.00" },
extraseg { 15,10,"15.00-17.00" },
extraseg { 16,10,"15.00-17.00" },
extraseg { 17,10,"15.00-17.00" },
extraseg { 18,10,"15.00-17.00" },
}
