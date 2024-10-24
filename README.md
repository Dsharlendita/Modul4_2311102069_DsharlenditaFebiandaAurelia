# Modul 4 Prosedur

Modul 4 ini membahas pemrograman prosedural di Go (Golang), yang merupakan metode untuk menulis kode dengan menggunakan fungsi-fungsi untuk menyelesaikan berbagai tugas secara spesifik. Setiap fungsi dapat menerima input dan menghasilkan output, sehingga fungsi-fungsi tersebut dapat saling berkomunikasi dengan efisien.

## Penjelasan Kode

1. Program ini menghitung nilai permutasi dan kombinasi. Fungsi `faktorial` menghitung faktorial dari angka `n`, dengan hasil 1 jika `n` adalah 0, dan mengalikan angka dari 1 hingga `n` untuk nilai lainnya. Fungsi `permutasi` menggunakan rumus \( P(n, r) = \frac{n!}{(n-r)!} \) dan `kombinasi` menggunakan rumus \( C(n, r) = \frac{n!}{r!(n-r)!} \). Fungsi `main()` membaca empat angka dari pengguna dan menghitung permutasi serta kombinasi dari angka tersebut.

2. Program ini menentukan pemenang kompetisi berdasarkan waktu penyelesaian soal. Dengan konstanta `maxTime` yang diatur ke 301, fungsi `hitungSkor` menghitung jumlah soal yang diselesaikan dan total waktu yang dihabiskan. Fungsi `main()` meminta nama peserta dan waktu penyelesaian, yang dapat diakhiri dengan "Selesai". Setelah menghitung, program membandingkan hasil peserta dengan pemenang saat ini dan menampilkan nama pemenang, jumlah soal yang diselesaikan, dan total waktu yang dihabiskan.

3. Program ini mencetak deret angka berdasarkan aturan Conjecture Collatz (3n + 1). Fungsi `cetakDeret` mencetak nilai `n` hingga mencapai 1, di mana angka genap dibagi dua dan angka ganjil diperbarui menjadi `3n + 1`. Dalam fungsi `main()`, pengguna diminta memasukkan angka yang valid (lebih besar dari 0 dan kurang dari 1.000.000). Jika valid, fungsi `cetakDeret` dipanggil. Jika tidak, program akan meminta angka positif yang sesuai.
