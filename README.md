# UNISA Server

UNISA Server adalah aplikasi backend untuk manajemen data dengan API untuk registrasi pengguna, login, dan operasi CRUD pada entitas `ListTautan` dan `DataLeads`. Aplikasi ini menggunakan MongoDB sebagai database dan JWT untuk autentikasi.

## Table of Contents

- [Setup](#setup)
- [Endpoints](#endpoints)
  - [User Endpoints](#user-endpoints)
  - [ListTautan Endpoints](#listtautan-endpoints)
  - [DataLeads Endpoints](#dataleads-endpoints)
- [Authentication](#authentication)
- [Error Handling](#error-handling)
- [Contributing](#contributing)
- [License](#license)

## Setup

1. **Install Dependencies**
   ```bash
   git clone https://github.com/yourusername/unisa_server.git
   cd unisa_server
   ```
2. **Configuration**

   ```bash
    MONGODB_URI=your_mongodb_uri
    JWT_SECRET=your_secret_key
   ```

3. **Configuration**
   ```bash
    go run main.go
   ```

## Endpoints

## User Endpoints

**Register**

- Endpoint: POST /register
- Description: Mendaftar pengguna baru.
- Request Body:

  ```bash
  {
    "username": "user123",
    "password": "password123",
    "role": "user" // atau "admin"
  }
  ```

- Responses:

  - 201 Created

  ```bash
  {
    "status": true,
    "Status_Code": 201,
    "message": "User registered successfully",
    "data": null
  }
  ```

  - 400 Bad Request

  ```bash
  {
    "status": false,
    "Status_Code": 400,
    "message": "Invalid input"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error registering user"
  }
  ```

  **Login**

- Endpoint: POST /login
- Description: Login pengguna dan mendapatkan token JWT.
- Request Body:

  ```bash
  {
    "username": "user123",
    "password": "password123"
  }
  ```

- Responses:

  - 200 OK

  ```bash
  {
    "status": true,
    "Status_Code": 200,
    "message": "Login successful",
    "data": ["<jwt_token>"]
  }
  ```

  - 401 Unauthorized

  ```bash
  {
    "status": false,
    "Status_Code": 401,
    "message": "Invalid credentials"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error generating token"
  }
  ```

## ListTautan Endpoints

**ListTautan**

- Endpoint: GET /list_tautan/get
- Description: Mengambil entitas seluruh DataLeads.
- Query Parameter: id=<object_id>
- Responses:

  - 200 OK

  ```bash
  {
    "status": true,
    "Status_Code": 200,
    "message": "Data ditemukan",
    "data": {
      "ID": "<object_id>",
      "QRCodePath": "path/to/qrcode.png",
      "NamaProgram": "Program A",
      "NamaInstansi": "Instansi A",
      "NamaKegiatan": "Kegiatan A",
      "Alamat": "Alamat A",
      "NamaPIC": "PIC A",
      "NamaPICPTInstansi": "PIC PT A",
      "TanggalMulai": "2024-08-01",
      "TanggalAkhir": "2024-08-15"
    },
    ...
  }
  ```

  - 404 Not Found

  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Data tidak ditemukan
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error fetching ListTautan"
  }
  ```

**CreateListTautan**

- Endpoint: POST /list_tautan/create
- Description: Membuat entitas ListTautan baru dan menghasilkan kode QR.
- Request Body:

  ```bash
  {
    "NamaProgram": "Program A",
    "NamaInstansi": "Instansi A",
    "NamaKegiatan": "Kegiatan A",
    "Alamat": "Alamat A",
    "NamaPIC": "PIC A",
    "NamaPICPTInstansi": "PIC PT A",
    "TanggalMulai": "2024-08-01",
    "TanggalAkhir": "2024-08-15"
  }
  ```

- Responses:

  - 201 Created

  ```bash
  {
    "status": true,
    "Status_Code": 201,
    "message": "Update Success",
    "data": {
      "ID": "<object_id>",
      "QRCodePath": "path/to/qrcode.png",
      "NamaProgram": "Program A",
      "NamaInstansi": "Instansi A",
      "NamaKegiatan": "Kegiatan A",
      "Alamat": "Alamat A",
      "NamaPIC": "PIC A",
      "NamaPICPTInstansi": "PIC PT A",
      "TanggalMulai": "2024-08-01",
      "TanggalAkhir": "2024-08-15"
    }
  }
  ```

  - 400 Bad Request

  ```bash
  {
    "status": false,
    "Status_Code": 400,
    "message": "Invalid input"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error creating ListTautan"
  }
  ```

  **GetAllListTautan**

  ### Query Parameters

- Endpoint: GET /list_tautan
- Description:

  - page (optional): Halaman data yang ingin diambil. Default: 1.
  - limit (optional): Jumlah item per halaman. Default: 10, maksimum: 100.
  - search (optional): Kata kunci pencarian untuk menyaring data berdasarkan beberapa field.
  - sortBy (optional): Field yang digunakan untuk mengurutkan data secara abjad (A-Z).
  - sortOrder (optional): Urutan pengurutan data. Dapat berupa asc (ascending) atau desc (descending). Default: asc.

- Responses:

  - 200 OK

  ```bash
  {
    "status": true,
    "Status_Code": 200,
    "message": "Data ditemukan",
    "data": [{
      "ID": "<object_id>",
      "QRCodePath": "path/to/qrcode.png",
      "NamaProgram": "Program A",
      "NamaInstansi": "Instansi A",
      "NamaKegiatan": "Kegiatan A",
      "Alamat": "Alamat A",
      "NamaPIC": "PIC A",
      "NamaPICPTInstansi": "PIC PT A",
      "TanggalMulai": "2024-08-01",
      "TanggalAkhir": "2024-08-15"
    },
    ...
    ]
  }
  ```

  - 404 Not Found

  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Query tidak boleh kosong"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```

  **GetListTautan**

- Endpoint: GET /list_tautan/get
- Description: Mengambil entitas ListTautan berdasarkan ID.

- Responses:

  - 200 OK

  ```bash
  {
    "status": true,
    "Status_Code": 200,
    "message": "Data ditemukan",
    "data": {
      "ID": "<object_id>",
      "QRCodePath": "path/to/qrcode.png",
      "NamaProgram": "Program A",
      "NamaInstansi": "Instansi A",
      "NamaKegiatan": "Kegiatan A",
      "Alamat": "Alamat A",
      "NamaPIC": "PIC A",
      "NamaPICPTInstansi": "PIC PT A",
      "TanggalMulai": "2024-08-01",
      "TanggalAkhir": "2024-08-15"
    }
  }
  ```

  - 404 Not Found

  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Query tidak boleh kosong"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```

  **UpdateListTautan**

- Endpoint: PUT /list_tautan/update
- Description: Memperbarui entitas ListTautan berdasarkan ID dan update kode QR.
- Query Parameter: **id=<object_id>**
- Request Body:

  ```bash
  {
    "NamaProgram": "Program A",
    "NamaInstansi": "Instansi A",
    "NamaKegiatan": "Kegiatan A",
    "Alamat": "Alamat A",
    "NamaPIC": "PIC A",
    "NamaPICPTInstansi": "PIC PT A",
    "TanggalMulai": "2024-08-01",
    "TanggalAkhir": "2024-08-15"
  }
  ```

- Responses:
- 201 Created

```bash
{
  "status": true,
  "Status_Code": 201,
  "message": "Update Success",
  "data": {
    "ID": "<object_id>",
    "QRCodePath": "path/to/qrcode.png",
    "NamaProgram": "Program A",
    "NamaInstansi": "Instansi A",
    "NamaKegiatan": "Kegiatan A",
    "Alamat": "Alamat A",
    "NamaPIC": "PIC A",
    "NamaPICPTInstansi": "PIC PT A",
    "TanggalMulai": "2024-08-01",
    "TanggalAkhir": "2024-08-15"
  }
}
```

- 204 No Content
- 404 Not Found

```bash
{
  "status": false,
  "Status_Code": 404,
  "message": "Query tidak boleh kosong"
}
```

- 500 Internal Server Error

```bash
{
  "status": false,
  "Status_Code": 500,
  "message": "Terjadi kesalahan: [error_message]"
}
```

     **DeleteListTautan**

- Endpoint: DELETE /list_tautan/delete
- Description: Menghapus entitas ListTautan berdasarkan ID.
- Query Parameter: **id=<object_id>**
- Request Body:

  ```bash
  {
    "NamaProgram": "Program A",
    "NamaInstansi": "Instansi A",
    "NamaKegiatan": "Kegiatan A",
    "Alamat": "Alamat A",
    "NamaPIC": "PIC A",
    "NamaPICPTInstansi": "PIC PT A",
    "TanggalMulai": "2024-08-01",
    "TanggalAkhir": "2024-08-15"
  }
  ```

- Responses:
  - 204 No Content
  - 404 Not Found
  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Query tidak boleh kosong"
  }
  ```
  - 500 Internal Server Error
  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```
  **DeleteListTautan**
- Endpoint: DELETE /listtautan
- Description: Menghapus entitas ListTautan berdasarkan ID.
- Query Parameter: **id=<object_id>**
- Request Body:

  ```bash
  {
    "NamaProgram": "Program A",
    "NamaInstansi": "Instansi A",
    "NamaKegiatan": "Kegiatan A",
    "Alamat": "Alamat A",
    "NamaPIC": "PIC A",
    "NamaPICPTInstansi": "PIC PT A",
    "TanggalMulai": "2024-08-01",
    "TanggalAkhir": "2024-08-15"
  }
  ```

- Responses:
  - 204 No Content
  - 404 Not Found
  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Query tidak boleh kosong"
  }
  ```
  - 500 Internal Server Error
  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```

## DataLeads Endpoints

**GetAllDataLeads**

### Query Parameters

- Endpoint: GET /data_leads
- Description:

  - page (optional): Halaman data yang ingin diambil. Default: 1.
  - limit (optional): Jumlah item per halaman. Default: 10, maksimum: 100.
  - search (optional): Kata kunci pencarian untuk menyaring data berdasarkan beberapa field.
  - sortBy (optional): Field yang digunakan untuk mengurutkan data secara abjad (A-Z).
  - sortOrder (optional): Urutan pengurutan data. Dapat berupa asc (ascending) atau desc (descending). Default: asc.

- Responses:

  - 200 OK

  ```bash
  {
    "status": true,
    "Status_Code": 200,
    "message": "Data ditemukan",
    "data": [{
      "ID": "<object_id>",
      "NamaLengkap": "John Doe",
      "Alamat": "Jl. Example No. 123",
      "TempatLahir": "Jakarta",
      "TanggalLahir": "1990-01-01",
      "Email": "johndoe@example.com",
      "NoHp": "08123456789",
      "NamaPTInstansi": "PT Example",
      "Jabatan": "Manager"
    },
    ...
    ]
  }
  ```

  - 404 Not Found

  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Data tidak ditemukan"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```

  **GetListTautanById**

- Endpoint: GET /data_leads/get
- Description: Mengambil entitas DataLeads berdasarkan ID.

- Responses:

  - 200 OK

  ```bash
  {
    "status": true,
    "Status_Code": 200,
    "message": "Data ditemukan",
    "data": {
      "ID": "<object_id>",
      "QRCodePath": "path/to/qrcode.png",
      "NamaProgram": "Program A",
      "NamaInstansi": "Instansi A",
      "NamaKegiatan": "Kegiatan A",
      "Alamat": "Alamat A",
      "NamaPIC": "PIC A",
      "NamaPICPTInstansi": "PIC PT A",
      "TanggalMulai": "2024-08-01",
      "TanggalAkhir": "2024-08-15"
    }
  }
  ```

  - 404 Not Found

  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Query tidak boleh kosong"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```

  **CreateDataLeads**

- - Endpoint: POST /data_leads/create
- Description: Membuat entitas DataLeads baru.
- Request Body:

  ```bash
  {
    "NamaLengkap": "John Doe",
    "Alamat": "Jl. Example No. 123",
    "TempatLahir": "Jakarta",
    "TanggalLahir": "1990-01-01",
    "Email": "johndoe@example.com",
    "NoHp": "08123456789",
    "NamaPTInstansi": "PT Example",
    "Jabatan": "Manager"
  }
  ```

- Responses:

  - 201 Created

  ```bash
  {
    "status": true,
    "Status_Code": 201,
    "message": "DataLeads created successfully",
    "data": {
      "ID": "<object_id>",
      "NamaLengkap": "John Doe",
      "Alamat": "Jl. Example No. 123",
      "TempatLahir": "Jakarta",
      "TanggalLahir": "1990-01-01",
      "Email": "johndoe@example.com",
      "NoHp": "08123456789",
      "NamaPTInstansi": "PT Example",
      "Jabatan": "Manager"
    }
  }
  ```

  - 400 Bad Request

  ```bash
  {
    "status": false,
    "Status_Code": 400,
    "message": "Invalid input"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error creating DataLeads"
  }
  ```

**UpdateDataLeads**

- Endpoint: PUT /data_leads/update
- Description: Memperbarui entitas ListTautan berdasarkan ID dan update kode QR.
- Query Parameter: **id=<object_id>**
- Request Body:

  ```bash
  {
    "NamaLengkap": "John Doe Updated",
    "Alamat": "Jl. Example Updated No. 123",
    "TempatLahir": "Jakarta Updated",
    "TanggalLahir": "1990-01-02",
    "Email": "johndoeupdated@example.com",
    "NoHp": "08123456780",
    "NamaPTInstansi": "PT Example Updated",
    "Jabatan": "Senior Manager"
  }
  ```

- Responses:
- 201 Created

```bash
{
  "status": true,
  "Status_Code": 201,
  "message": "Update Success",
  "data": {
    "NamaLengkap": "John Doe Updated",
    "Alamat": "Jl. Example Updated No. 123",
    "TempatLahir": "Jakarta Updated",
    "TanggalLahir": "1990-01-02",
    "Email": "johndoeupdated@example.com",
    "NoHp": "08123456780",
    "NamaPTInstansi": "PT Example Updated",
    "Jabatan": "Senior Manager"
}
```

- 204 No Content
- 404 Not Found

```bash
{
  "status": false,
  "Status_Code": 404,
  "message": "Invalid input"
}
```

- 500 Internal Server Error

```bash
{
  "status": false,
  "Status_Code": 500,
  "message": "Error updating DataLeads"
}
```

**DeleteDataLeads**

- Endpoint: DELETE /data_leads/delete
- Description: : Menghapus entitas DataLeads berdasarkan ID.
- Query Parameter: **id=<object_id>**
- Request Body:

  ```bash
  {
    "NamaProgram": "Program A",
    "NamaInstansi": "Instansi A",
    "NamaKegiatan": "Kegiatan A",
    "Alamat": "Alamat A",
    "NamaPIC": "PIC A",
    "NamaPICPTInstansi": "PIC PT A",
    "TanggalMulai": "2024-08-01",
    "TanggalAkhir": "2024-08-15"
  }
  ```

- Responses:
  - 204 No Content
  - 404 Not Found
  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Query tidak boleh kosong"
  }
  ```
  - 500 Internal Server Error
  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Terjadi kesalahan: [error_message]"
  }
  ```



## Mahasiswa Endpoints

**PostPresensiMahasiswa**

### Query Parameters

- Endpoint: POST /mahasiswa/register-mahasiswa
- Description
  Endpoint ini digunakan untuk mendaftarkan mahasiswa baru ke dalam sistem.
- Request Body:

  ```bash
  {
    "nama_lengkap": "Nama Lengkap Mahasiswa",
    "alamat": "Alamat Mahasiswa",
    "tempat_lahir": "Tempat Lahir Mahasiswa",
    "tanggal_lahir": "YYYY-MM-DD",
    "email": "email@example.com",
    "no_hp": 985908209380238,  // Nomor HP Mahasiswa sebagai angka
    "nama_pt_instansi": "Nama PT/Instansi",
    "jabatan": "Jabatan Mahasiswa",
    "nama_instagram": "Nama Instagram (opsional)",
    "nama_facebook": "Nama Facebook (opsional)",
    "sumber_informasi": "Sumber Informasi (opsional)"
  }
  ```
- Response
  - 201 OK
    ```bash
  {
    "status": true,
    "Status_Code": 201,
    "message": "Create Success",
    "data": {
      "ID": "ObjectID",
      "NamaLengkap": "Nama Lengkap Mahasiswa",
      "Alamat": "Alamat Mahasiswa",
      "TempatLahir": "Tempat Lahir Mahasiswa",
      "TanggalLahir": "YYYY-MM-DD",
      "Email": "email@example.com",
      "NoHp": 985908209380238,  // Nomor HP Mahasiswa sebagai angka
      "NamaPTInstansi": "Nama PT/Instansi",
      "Jabatan": "Jabatan Mahasiswa",
      "NamaInstagram": "Nama Instagram (opsional)",
      "NamaFacebook": "Nama Facebook (opsional)",
      "SumberInformasi": "Sumber Informasi (opsional)"
    }
  }
  ```

  - 400 Bad Request

  ```bash
  {
    "status": false,
    "Status_Code": 400,
    "message": "Invalid input"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error creating Mahasiswa"
  }
  ```

 **PostPresensiMahasiswa**

### Query Parameters

- Endpoint: POST /mahasiswa/presensi-mahasiswa
- Description
  Endpoint ini digunakan untuk mencatat kehadiran mahasiswa berdasarkan data yang ada di koleksi mahasiswa

- Request Body:

  ```bash
  {
    "nama_lengkap": "Nama Lengkap Mahasiswa",
    "alamat": "Alamat Mahasiswa",
    "tempat_lahir": "Tempat Lahir Mahasiswa",
    "tanggal_lahir": "YYYY-MM-DD",
    "email": "email@example.com",
    "no_hp": 985908209380238,  // Nomor HP Mahasiswa sebagai angka
    "nama_pt_instansi": "Nama PT/Instansi",
    "jabatan": "Jabatan Mahasiswa",
    "nama_instagram": "Nama Instagram (opsional)",
    "nama_facebook": "Nama Facebook (opsional)"
  }
  ```
- Response
  - 201 OK
    ```bash
  {
    "status": true,
    "Status_Code": 201,
    "message": "Presensi Success",
    "data": {
      "nama_lengkap": "Nama Lengkap Mahasiswa",
      "alamat": "Alamat Mahasiswa",
      "tanggal_lahir": "YYYY-MM-DD",
      "email": "email@example.com",
      "no_hp": 985908209380238,  // Nomor HP Mahasiswa sebagai angka
      "nama_pt_instansi": "Nama PT/Instansi",
      "jabatan": "Jabatan Mahasiswa",
      "nama_instagram": "Nama Instagram (opsional)",
      "nama_facebook": "Nama Facebook (opsional)",
      "tanggal_presensi": "YYYY-MM-DD",
      "status": "Hadir"
    }
  }

  ```

  - 400 Bad Request

  ```bash
  {
    "status": false,
    "Status_Code": 400,
    "message": "Invalid input"
  }
  ```
  - 404 Not Found

  ```bash
  {
    "status": false,
    "Status_Code": 404,
    "message": "Mahasiswa not found"
  }
  ```

  - 500 Internal Server Error

  ```bash
  {
    "status": false,
    "Status_Code": 500,
    "message": "Error creating Mahasiswa"
  }
  ```