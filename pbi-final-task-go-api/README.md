# Final Task Project Based Internship BTPN Syariah X Rakamin

## oleh [Gaddy Juanda Sam](https://github.com/Gaddy661)

### Pendahuluan

API ini dibuat untuk menyelesaikan Final Task dari Project Based Internship BTPN Syariah X Rakamin. API ini digunakan untuk mengatur foto pengguna, seperti upload, edit, dan hapus foto.

### Fitur-Fitur API

* Pengguna dapat melakukan register dan login ke akun masing-masing.

* Pengguna yang telah di-otentikasi hanya dapat meng-edit dan menghapus akun milik pengguna sendiri.

* Pengguna yang telah di-otentikasi hanya dapat melihat, meng-edit, dan menghapus foto-foto milik pengguna sendiri.

### Environment Variables

| Variables | Usage |
| --- | --- |
| PORT | Untuk mengatur nomor port localhost yang dipakai oleh API, saat ini adalah 3000 |
| SECRET | Untuk keperluan menghasilkan token dengan JWT, saat ini adalah "haikawankawan" |
| DB_STRING | Untuk menghubungkan API dengan database |

### Endpoint yang Tersedia

| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| POST | /users/register | Untuk membuat akun baru |
| POST | /users/login | Untuk login akun yang ada |
| PUT | /users/:userId | Untuk mengubah info akun milik pengguna |
| DELETE | /users/:userId | Untuk menghapus akun milik pengguna beserta foto-fotonya |
| POST | /photos | Untuk meng-upload sebuah foto |
| GET | /photos | Untuk mendapatkan semua foto milik pengguna |
| PUT | /photos/:photoId | Untuk meng-edit sebuah foto milik pengguna |
| DELETE | /photos/:photoId | Untuk menghapus sebuah foto milik pengguna |

### Teknologi yang Digunakan

* [PostgreSQL](https://www.postgresql.org/) adalah sistem manajemen basis data relasional yang gratis dan open-source yang menekankan pada perluasan dan kepatuhan terhadap SQL. Layanan hosting basis data yang digunakan saat pembuatan API ini adalah [ElephantSQL](https://www.elephantsql.com/).

* [Gin Web Framework](https://gin-gonic.com/) adalah sebuah framework web yang ditulis dalam bahasa Go. Framework ini memiliki fitur API seperti martini dengan performa hingga 40 kali lebih cepat berkat httprouter.

* [GORM (Go Object Relational Mapper)](https://gorm.io/) adalah library populer untuk bahasa pemrograman Go yang menyederhanakan bekerja dengan basis data.

* [jwt-go](https://github.com/golang-jwt/jwt) adalah library Go yang populer untuk bekerja dengan JSON Web Token (JWT) yang digunakan untuk keperluan otentikasi pengguna.

* [govalidator](https://github.com/asaskevich/govalidator) adalah sekumpulan tools untuk memvalidasi dan membersihkan data dalam bahasa pemrograman Go.