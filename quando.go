package main

type mytb struct {
ncor int
gio int
mese int
primo int
}

var quando = []mytb {
mytb { 150,7,1,4 },
mytb { 156,8,01,0 },
mytb { 163,13,01,0 },
mytb { 132,14,01,1 },
mytb { 150,14,01,0 },
mytb { 153,14,01,1 },
mytb { 130,15,01,0 },
mytb { 156,15,01,0 },
mytb { 134,16,01,1 },
mytb { 137,20,01,1 },
mytb { 163,20,01,0 },
mytb { 132,21,01,0 },
mytb { 150,21,01,0 },
mytb { 153,21,01,0 },
mytb { 130,22,01,0 },
mytb { 156,22,01,0 },
mytb { 134,23,01,0 },
mytb { 137,27,01,0 },
mytb { 163,27,01,0 },
mytb { 132,28,01,0 },
mytb { 150,28,01,0 },
mytb { 153,28,01,0 },
mytb { 130,29,01,0 },
mytb { 156,29,01,0 },
mytb { 134,30,01,0 },
mytb { 137,3,02,0 },
mytb { 163,3,02,0 },
mytb { 132,4,02,0 },
mytb { 150,4,02,0 },
mytb { 153,4,02,0 },
mytb { 130,5,02,0 },
mytb { 156,5,02,0 },
mytb { 134,6,02,0 },
mytb { 137,10,02,0 },
mytb { 163,10,02,0 },
mytb { 132,11,02,0 },
mytb { 150,11,02,0 },
mytb { 153,11,02,0 },
mytb { 130,12,02,0 },
mytb { 156,12,02,0 },
mytb { 134,13,02,0 },
mytb { 137,17,02,0 },
mytb { 163,17,02,0 },
mytb { 132,18,02,0 },
mytb { 153,18,02,0 },
mytb { 130,19,02,0 },
mytb { 160,19,02,1 },
mytb { 134,20,02,0 },
mytb { 137,24,02,0 },
mytb { 161,24,02,1 },
mytb { 153,25,02,0 },
mytb { 130,26,02,0 },
mytb { 160,26,02,0 },
mytb { 137,2,03,0 },
mytb { 161,2,03,0 },
mytb { 142,3,03,1 },
mytb { 153,3,03,0 },
mytb { 143,4,03,1 },
mytb { 160,4,03,0 },
mytb { 144,5,03,1 },
mytb { 137,9,03,0 },
mytb { 161,9,03,0 },
mytb { 142,10,03,0 },
mytb { 143,11,03,0 },
mytb { 160,11,03,0 },
mytb { 144,12,03,0 },
mytb { 141,16,03,0 },
mytb { 161,16,03,0 },
mytb { 142,17,03,0 },
mytb { 143,18,03,0 },
mytb { 160,18,03,0 },
mytb { 144,19,03,0 },
mytb { 141,23,03,0 },
mytb { 161,23,03,0 },
mytb { 142,24,03,0 },
mytb { 155,24,03,1 },
mytb { 143,25,03,0 },
mytb { 160,25,03,0 },
mytb { 144,26,03,0 },
mytb { 141,30,03,0 },
mytb { 161,30,03,0 },
mytb { 155,31,03,0 },
mytb { 143,1,04,0 },
mytb { 160,1,04,0 },
mytb { 141,6,04,0 },
mytb { 161,6,04,0 },
mytb { 155,7,04,0 },
mytb { 160,8,04,0 },
mytb { 155,14,04,0 },
mytb { 158,15,04,1 },
mytb { 135,16,04,1 },
mytb { 138,20,04,1 },
mytb { 161,20,04,0 },
mytb { 139,21,04,1 },
mytb { 155,21,04,0 },
mytb { 151,22,04,1 },
mytb { 158,22,04,0 },
mytb { 135,23,04,0 },
mytb { 138,27,04,0 },
mytb { 162,27,04,1 },
mytb { 139,28,04,0 },
mytb { 155,28,04,0 },
mytb { 151,29,04,0 },
mytb { 158,29,04,0 },
mytb { 135,30,04,0 },
mytb { 152,30,04,1 },
mytb { 138,4,05,0 },
mytb { 162,4,05,0 },
mytb { 139,5,05,0 },
mytb { 155,5,05,0 },
mytb { 151,6,05,0 },
mytb { 158,6,05,0 },
mytb { 135,7,05,0 },
mytb { 152,7,05,0 },
mytb { 146,8,05,1 },
mytb { 138,11,05,0 },
mytb { 162,11,05,0 },
mytb { 139,12,05,0 },
mytb { 155,12,05,0 },
mytb { 151,13,05,0 },
mytb { 158,13,05,0 },
mytb { 135,14,05,0 },
mytb { 152,14,05,0 },
mytb { 146,15,05,0 },
mytb { 138,18,05,0 },
mytb { 162,18,05,0 },
mytb { 158,20,05,0 },
mytb { 152,21,05,0 },
mytb { 146,22,05,0 },
mytb { 138,25,05,0 },
mytb { 162,25,05,0 },
mytb { 158,27,05,0 },
mytb { 152,28,05,0 },
mytb { 162,1,06,0 },
mytb { 158,3,06,0 },
mytb { 152,4,06,0 },
mytb { 162,8,06,0 },
mytb { 162,15,06,0 },
mytb { 131,14,10,1 },
mytb { 157,14,10,1 },
mytb { 154,15,10,1 },
mytb { 133,16,10,1 },
mytb { 159,16,10,1 },
mytb { 131,21,10,0 },
mytb { 157,21,10,0 },
mytb { 154,22,10,0 },
mytb { 133,23,10,0 },
mytb { 159,23,10,0 },
mytb { 140,24,10,1 },
mytb { 131,28,10,0 },
mytb { 157,28,10,0 },
mytb { 154,29,10,0 },
mytb { 133,30,10,0 },
mytb { 159,30,10,0 },
mytb { 140,31,10,0 },
mytb { 131,4,11,0 },
mytb { 157,4,11,0 },
mytb { 145,5,11,1 },
mytb { 154,5,11,0 },
mytb { 133,6,11,0 },
mytb { 159,6,11,0 },
mytb { 140,7,11,0 },
mytb { 131,11,11,0 },
mytb { 157,11,11,0 },
mytb { 145,12,11,0 },
mytb { 154,12,11,0 },
mytb { 130,13,11,1 },
mytb { 148,13,11,1 },
mytb { 159,13,11,0 },
mytb { 140,14,11,0 },
mytb { 131,18,11,0 },
mytb { 157,18,11,0 },
mytb { 145,19,11,0 },
mytb { 154,19,11,0 },
mytb { 130,20,11,0 },
mytb { 148,20,11,0 },
mytb { 159,20,11,0 },
mytb { 140,21,11,0 },
mytb { 141,25,11,1 },
mytb { 157,25,11,0 },
mytb { 145,26,11,0 },
mytb { 154,26,11,0 },
mytb { 130,27,11,0 },
mytb { 148,27,11,0 },
mytb { 159,27,11,0 },
mytb { 140,28,11,0 },
mytb { 141,2,12,0 },
mytb { 157,2,12,0 },
mytb { 145,3,12,0 },
mytb { 150,3,12,1 },
mytb { 154,3,12,0 },
mytb { 130,4,12,6 },
mytb { 148,4,12,0 },
mytb { 159,4,12,0 },
mytb { 140,5,12,0 },
mytb { 141,9,12,0 },
mytb { 163,9,12,1 },
mytb { 150,10,12,0 },
mytb { 130,11,12,0 },
mytb { 148,11,12,0 },
mytb { 156,11,12,1 },
mytb { 140,12,12,0 },
mytb { 141,16,12,0 },
mytb { 163,16,12,0 },
mytb { 150,17,12,0 },
mytb { 130,18,12,0 },
mytb { 148,18,12,0 },
mytb { 156,18,12,0 },
mytb { 904,19,12,5 },
}
const verqu="200109.1432"
