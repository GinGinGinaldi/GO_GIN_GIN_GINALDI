# Docker

1. Definisi dan Fungsi
Docker volumes adalah cara untuk menyimpan data yang dihasilkan dan digunakan oleh kontainer Docker. Mereka memungkinkan data untuk dipersistenkan secara terpisah dari siklus hidup kontainer, sehingga data tetap ada meskipun kontainer dihentikan atau dihapus.
2. Pembuatan dan Penggunaan
Volume dapat dibuat menggunakan perintah docker volume create dan dapat dipasang ke kontainer dengan menggunakan flag -v atau --mount. Ini memungkinkan beberapa kontainer untuk berbagi volume yang sama, sehingga memudahkan kolaborasi antar aplikasi.
3. Kasus Penggunaan
Volumes sering digunakan untuk menyimpan data yang perlu bertahan lama, seperti file database atau konfigurasi aplikasi. Dengan menggunakan volumes, kita dapat memastikan bahwa data penting tidak hilang saat kontainer di-restart atau dihapus.