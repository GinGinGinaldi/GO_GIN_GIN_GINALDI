# Clean Architecture


1. Clean dan Hexagonal Architecture 
   Clean Architecture dan Hexagonal Architectureberfokus pada pemisahan logika bisnis dari detail implementasi. Keduanya menciptakan struktur aplikasi yang bersih, meningkatkan pemeliharaan dan pengujian. Clean Architecture membagi aplikasi ke dalam lapisan-lapisan, sedangkan Hexagonal Architecture menggunakan antarmuka (ports) untuk interaksi dengan dunia luar.

2. Context di Golang 
   Context di Golang penting untuk mengelola batas waktu dan pembatalan. Ini memungkinkan komunikasi yang efisien antar goroutines, serta membantu menghindari kebocoran sumber daya dalam aplikasi yang mengikuti Clean Architecture.

3. Dari MVC ke Clean Code
   Peralihan dari MVC ke Clean Architecturemelibatkan pemisahan yang lebih jelas antara logika bisnis dan antarmuka pengguna. Hal ini meningkatkan keterbacaan dan memudahkan perbaikan serta pengembangan kode yang lebih bersih.

4. Menulis Unit Test
   Unit testing adalah kunci dalam Clean Architecture, karena memisahkan logika bisnis dari implementasi memudahkan pengujian. Penggunaan interface dan dependency injection memungkinkan pengujian yang efektif, meningkatkan keandalan kode dengan mengidentifikasi dan memperbaiki bug lebih cepat.