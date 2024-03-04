# (06) String - Advance Function - Pointer - Method - Struct and Interface

3 Point Penting :
1. Pada string terdapat berbagai macam fungsi bawaannya, seperti len, compare, contains, substring, replace, dan insert. Selanjutnya Pointer, Pointer adalah sebuah variable yang menyimpan alamat memory dari variable lainnya, pada pointer terdapat 2 operator penting yaitu:
- Deferencing (*) = Berfungsi untuk mendeklarasikan variable pointer dan mengakses value yang disimpan pada suatu alamat
- Referencing (&) = Berfungsi untuk melakukan return alamat dari suatu memori dan mengakses alamat dari suatu memori ke pointer
2. Struct adalah suatu tipe yang didefinisikan oleh pengguna yang berisi kumpulan bidang/properti bernama atau fungsi (metode). kemudian, Interface adalah kumpulan tanda tangan metode yang dapat diimplementasikan oleh sebuah objek. Oleh karena itu, antarmuka menentukan perilaku objek. Sedangkan Method adalah fungsi yang terhubung dengan suatu tipe. Beberapa keunggulan method daripada function adalah sebagai berikut:
- Membantu Anda menulis kode gaya berorientasi objek di Go
- Metode membantu Anda menghindari konflik penamaan
- Panggilan metode jauh lebih mudah dibaca dan dipahami daripada panggilan fungsi
3. Pada bahasa pemrograman golang tidak terdapat Try-Catch, namun sebagai gantinya terdapat Panic-Recover
- Panic = Ketika runtime Go mendeteksi kesalahan-kesalahan ini, itu menyebabkan panic.
- Recover = Untuk menambahkan kemampuan untuk pulih dari kesalahan panic, entah tambahkan sebuah fungsi anonim atau tentukan sebuah fungsi kustom dan panggil dengan kata kunci 'defer' dari dalam metode, di mana panic mungkin terjadi dari panggilan internal lainnya.