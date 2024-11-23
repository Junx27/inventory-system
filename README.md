## Menjalankan proyek

Langkah pertama clone repository dengan memasukan perintah sebagai berikut

```bash
cd inventory-system
```

Menjalankan program dengan menggunakan Docker, pastikan lokal komputer sudah terinstal Docker jika belum dapat mengunjungi website `https://docs.docker.com/engine/install/` sesuaikan dengan spesifikasi sistem operasi. Jika sudah dapat menjalankan dengan perintah sebagai berikut:

## 1. Menjalankan program utama dengan menggunakan Docker

Masukan perintah sebagai berikut untuk menjalankan program utama:

```bash
 docker-compose up --build -d
```

maka pada terminal akan menampilkan notifikasi sebagai berikut:

```bash
✔ Network inventory-system_default         Created                                                                                                                 0.0s
 ✔ Container inventory-system-postgres-1    Started                                                                                                                 0.3s
 ✔ Container inventory-system-go-app-1      Started                                                                                                                 0.6s
 ✔ Container inventory-system-go-migrate-1  Started                                                                                                                 0.5s
 ✔ Container inventory-system-go-seeder-1   Started                                                                                                                 0.5s
 ✔ Container inventory-system-go-drop-1     Started
```

## 2. Menjalankan migrasi tabel database dengan menggunakan Docker

Masukan perintah berikut untuk menjalankan migrasi tabel database dengan menggunakan servis sebagai berikut:

```bash
 docker-compose run --rm go-migrate
```

## 3. Menjalankan drop tabel database dengan menggunakan Docker

Masukan perintah drop table untuk menghapus tabel yang sudah ada pada databae, lakukan perintah ini jika perlu dengan menggunakan servis sebagai berikut:

```bash
docker-compose run --rm go-drop
```

## 4. Menjalankan migrasi database dengan menggunakan Docker

Masukan perintah berikut untuk melakukan seeder ke dalam database menggunakan data sementara, jika tabel sudah di drop maka lakukan proses migrasi agar proses seeder berhasil dengan menggunakan servis sebagai berikut:

```bash
docker-compose run --rm go-seed
```

## 5. Menghentikan program dengan menggunakan Docker

Masukan perintah berikut untuk menghentikan program

```bash
docker-compose down
```

maka pada terminal akan menampilkan notifikasi sebagai berikut:

```bash
 Container inventory-system-go-drop-1     Removed                                                                                                                 0.0s
 ✔ Container inventory-system-go-seeder-1   Removed                                                                                                                 0.0s
 ✔ Container inventory-system-go-migrate-1  Removed                                                                                                                 0.0s
 ✔ Container inventory-system-go-app-1      Removed                                                                                                                 0.2s
 ✔ Container inventory-system-postgres-1    Removed                                                                                                                 0.1s
 ✔ Network inventory-system_default         Removed
```

## Preview Docker
