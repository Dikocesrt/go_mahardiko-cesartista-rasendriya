# (20) Clean Architecture

3 Point Penting :
1. Clean Architecture adalah sebuah pendekatan dalam pengembangan perangkat lunak yang menekankan pada pemisahan konsep dan tanggung jawab dalam sebuah aplikasi. Pendekatan ini didasarkan pada prinsip bahwa kode program harus tersusun dengan baik, mudah dimengerti, dan mudah untuk dikelola.
2. Pada clean architecture, sebuah aplikasi dibagi menjadi beberapa lapisan, yang masing-masing memiliki tanggung jawabnya sendiri. Lapisan-lapisan tersebut umumnya mencakup:
- Lapisan Domain: Merupakan inti dari aplikasi, berisi aturan bisnis dan logika aplikasi. Lapisan ini bersifat independen dari infrastruktur atau kerangka kerja tertentu.
- Lapisan Use Case (Application): Berisi implementasi dari fungsionalitas yang spesifik untuk aplikasi tersebut. Lapisan ini bertanggung jawab untuk menerapkan aturan bisnis dari lapisan domain ke dalam kasus penggunaan yang konkret.
- Lapisan Infrastruktur: Merupakan lapisan yang bertanggung jawab untuk berkomunikasi dengan sumber daya eksternal, seperti database, sistem file, atau layanan web eksternal. Lapisan ini terisolasi dari lapisan domain dan use case sehingga memungkinkan untuk melakukan perubahan infrastruktur tanpa harus memengaruhi bagian lain dari aplikasi.
3. Clean Architecture bertujuan untuk memisahkan kode menjadi komponen-komponen yang terpisah, sehingga memudahkan pengujian, pemeliharaan, dan pengembangan lebih lanjut. Dengan pendekatan ini, aplikasi menjadi lebih fleksibel dan mudah diubah tanpa harus mengganggu bagian lain dari sistem.