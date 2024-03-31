# (18) Middleware

3 Point Penting :
1. Pengertian Middleware, sebuah konsep yang memungkinkan pengguna untuk menambahkan fungsionalitas tambahan ke dalam aplikasi web mereka. Middleware berperan sebagai jembatan antara permintaan yang masuk dan tanggapan yang keluar. Setiap kali permintaan HTTP masuk ke aplikasi, middleware dapat menangani permintaan tersebut, melakukan operasi tertentu, dan kemudian meneruskannya ke middleware berikutnya atau ke handler akhir yang menangani permintaan.
2. Dalam Echo (Go), terdapat dua metode yang digunakan untuk menerapkan middleware: Use() dan Pre(). Meskipun keduanya memiliki tujuan yang serupa, yaitu menambahkan fungsionalitas tambahan ke dalam aplikasi, ada perbedaan penting antara keduanya:
- Use()
	- Penjelasan, Metode Use() digunakan untuk mendaftarkan middleware yang akan dijalankan sebelum handler utama.
	- Cara Kerja, Middleware yang didaftarkan dengan Use() akan dijalankan secara serentak pada setiap permintaan sebelum handler utama dieksekusi.
	- Contoh, e.Use(middleware.Logger())
- Pre()
	- Penjelasan, Metode Pre() digunakan untuk mendaftarkan middleware yang akan dijalankan sebelum middleware lainnya, tetapi setelah middleware yang sudah ada.
	- Cara Kerja, Middleware yang didaftarkan dengan Pre() akan dijalankan sebelum middleware lainnya, tetapi setelah middleware yang sudah ada dengan Use() dijalankan.
	- Contoh, e.Pre(middleware.AddTrailingSlash())
3. Jenis Jenis Middleware Pada Echo
- JWT
	- Pengertian, JWT (JSON Web Token) adalah standar terbuka (RFC 7519) yang mendefinisikan format token yang aman untuk mentransfer klaim antara dua pihak.
	- Fungsi, Middleware JWT memverifikasi dan mengekstrak token JWT dari header permintaan, memvalidasi token, dan mengizinkan atau menolak akses berdasarkan klaim yang terkandung dalam token.
- Logger
	- Pengertian, Middleware Logger digunakan untuk mencatat informasi permintaan HTTP seperti metode, path, kode status, waktu respons, dan lainnya.
	- Fungsi, Merekam log permintaan dan tanggapan ke konsol atau penyimpanan log lainnya, yang membantu dalam pemantauan dan pemecahan masalah.
- Rate Limiter
	- Pengertian, Middleware Rate Limiter digunakan untuk membatasi jumlah permintaan yang diterima oleh aplikasi dalam satu periode waktu tertentu.
	- Fungsi, Memungkinkan pengguna untuk mengatur tingkat batasan permintaan berdasarkan IP pengguna, jenis permintaan, atau kriteria lainnya untuk mencegah serangan DDoS atau penggunaan berlebihan sumber daya server.