# (15) RESTful API

3 Point Penting :
1. API, API (Application Programming Interface) adalah serangkaian aturan dan protokol yang memungkinkan berbagai aplikasi perangkat lunak untuk berkomunikasi dan berinteraksi satu sama lain.
2. RESTful, REST adalah REpresentation State Transfer, REST menggunakan HTTP Protocol, contohnya : https://www.instagram.com/users/dikocesrt_, Beberapa contoh HTTP request methos yaitu seperti GET, POST, PUT, DELETE, HEAD, OPTION, PATCH. Request dan Response dapat berformat JSON, XML, atau SOAP. Beberapa best practice REST API Design adalah sebagai berikut:
- Use Nouns Instead of Verbs, contohnya seperti GET /books/123
- Naming Using Plural Nouns, contohnya seperti GET /cars/123
- Use Resource Nesting to show relations or hierarchy, contohnya seperti /users -> /users/123 -> /users/123/orders -> /users/123/orders/0001
- Handle Trailing Slashes Gracefully
- Filtering, Sorting, Paging, and Field Selection
3. POSTMAN, Postman adalah klien HTTP untuk menguji layanan web. Postman memudahkan pengujian, pengembangan, dan dokumentasi API dengan memungkinkan pengguna untuk dengan cepat menyusun permintaan HTTP yang sederhana maupun kompleks.