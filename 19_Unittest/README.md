# (19) Unittesting

3 Point Penting :
1. Unit testing adalah proses pengujian perangkat lunak di mana bagian-bagian kecil atau "unit" dari program dites secara terisolasi untuk memastikan bahwa masing-masing unit berfungsi sebagaimana mestinya. Tujuan utama dari unit testing adalah untuk memverifikasi bahwa setiap unit kode bekerja dengan benar sesuai dengan spesifikasi yang telah ditentukan sebelumnya. 
2. Runner dan Mocking
- Runner dalam konteks unit testing adalah bagian dari kerangka kerja pengujian yang bertanggung jawab untuk mengeksekusi serangkaian tes dan melaporkan hasilnya. Runner biasanya digunakan untuk menjalankan seluruh rangkaian tes dalam suatu paket atau proyek pengujian. Di dalam Go, runner ini dikelola oleh paket testing dan dieksekusi menggunakan perintah go test.
- Mocking adalah teknik yang digunakan dalam unit testing untuk membuat objek palsu yang mensimulasikan perilaku objek nyata. Objek palsu ini digunakan untuk menggantikan objek nyata yang mungkin sulit atau tidak praktis untuk diuji dalam lingkungan pengujian. Dalam Go, mocking dapat dilakukan menggunakan paket seperti gomock atau dengan membuat struktur atau interface palsu secara manual.
3. Coverage dalam konteks unit testing mengacu pada seberapa banyak dari kode program yang telah diuji oleh serangkaian tes. Ini mengukur sejauh mana tes telah menjalankan kode program, memberikan gambaran tentang seberapa baik kode tersebut telah diuji. Terdapat beberapa jenis coverage yang umum digunakan dalam unit testing:
- Statement Coverage: Mengukur seberapa banyak statement kode yang dieksekusi oleh tes.
- Branch Coverage: Mengukur seberapa banyak percabangan kode yang dieksekusi oleh tes, termasuk if, else, switch, dan sebagainya.
- Function Coverage: Mengukur seberapa banyak fungsi atau metode yang dipanggil oleh tes.
- Line Coverage: Mengukur seberapa banyak baris kode yang dieksekusi oleh tes.
- Package Coverage: Mengukur seberapa banyak dari package yang telah diuji oleh tes.