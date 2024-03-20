# (12) CLI

3 Point Penting :
1. CLI, CLI atau Command Line Interface adalah antarmuka teks yang cepat dan kuat yang digunakan oleh pengembang untuk berkomunikasi dengan komputer secara lebih efektif dan efisien untuk menyelesaikan beragam tugas. CLI ini memiliki beberapa keuntungan dibandingkan GUI, yaitu :
- Kontrol granular dari sistem operasi atau aplikasi.
- Manajemen yang lebih cepat dari sejumlah besar sistem operasi.
- Kemampuan untuk menyimpan skrip untuk mengotomatisasi tugas-tugas reguler.
- Antarmuka yang memungkinkan pengetahuan untuk membantu dalam pemecahan masalah, seperti masalah koneksi jaringan.
2. Shell Most Popular Commands
- Directory Commands
	- pwd = Menampilkan direktori kerja saat ini.
	- ls = Menampilkan daftar file dan direktori dalam direktori saat ini.
	- mkdir = Membuat direktori baru.
	- cd = Berpindah ke direktori tertentu.
	- rm = Menghapus file atau direktori.
	- cp = Menyalin file atau direktori dari satu lokasi ke lokasi lain.
	- mv = Memindahkan atau mengubah nama file atau direktori.
	- ln = Membuat tautan (link) antara file atau direktori.
- Files Commands
	- create
		- touch = Membuat file kosong atau mengubah tanggal modifikasi file.
	- view
		- head = Menampilkan beberapa baris pertama dari sebuah file.
		- cat = Menggabungkan dan menampilkan isi dari satu atau beberapa file.
		- tail = Menampilkan beberapa baris terakhir dari sebuah file.
		- less = Menampilkan isi file dengan kemampuan navigasi halaman dan mundur.
	- editor
		- vim = Editor teks yang kuat dan fleksibel dengan antarmuka mode teks dan mode perintah.
		- nano = Editor teks sederhana dengan antarmuka baris perintah yang ramah pengguna.
	- permission
		- chown = Mengubah kepemilikan (owner) dari file atau direktori.
		- chmod = Mengubah izin (permissions) file atau direktori.
	- different
		- diff = Membandingkan dua file dan menampilkan perbedaan di antara keduanya.
- Network Commands
	- ping = Memeriksa koneksi jaringan dengan mengirim paket ke tujuan tertentu dan menunggu respons balik.
	- ssh = Mengakses dan mengelola server atau komputer jarak jauh melalui protokol SSH (Secure Shell).
	- curl = Mengunduh atau mengirim data melalui URL menggunakan berbagai protokol seperti HTTP, HTTPS, FTP, dan lainnya.
	- wget = Mengunduh file atau konten dari internet menggunakan protokol HTTP, HTTPS, atau FTP.
	- netstat = Menampilkan informasi tentang koneksi jaringan, tabel routing, dan statistik antarmuka jaringan pada sistem.
- Utility Commands
	- man = Menampilkan manual (dokumentasi) untuk perintah atau topik tertentu.
	- env = Menampilkan atau mengatur variabel lingkungan sistem.
	- echo = Menampilkan teks atau nilai variabel ke output standar.
	- sudo = Mengizinkan pengguna menjalankan perintah dengan hak akses superuser (root).
	- grep = Mencari teks dalam file atau output menggunakan pola pencarian yang ditentukan.
3. Shell Script, pada shell kita juga dapat membuat file script yang berekstensi .sh yang dapat di compile menggunakan shell, file sh ini memiliki fungsi yang cukup banyak salah satunya mengakses suatu endpoint dan mendapatkan response dari endpoint tersebut, lalu response tersebut dapat di save ke file yang berada di current directory