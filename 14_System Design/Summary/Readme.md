# System Design
Desain Modular: Microservices memecah sistem besar menjadi layanan-layanan kecil yang independen. Setiap layanan menangani fungsi spesifik, sehingga lebih mudah untuk dikembangkan, diuji, dan diperbaiki tanpa mengganggu sistem lain.

Komunikasi via API: Microservices berinteraksi satu sama lain melalui API atau pesan, bukan langsung memanggil fungsi di dalam satu aplikasi besar, sehingga memberikan fleksibilitas dan skalabilitas.

Skalabilitas dan Toleransi Kesalahan: Karena layanan terpisah, setiap microservice bisa diskalakan atau diperbaiki secara individual. Jika satu layanan gagal, yang lain tetap berjalan, meningkatkan keandalan sistem.