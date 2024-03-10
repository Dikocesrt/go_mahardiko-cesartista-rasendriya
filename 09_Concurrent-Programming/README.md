# (09) Concurrent Programming

3 Point Penting :
1. Sequential Program, Parallel Program, dan Concurrent Program
- Sequential = Dalam program sequential, sebelum tugas baru dimulai, tugas sebelumnya harus selesai terlebih dahulu.
- Parallel = Dalam program paralel, beberapa tugas dapat dieksekusi secara bersamaan. (memerlukan mesin multi-core)
- Concurrent = Dalam program concurrent, beberapa tugas dapat dieksekusi secara independen dan mungkin tampak dilakukan secara bersamaan. Namun, program concurrent ini juga memungkinkan paralelisme, Concurrent dalam Go (goroutines) membuatnya mudah untuk membangun program paralel yang memanfaatkan mesin dengan beberapa prosesor (kebanyakan mesin saat ini).
2. 3 Fitur Go
- Goroutine = fungsi atau metode yang berjalan secara konkuren (independen) dengan fungsi atau metode lainnya. Pada goroutine terdapat gomaxprocs, gomaxprocs digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk eksekusi Goroutine dalam program Go tertentu.
- Channel = objek komunikasi yang digunakan oleh goroutine untuk berkomunikasi satu sama lain.
- Select = select memudahkan pengendalian komunikasi data melalui satu atau banyak channel.
3. Race Condition
Race condition adalah ketika 2 thread mengakses memori pada waktu yang sama, salah satunya sedang menulis. Race condition terjadi karena akses yang tidak disinkronkan ke memori bersama. 3 cara untuk mengatasi race condition adalah
- Blocking dengan Waitgroups = cara ini merupakan cara yang paling straigtforward untuk menyelesaikan race condition adalah dengan memblokir akses baca sampai operasi tulis telah selesai.
- Channel Blocking = channel blocking memungkinkan goroutine kita untuk melakukan sinkronisasi satu sama lain untuk sementara waktu.
- Mutex = Ketika kita tidak peduli tentang urutan pembacaan dan penulisan, kita hanya memerlukan agar mereka tidak terjadi secara bersamaan. Jika ini terdengar seperti kasus penggunaan Anda, maka kita harus mempertimbangkan untuk menggunakan mutex.