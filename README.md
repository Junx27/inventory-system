## Menjalankan proyek

Langkah pertama clone repository dengan memasukan perintah sebagai berikut

```bash
git clone https://github.com/Junx27/inventory-system
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
docker-compose run --rm go-seeder
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

## Preview endpoint program

## 1. Products

## 1.1. Add Products

Untuk menambahkan data produk dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `POST`:

```bash
http://localhost:8080/api/v1/products
```

Contoh memasukan data format `json`:

```json
{
  "name": "New Product 2",
  "price": 100,
  "category": "new",
  "description": "This is a test product"
}
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "message": "Product added successfully",
  "product": {
    "id": 5,
    "name": "New Product 2",
    "price": 100,
    "category": "new",
    "description": "This is a test product",
    "image": "",
    "inventory": {
      "id": 0,
      "qty": 0,
      "location": ""
    },
    "orders": null
  }
}
```

## 1.2. Get All Products

Untuk mendapatkan data semua produk dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `GET`:

```bash
http://localhost:8080/api/v1/products
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "products": [
    {
      "id": 1,
      "name": "Product A",
      "price": 100,
      "category": "Category A",
      "description": "Description for Product A",
      "image": "image_a.jpg",
      "inventory": {
        "id": 1,
        "qty": 50,
        "location": "Warehouse A"
      },
      "orders": [
        {
          "id": 1,
          "qty": 5
        },
        {
          "id": 2,
          "qty": 9
        }
      ]
    },
    {
      "id": 2,
      "name": "Product B",
      "price": 200,
      "category": "Category B",
      "description": "Description for Product B",
      "image": "image_b.jpg",
      "inventory": {
        "id": 2,
        "qty": 30,
        "location": "Warehouse B"
      },
      "orders": [
        {
          "id": 3,
          "qty": 2
        }
      ]
    },
    {
      "id": 3,
      "name": "Product C",
      "price": 300,
      "category": "Category C",
      "description": "Description for Product C",
      "image": "image_c.jpg",
      "inventory": {
        "id": 3,
        "qty": 100,
        "location": "Warehouse C"
      },
      "orders": [
        {
          "id": 4,
          "qty": 7
        },
        {
          "id": 5,
          "qty": 20
        }
      ]
    },
    {
      "id": 4,
      "name": "Product D",
      "price": 150,
      "category": "Category D",
      "description": "Description for Product D",
      "image": "image_d.jpg",
      "inventory": {
        "id": 4,
        "qty": 20,
        "location": "Warehouse D"
      },
      "orders": [
        {
          "id": 6,
          "qty": 1
        }
      ]
    }
  ]
}
```

## 1.3. Get Products By ID

Untuk mendapatkan data produk berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `GET`:

```bash
http://localhost:8080/api/v1/products/1
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "product": {
    "id": 1,
    "name": "Product A",
    "price": 100,
    "category": "Category A",
    "description": "Description for Product A",
    "image": "image_a.jpg",
    "inventory": {
      "id": 1,
      "qty": 50,
      "location": "Warehouse A"
    },
    "orders": [
      {
        "id": 1,
        "qty": 5
      },
      {
        "id": 2,
        "qty": 9
      }
    ]
  }
}
```

## 1.4. Update Products By ID

Untuk mengupdate produk berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `PUT`:

```bash
http://localhost:8080/api/v1/products/1
```

dan memasukan data `json` sebagai berikut:

```json
{
  "name": "New Product 1",
  "price": 100,
  "category": "new",
  "description": "This is a test product"
}
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "message": "Product updated successfully",
  "product": {
    "id": 1,
    "name": "New Product 1",
    "price": 100,
    "category": "new",
    "description": "This is a test product",
    "image": "image_a.jpg",
    "inventory": {
      "id": 0,
      "qty": 0,
      "location": ""
    },
    "orders": null
  }
}
```

Mengecek perubahan yang telah dibuat dengan mendapatkan semua data produk:

```json
{
  "products": [
    {
      "id": 2,
      "name": "Product B",
      "price": 200,
      "category": "Category B",
      "description": "Description for Product B",
      "image": "image_b.jpg",
      "inventory": {
        "id": 2,
        "qty": 30,
        "location": "Warehouse B"
      },
      "orders": [
        {
          "id": 3,
          "qty": 2
        }
      ]
    },
    {
      "id": 3,
      "name": "Product C",
      "price": 300,
      "category": "Category C",
      "description": "Description for Product C",
      "image": "image_c.jpg",
      "inventory": {
        "id": 3,
        "qty": 100,
        "location": "Warehouse C"
      },
      "orders": [
        {
          "id": 4,
          "qty": 7
        },
        {
          "id": 5,
          "qty": 20
        }
      ]
    },
    {
      "id": 4,
      "name": "Product D",
      "price": 150,
      "category": "Category D",
      "description": "Description for Product D",
      "image": "image_d.jpg",
      "inventory": {
        "id": 4,
        "qty": 20,
        "location": "Warehouse D"
      },
      "orders": [
        {
          "id": 6,
          "qty": 1
        }
      ]
    },
    {
      "id": 1,
      "name": "New Product 1",
      "price": 100,
      "category": "new",
      "description": "This is a test product",
      "image": "image_a.jpg",
      "inventory": {
        "id": 1,
        "qty": 50,
        "location": "Warehouse A"
      },
      "orders": [
        {
          "id": 1,
          "qty": 5
        },
        {
          "id": 2,
          "qty": 9
        }
      ]
    }
  ]
}
```

## 1.5. Delete Products By ID

Untuk menghapus produk berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `DELETE`:

```bash
http://localhost:8080/api/v1/products/1
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "message": "Product deleted successfully"
}
```

## 1.6 Upload Image Products By ID

Untuk mengupload image produk berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `POST`:

```bash
http://localhost:8080/api/v1/products/1/upload-image
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "image": "uploads/products/1.jpg",
  "message": "Image uploaded successfully"
}
```

## 1.7. Download Image Products By ID

```bash
http://localhost:8080/api/v1/products/1/download-image
```

## 2. Inventories

## 2.1. Add Inventories

Untuk menambahkan data inventories dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `POST`:

```bash
http://localhost:8080/api/v1/inventories
```

Contoh data `json` yang dimasukan:

```json
{
  "location": "Warehose B",
  "product_id": 1,
  "qty": 100
}
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "inventory": {
    "id": 5,
    "product_id": 1,
    "qty": 100,
    "location": "Warehose B",
    "product": {
      "id": 0,
      "name": "",
      "price": 0,
      "category": "",
      "description": "",
      "image": ""
    }
  },
  "message": "Inventory created successfully"
}
```

## 2.2. Get All Inventories

Untuk mendapatkan data semua inventories dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `GET`:

```bash
http://localhost:8080/api/v1/inventories
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "inventories": [
    {
      "id": 2,
      "qty": 30,
      "location": "Warehouse B"
    },
    {
      "id": 3,
      "qty": 100,
      "location": "Warehouse C"
    },
    {
      "id": 4,
      "qty": 20,
      "location": "Warehouse D"
    }
  ],
  "message": "Inventories fetched successfully"
}
```

## 2.3. Get Inventories By ID

Untuk mendapatkan data inventories berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `GET`:

```bash
http://localhost:8080/api/v1/inventories/1
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "inventory": {
    "id": 1,
    "location": "Warehouse A",
    "product": {
      "id": 1,
      "name": "Product A",
      "price": 100,
      "category": "Category A",
      "description": "Description for Product A",
      "image": ""
    },
    "product_id": 1,
    "qty": 50
  },
  "message": "Inventory fetched successfully"
}
```

## 2.4. Update Inventories By ID

Untuk mengupdate inventories berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `PUT`:

```bash
http://localhost:8080/api/v1/inventories/1
```

Contoh data `json` yang dimasukan:

```json
{
  "location": "Warehose B Update",
  "product_id": 1,
  "qty": 100
}
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "inventory": {
    "id": 5,
    "product_id": 1,
    "qty": 100,
    "location": "Warehose B Update",
    "product": {
      "id": 0,
      "name": "",
      "price": 0,
      "category": "",
      "description": "",
      "image": ""
    }
  },
  "message": "Inventory updated successfully"
}
```

## 2.5. Delete Inventories By ID

Untuk menghapus inventories berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `DELETE`:

```bash
http://localhost:8080/api/v1/inventories/1
```

Maka data yang ditampilkan adalah sebagai berikut:

```json
{
  "message": "Inventory deleted successfully"
}
```

## 2. Orders

## 2.1. Add Orders

Untuk menambahkan data orders dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `POST`:

```bash
http://localhost:8080/api/v1/orders
```

Maka data yang ditampilkan adalah sebagai berikut:

## 2.2. Get All Orders

Untuk mendapatkan data semua orders dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `GET`:

```bash
http://localhost:8080/api/v1/orders
```

Maka data yang ditampilkan adalah sebagai berikut:

## 2.3. Get Orders By ID

Untuk mendapatkan data orders berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `GET`:

```bash
http://localhost:8080/api/v1/orders/1
```

Maka data yang ditampilkan adalah sebagai berikut:

## 2.4. Update Orders By ID

Untuk mengupdate orders berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `PUT`:

```bash
http://localhost:8080/api/v1/orders/1
```

Maka data yang ditampilkan adalah sebagai berikut:

## 2.5. Delete Orders By ID

Untuk menghapus orders berdasarkan id dapat menjalankan perintah sebagai berikut pada POSTMAN dengan method `DELETE`:

```bash
http://localhost:8080/api/v1/orders/1
```

Maka data yang ditampilkan adalah sebagai berikut:
