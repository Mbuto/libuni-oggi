package main

const orari_segr_default = "15:30-18:00"
const giorno_segr_default = 10 // 1 = "lunedì"
const mesi_segr_da = 9
const mesi_segr_a = 12

type extraseg struct {
gio int
mese int
orario string
}

var sisegr = []extraseg {
{ 12, 10, "15:30-17:00" },
{ 13, 10, "15:30-17:00" },
{ 14, 10, "15:30-17:00" },
{ 18, 10, "15:30-17:00" },
{ 19, 10, "15:30-17:00" },
{ 20, 10, "15:30-17:00" },
{ 21, 10, "15:30-17:00" },
}
