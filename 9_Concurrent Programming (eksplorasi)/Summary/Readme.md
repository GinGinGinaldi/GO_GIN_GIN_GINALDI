# Concurrent Programming (eksplorasi)

1. Konkurensi vs. Sekuensial dan Paralel
Dalam pemrograman, konkurensi memungkinkan kita menjalankan beberapa tugas secara bersamaan dan berinteraksi, berbeda dari pemrograman sekuensial yang satu per satu dan paralel yang menggunakan beberapa prosesor. Contoh aplikasinya bisa dilihat di website "Big Search" yang harus menangani banyak permintaan sekaligus.
2. Goroutines dan Channels
Di Golang, kita menggunakan goroutines untuk menjalankan fungsi secara bersamaan. Untuk komunikasi antar goroutines, kita memanfaatkan channels yang aman untuk mengirim data. Dengan menggunakan select, kita juga bisa menunggu dan merespons data dari beberapa channel secara efisien.
3. Mengatasi Masalah Konkurensi
Salah satu tantangan yang kita hadapi adalah race condition, di mana beberapa goroutines mengakses data yang sama secara bersamaan. Untuk mencegahnya, kita bisa menggunakan mutex untuk mengunci akses ke data dan WaitGroups untuk memastikan semua goroutines selesai sebelum melanjutkan. Channels juga bisa bersifat blocking, yang membantu mengatur alur eksekusi dan mencegah kondisi balapan.