# (11) Join - Union - Agregasi - Subquery - Function

3 Point Penting :
1. Join dan Union
- Join, terdapat 3 jenis join yaitu :
	- Inner Join, inner join akan mengembalikan baris baris dari dua table atau lebih yang memenuhi syarat
	- Left Join, Left join akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join.
	- Right Join, Right join akan mengembalikan semua baris dari tabel sebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yang memenuhi kondisi join. Teknik ini merupakan kebalikan dari left join.
- Union, Ada hal yang perlu diperhatikan dari union adalah jumlah field yang dikeluarkan/dipanggil harus sama.
2. Fungsi Agregasi
- MIN, fungsi di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal
- MAX, digunakan untuk mendapatkan nilai maximum atau nilai terbesar dari sebuah data record di tabel
- SUM, Digunakan untuk mendapatkan jumlah total nilai dari sebuah data atau record di tabel
- AVG, Digunakan untuk mencari nilai rata-rata (average) dari sebuah data atau record di tabel
- COUNT, Digunakan untuk mencari jumlah dari sebuah data atau record di tabel
- HAVING, Digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi aggregat
3. Subquery dan Function
- Subquery, subquery digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil, subquery dapat digunakan dengan SELECT, INSERT, UPDATE, dan DELETE statements bersama dengan operator seperti
- Function, sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya