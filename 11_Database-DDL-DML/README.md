# (10) Database - DDL - DML

3 Point Penting :
1. Entitas dan Attribute
- Entitas = Membahas mengenai Database Schema, pastinya tidak lepas dari yang namanya Entitas, Entitas mengacu pada objek yang dapat diidentifikasi secara unik dan memiliki atribut yang mendefinisikan karakteristik atau sifatnya. Entitas dalam database sering kali merepresentasikan objek nyata dalam dunia nyata, seperti orang, tempat, barang, atau konsep, yang perlu disimpan dan dikelola dalam sistem basis data. Salah satu contoh entitas adalah karyawan.
- Attribute = karakteristik atau sifat yang didefinisikan untuk suatu entitas. Mereka merepresentasikan data yang berkaitan dengan entitas tersebut dan membantu dalam mendeskripsikan atau mengidentifikasi entitas tersebut. Setiap entitas dalam database memiliki sejumlah atribut yang mewakili berbagai aspek atau informasi yang relevan tentang entitas tersebut. Misalnya, jika kita memiliki entitas "karyawan" dalam sebuah basis data HR, atribut-atribut yang mungkin terkait dengan entitas tersebut bisa mencakup: nama, alamat, gaji, dll.
2. Relasi
- One to One, salah satu contoh relasi one to one yaitu 1 user hanya memiliki 1 foto profil
- One to Many, salah satu contoh relasi one to many yaitu 1 user bisa memiliki banyak tweets
- Many to Many, salah satu contoh relasi many to many yaitu user dengan product makanan
3. Perintah SQL DDL dan DML
- DDL (Data Definition Language), DDL digunakan untuk mendefinisikan struktur dan karakteristik dari objek-objek basis data. Beberapa statement DDL yaitu CREATE DATABASE, USE, CREATE TABLE, DROP TABLE, dan RENAME TABLE
- DML (Data Manipulation Language), DML merupakan perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database. Beberapa statement DML yaitu INSERT INTO, SELECT, UPDATE, DELETE FROM. 