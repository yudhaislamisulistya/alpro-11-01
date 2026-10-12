<!-- i = 0
<perulangan>
output(i)
i = i * 2
i++
output("rekayasa perangkat lunak")
<akhir perulangan>
output(i)
 -->

i = 2
<perulangan dengan batas 3>
output(i * 2)
i = i - 1
i++
output(i)
<akhir perulangan jika mencapai 3 perulangan>


Perulangan Pertama
Nilai i awal = 2
Menampilkan data dari i = 2 * 2 = 4
i = i - 1 => (2 - 1 = 1)
i++ (increment) = 1 + 1 = 2
Menampilkan data dari i adalah 2

Perylangan kedua
nilai i = 2
Menampilkan data dari i = 2 * 2 = 4

a = 10
b = 20
for iteration <- a to b do​
    <action>​
endfor​

for (int a = 2; a < 10; a+=2){
    <action>
}

a = 2 -> 2 < 10 -> TRUE -> Do a+=2 -> 2 + 2 = 4
a = 4 -> 4 < 10 -> TRUE -> Do a+=2 -> 4 + 2 = 6
a = 6 -> 6 < 10 -> TRUE -> Do a+=2 -> 6 + 2 = 8
a = 8 -> 8 < 10 -> TRUE -> Do a+=2 -> 8 + 2 = 10

a = 10 -> 10 < 10 -> FALSE


a = 0 -> 0 < 10 -> TRUE -> Do
a = 1 -> 1 < 10 -> TRUE -> Do
...
a = 9 -> 9 < 10 -> TRUE -> Do

a = 10 -> 10 < 10 -> FALSE -> End


for(int i = 11; i >= 0; i-=3){
    <action>
}


i = 11 -> 11 >= 0 -> TRUE -> Do i-=3 -> 11 - 3 = 8
i = 8 -> 8 >= 0 -> TRUE -> Do i-=3 -> 8 - 3 = 5
i = 5 -> 5 >= 0 -> TRUE -> Do i-=3 -> 5 - 3 = 2
i = 2 -> 2 >= 0 -> TRUE -> Do i-=3 -> 2 - 3 = -1
