# (10) Clean Code

3 Point Penting :
1. Karakteristik
- Mudah dipahami = penamaan variable atau lainnya yang mudah dipahami merupakan karakteristik dari clean code
- Mudah dieja dan dicari = penamaan variable atau lainnya jangan di singkat singkat karena akan susah untuk dieja dan dicari
- Singkat namun mendeskripsikan konteks = penamaan penamaan variable atau lainnya dapat mendeskripsikan konteksnya
- Konsisten = penamaan variable atau lainnya konsisten seperti uppercase dan lowercasenya atau camelCase dan snake_casenya
- Hindari penambahan konteks yang tidak perlu = jangan ada pengulangan nama konteks yang sudah pasti
- Komentar = memang komentar merupakan hal yang penting agar membantu kode kita mudah untuk dipahami, namun terlalu banyak komentar dapat menyebabkan kode malah susah untuk dipahami
- Good function = jangan terlalu banyak argumen pada parameter, apabila memang banyak maka dijadikan object maka akan membuat kode semakin clean dan juga 1 fungsi hanya memiliki 1 fungsionalitas saja
- Gunakan konvensi = seperti menggunakan style guide dari airbnb atau google
- Formatting = menggunakan contoh formatting seperti lebar baris code 80 - 120 karakter, satu class 300 - 500 baris, menggunakan prettier, dll
2. Clean Code Principles
- KISS (Keep It So Simple) = Hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B, mengecek fungsi C, dst.
- DRY (Don't Repeat Yourself) = Duplikasi kode terjadi karena sering copy-paste, untuk menghindari duplikasi kode buatlah fungsi yang dapat digunakan secara berulang-ulang
3. Refactoring
Proses restrukturisasi kode yang dibuat, dengan cara merubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor