# (12) System Designs

3 Point Penting :
1. Diagram Design, Sebuah diagram adalah representasi simbolis dari informasi menggunakan teknik visualisasi. Beberapa diagram design tools adalah smartdraw, Lucidchart, Whimsical, draw.io, visio. Software design yang biasa dipakai adalah
- Flowchart, representasi grafis dari alur kerja atau proses yang menggunakan simbol-simbol standar untuk menggambarkan langkah-langkah serta hubungan antara mereka.
- Use Case, teknik pemodelan yang digunakan dalam rekayasa perangkat lunak untuk menggambarkan interaksi antara sistem dan aktor yang terlibat dalam menjalankan suatu skenario tertentu.
- ERD (Entity-Relationship Diagram) adalah representasi visual dari struktur data yang digunakan untuk menggambarkan hubungan antara entitas dalam suatu sistem informasi.
- High Level Architecture (HLA) adalah sebuah standar yang digunakan dalam simulasi komputer untuk memfasilitasi interoperabilitas antara berbagai sistem simulasi yang berbeda.
2. Karakteristik Sistem Distribusi, Job Queue, dan Load Balancing
- Karakteristik dari sistem terdistribusi yaitu:
	- Scalability, terdapat 2 macam jenis scalability yaitu:
		- Vertical Scale, vertical scale adalah peningkatan kinerja suatu sistem dengan menambah atau meningkatkan kapasitas sumber daya (seperti CPU, RAM, atau penyimpanan) pada satu mesin tunggal.
		- Horizontal Scale, pendekatan dalam meningkatkan kinerja sistem dengan menambahkan lebih banyak mesin atau instance, yang memungkinkan untuk meningkatkan kapasitas secara horizontal.
	- Reliability, kemampuan suatu sistem atau produk untuk beroperasi dengan konsistensi yang tinggi dan minim gangguan atau kegagalan selama periode waktu tertentu.
	- Availability, kemampuan suatu sistem atau layanan untuk tetap berfungsi dan tersedia bagi pengguna dalam rentang waktu yang diinginkan, tanpa adanya gangguan atau downtime yang signifikan.
	- Efficiency, kemampuan sistem untuk menggunakan sumber daya secara optimal dalam menangani permintaan, melakukan komunikasi, dan menjalankan tugas secara efisien di seluruh jaringan terdistribusi.
	- Manageability, kemampuan untuk dengan mudah mengelola, memantau, dan mengontrol sumber daya serta aktivitas sistem terdistribusi secara efisien dan efektif.
- Job Queue, struktur data yang digunakan dalam sistem komputer untuk mengatur dan mengelola antrian tugas atau pekerjaan yang harus dieksekusi. Dalam job queue, setiap tugas atau pekerjaan ditempatkan dalam antrian berdasarkan urutan kedatangan atau prioritasnya. Sistem menggunakan job queue untuk mengelola sumber daya dan alokasi waktu, memastikan bahwa tugas dieksekusi dengan benar sesuai dengan aturan yang ditetapkan.
- Load Balancing, teknik yang digunakan dalam sistem komputer untuk mendistribusikan beban kerja atau lalu lintas jaringan secara merata di antara beberapa sumber daya komputasi, seperti server, komputer, atau jaringan, dengan tujuan meningkatkan kinerja sistem, mencegah kelebihan beban pada satu titik, dan meningkatkan ketersediaan layanan. Load balancing biasanya digunakan dalam lingkungan dengan tingkat lalu lintas yang tinggi, seperti situs web, aplikasi online, dan jaringan besar, untuk memastikan ketersediaan layanan yang optimal dan responsif bagi pengguna.
3. Monolithic VS Microservice dan SQL VS No SQL
- Monolithic VS Microservice
	- Monolithic, Aplikasi monolit memiliki basis kode tunggal dengan beberapa modul. Modul-modulnya dibagi sebagai modul untuk fitur bisnis atau fitur teknis. Ini memiliki sistem pembangunan tunggal yang membangun seluruh aplikasi dan/atau dependensinya. Ini juga memiliki biner eksekusi tunggal atau biner yang dapat diimplementasikan.
	- Microservice, layanan yang dapat diimplementasikan secara independen yang dimodelkan di sekitar sebuah domain bisnis. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur menawarkan banyak opsi untuk memecahkan masalah yang mungkin Anda hadapi. Dengan demikian, arsitektur mikroservice didasarkan pada beberapa mikroservice yang bekerja sama.
- SQL VS No SQL
	- SQL, Database relasional terstruktur dan memiliki skema yang telah ditentukan sebelumnya seperti buku telepon yang menyimpan nomor telepon dan alamat.
	- No SQL, Database non-relasional tidak terstruktur, dan memiliki skema dinamis seperti folder-file yang menyimpan segala sesuatu mulai dari alamat dan nomor telepon seseorang hingga 'like' di Facebook dan preferensi belanja online mereka.