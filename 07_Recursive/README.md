# (07) Recursive - Number Theory - Sorting - Searching

3 Point Penting :
1. Recursif
Rekursif adalah suatu keadaan di mana sebuah fungsi memecahkan masalah dengan memanggil dirinya sendiri secara berulang. Jika masalahnya cukup kecil, fungsi rekursif dapat menghasilkan jawaban secara langsung. Jika masalahnya terlalu besar, fungsi akan memanggil dirinya sendiri dengan cakupan masalah yang lebih kecil. Ada dua hal yang perlu dipertimbangkan saat menggunakan strategi rekursif:
- Base Case = Apa kasus paling sederhana dari masalah ini?
- Recurrence relations = Apa hubungan rekursif dari masalah ini dengan masalah yang lebih kecil yang serupa?
Salah satu contoh penerapan rekursif adalah pada saat membuat fungsi untuk menghitung faktorial.
2. Teori bilangan
Teori bilangan adalah cabang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam bidang teori bilangan, seperti Bilangan Prima, Faktor Persekutuan Terbesar (GCD), Kelipatan Persekutuan Terkecil (LCM), Faktorial, Faktor Prima, dll.
3. Searching dan Sorting
- Searching = proses untuk menemukan posisi nilai yang diberikan dalam daftar nilai. Ada beberapa contoh searching seperti : 
   - Linier Search, dengan kompleksitas waktu O(n)
   - Binary Search, dengan kompleksitas waktu O(log n)
   - Builtins Search, yaitu search yang sudah disediakan oleh bahasa pemrograman Golang, contoh : sort.SearchInts()
- Sorting = proses penyusunan data dalam urutan tertentu. Biasanya, kita mengurutkan berdasarkan nilai elemen-elemennya. Kita dapat mengurutkan angka, kata-kata, pasangan nilai, dll. Sebagai contoh, kita dapat mengurutkan siswa berdasarkan tinggi badan mereka, dan kita dapat mengurutkan kota-kota secara alfabetis atau berdasarkan jumlah penduduknya. Urutan yang paling umum digunakan adalah urutan numerik dan urutan alfabetis. Ada beberapa contoh sorting seperti :
   - Selection Sort, dengan kompleksitas waktu O(n^2)
   - Counting Sort, dengan kompleksitas waktu O(n + k)
   - Merge Sort, dengan kompleksitas waktu O(log n)
   - Builtins Sort, yaitu sorting yang sudah disediakan oleh bahasa pemrograman Golang, contoh : sort.Strings()