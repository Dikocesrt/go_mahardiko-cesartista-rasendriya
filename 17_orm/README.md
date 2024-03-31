# (17) ORM dan MVC

3 Point Penting :
1. Pengertian ORM dan GORM
- Pengertian ORM, Object-Relational Mapping (ORM) adalah sebuah konsep dalam pengembangan perangkat lunak yang menghubungkan objek dalam bahasa pemrograman dengan entitas dalam basis data relasional. Ini memungkinkan pengembang untuk berinteraksi dengan basis data menggunakan objek dan bahasa pemrograman, tanpa harus menulis kueri SQL secara eksplisit.
- Pengertian GORM, GORM adalah salah satu ORM yang populer untuk bahasa pemrograman Go. Ini menyediakan alat dan fitur untuk memetakan struktur data Go ke dalam basis data relasional dan melakukan operasi CRUD.
2. Kelebihan dan Kekurangan GORM
- Kelebihan
	- Mudah Digunakan: GORM memiliki sintaks yang mudah dipahami dan digunakan, membuatnya cocok untuk pengembang yang baru mengenal Go dan ORM.
	- Dokumentasi yang Baik: GORM memiliki dokumentasi yang lengkap dan bermanfaat, yang memudahkan pengguna untuk mempelajari dan menggunakan framework ini.
	- Komunitas yang Aktif: GORM didukung oleh komunitas yang aktif, yang berarti pengguna dapat dengan mudah menemukan dukungan dan sumber daya tambahan.
- Kekurangan
	- Performa: Seperti halnya dengan banyak ORM, performa GORM dapat menjadi masalah dalam aplikasi yang memerlukan kueri basis data yang sangat dioptimalkan.
	- Ketergantungan: Penggunaan GORM menambahkan ketergantungan pada library eksternal, yang dapat meningkatkan kompleksitas proyek.
3. MVC, Model-View-Controller (MVC) adalah sebuah pola desain (design pattern) yang digunakan dalam pengembangan perangkat lunak untuk memisahkan logika aplikasi menjadi tiga komponen yang terpisah: Model, View, dan Controller. Ini membantu dalam mengorganisir kode dan memisahkan tugas-tugas yang berbeda dalam aplikasi.
- Model, Model merepresentasikan struktur data dan logika bisnis dari aplikasi. Ini berinteraksi dengan basis data untuk melakukan operasi CRUD (Create, Read, Update, Delete) dan memanipulasi data.
- View, View adalah tampilan dari data yang disajikan kepada pengguna. Ini menangani presentasi data kepada pengguna dalam bentuk yang dapat dipahami dan interaktif.
- Controller, Controller bertindak sebagai perantara antara Model dan View. Ini menerima input dari pengguna melalui View, memprosesnya menggunakan Model, dan kemudian mengirimkan output kembali ke View.