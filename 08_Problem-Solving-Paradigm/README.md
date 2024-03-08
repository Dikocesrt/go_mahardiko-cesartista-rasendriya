# (08) Problem Solving Paradigm - Brute Force - Greedy - Dynamic Programming

3 Point Penting :
1. Problem Solving Paradigm
Problem Solving Paradigm adalah pendekatan yang umumnya digunakan untuk menyelesaikan masalah, beberapa contohnya yaitu :
- Complete Search/Brute Force
- Algoritma Divide and Conquer
- Algoritma Greedy 
- Dynamic Programming
2. Algoritma Complete Search/Brute Force, Divide and Conquer, dan Greedy
- Complete Search atau Brute Force = suatu metode untuk menyelesaikan masalah dengan menelusuri seluruh ruang pencarian untuk mendapatkan solusi yang diperlukan. Brute Force terjadi ketika tidak ada algoritma lain yang tersedia. Biasanya mudah ditulis karena sifatnya yang langsung. Secara teoritis, semua masalah dapat dipecahkan menggunakan pendekatan Brute Force terutama ketika Anda memiliki waktu yang tidak terbatas.
- Divide and Conquer = paradigma penyelesaian masalah di mana suatu masalah dibuat lebih sederhana dengan 'membagi'nya menjadi bagian-bagian yang lebih kecil dan kemudian menaklukkan setiap bagian. Langkah-langkahnya adalah:
   - Divide  : Memisahkan masalah besar menjadi masalah-masalah yang lebih kecil.
   - Conquer : Ketika masalah sudah cukup kecil untuk diselesaikan, langsung diselesaikan.
   - Combine : Jika diperlukan, solusi dari masalah-masalah yang lebih kecil digabungkan menjadi solusi untuk masalah yang lebih besar.
- Greedy = Sebuah algoritma dikatakan bersifat "greedy" jika pada setiap langkahnya membuat pilihan yang secara lokal optimal dengan harapan akhirnya mencapai solusi yang secara global optimal. Dalam beberapa kasus, pendekatan greedy berhasil - solusinya pendek dan berjalan secara efisien.
3. Dynamic Programming
Dynamic Programming adalah teknik algoritmik untuk menyelesaikan masalah optimisasi dengan memecahnya menjadi submasalah yang lebih sederhana dan memanfaatkan fakta bahwa solusi optimal untuk masalah keseluruhan tergantung pada solusi optimal untuk submasalahnya. 
- Karakteristik Dynamic Programming
   - Overlapping Subproblems = versi yang lebih kecil dari masalah asli. Setiap masalah memiliki submasalah yang tumpang tindih jika menemukan solusinya melibatkan penyelesaian submasalah yang sama berkali-kali.
   - Optimal Substructure Property = Setiap masalah memiliki properti substruktur optimal jika solusi optimal secara keseluruhan dapat dibangun dari solusi optimal dari submasalahnya
- Dynamic Programming Method
   - Top-down with memoization = Dalam pendekatan ini, kita mencoba untuk menyelesaikan masalah yang lebih besar dengan secara rekursif mencari solusi untuk submasalah yang lebih kecil. Setiap kali kita menyelesaikan sebuah submasalah, kita menyimpan hasilnya sehingga kita tidak perlu menyelesaikannya lagi jika dipanggil beberapa kali. Sebagai gantinya, kita dapat langsung mengembalikan hasil yang disimpan. Teknik penyimpanan hasil submasalah yang sudah diselesaikan ini disebut Memoisasi.
   - Bottom-up with Tabulation = kebalikan dari pendekatan top-down dan menghindari rekursi. Dalam pendekatan ini, kita menyelesaikan masalah "bottom-up" (artinya dengan menyelesaikan semua submasalah terkait terlebih dahulu). Ini biasanya dilakukan dengan mengisi tabel n-dimensi. Berdasarkan hasil dalam tabel, solusi untuk masalah top/asli kemudian dihitung.