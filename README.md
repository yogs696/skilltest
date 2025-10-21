<pre style="font-size: 1.4vw;">
<p align="center">
      _     _ _ _                        
     | |   (_) | |   _              _    
  ___| |  _ _| | |  | |_  ____  ___| |_  
 /___) | / ) | | |  |  _)/ _  )/___)  _) 
|___ | |< (| | | |  | |_( (/ /|___ | |__ 
(___/|_| \_)_|_|_|   \___)____|___/ \___)
</p>
</pre>
<p align="center">
<a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
</a>
<a href="/LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-green.svg">
</a>
</p>
<p align="center">
<b>Go - skilltest</b> is a test preparation </b>
</p>

# skilltest API Guide

## 🔀 How to Run :

```js
- clone the project using, git clone https://github.com/yogs696/skilltest.git
- open the directory project and run go mod tidy then go mod vendor
- create file .config.yaml like the example and set the configuration
- run cli for db migration using : go run main.go -db-migrate
- this project can be run using air or using command go run main.go --run
- if you want to run this project using air, make sure air already install  on your laptop
```

## 🔀 Compatible Route Endpoint

| NO  | Use             | Endpoint                    | Example                                       | Action |
| --- | --------------- | --------------------------- | --------------------------------------------- | ------ |
| 1   | register        | api/v1/auth/register        | http://localhost:4040/v1/auth/register        | POST   |
| 2   | Login           | api/v1/auth/login           | http://localhost:4040/v1/auth/login           | POST   |
| 3   | list Schedule   | api/v1/schedule/list        | http://localhost:4040/v1/schedule/list        | GET    |
| 3   | Create Schedule | api/v1/schedule/create      | http://localhost:4040/v1/schedule/create      | POST   |
| 4   | Update Schedule | api/v1/schedule/update/{id} | http://localhost:4040/v1/schedule/update/{id} | PUT    |
| 5   | Delete Schedule | api/v1/schedule/delete/{id} | http://localhost:4040/v1/schedule/delete/{id} | DELETE |

---

## 📖 Compatible JSON Payload skilltest API

This is the JSON payload that's sended to skilltest API

### 💸 Register Curl

```js
curl --location 'localhost:4040/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
	"username": "testing",
    "email": "test@gmail.com",
    "password": "admin123"
}'
```

### 💸 Register Response

```js
{
    "success": true,
    "code": 2400,
    "data": {
        "id": 1,
        "name": "testing",
        "email": "test@gmail.com",
        "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwic2VjIjoiZDQxZDhjZDk4ZjAwYjIwNGU5ODAwOTk4ZWNmODQyN2UiLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3RpbmcifQ.vJNGpo8x2H-Qe16Xk6AFc9tron5vzN2RQfhIjM4Yps_LxlAzsnIbr-uCs0m0G67LgQcHCXcdM9p4tLlVs9V9XqyddC9KnjIHh0VnzCwed_gapTP1dFGk_1dy0XUajnUfpuEE1QsBy19iidMD44tKewUjAdAQ1n92ZCifQq5wVc1saF0ExWCWQbWC-K0PyiiLMPLtDuH19r0xChqvt-_EDGFCHzUeKQYUW8w5uShHFnksRiw3Lx_NO-LMN5_K6VuobvQB_-11HzsS26AYWu4KpkbhaWdJ3qhr-F8RXV3oU2xWWKwKm7ifYCJIhjqVi51mQh2M8T5vgfqi1bAUgaxLKg"
    },
    "error": null
}

```

### 💸 Login Curl

```js
curl --location 'localhost:4040/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@gmail.com",
    "password": "admin123"
}'
```

### 💸 Login Response

```js
{
    "success": true,
    "code": 2400,
    "data": {
        "id": 1,
        "name": "testing",
        "email": "test@gmail.com",
        "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwic2VjIjoiZDQxZDhjZDk4ZjAwYjIwNGU5ODAwOTk4ZWNmODQyN2UiLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3RAZ21haWwuY29tIn0.CRUK403DMvcqEgNp2_b2oAT4GjnHZm-u1obQfZjMnIW5UnjLmqJzuZnC1tmiRTIxCPCk4mbXjZcYTG11IjDIWhH7s8ND1erp--w1YQ-HYSZhXJwKd6EOBFS1Qg-Ti-UvZPXB8Ya96zduwOX6VlefsuPoqu5RoljBhns9vahUMlofoCEoYej671vlYelUOm2eO1ANVcDnPdnAivAtiuphV11rqtfaFIGh7snmmAaT-IHqTQDvm8Z7b_3jpJXE-Xt8hAqhJUlep05Tvqqrd2fYfJy_P_EEMUVd1C94i-IU-vbYfuvPkdkuAanB9lO-ZqnQXk0okvtR5TCe-9l8KswJxw"
    },
    "error": null
}
```

### 💸 Curl Schedule List

```js
curl --location --request GET 'localhost:4040/v1/schedule/list' \
--header 'Authorization: Bearer your token' \
--header 'Content-Type: application/json' \
--data '{
    "draw": 0,
    "search": "2", //for cinema_id or movie_id
    "length": 10,
    "offset": 0
}'
```

### 💸 List Schedule Datatable Response

```js
{
    "draw": 1,
    "recordsTotal": 2,
    "filteredTotal": 2,
    "data": [
        {
            "id": 1,
            "cinema_id": 12,
            "movie_id": 2,
            "show_date": "2025-10-21T07:00:00+07:00",
            "start_time": "19:00",
            "end_time": "21:00",
            "created_at": "2025-10-21T10:28:21.63434+07:00",
            "updated_at": "2025-10-21T10:41:30.918438+07:00"
        },
        {
            "id": 2,
            "cinema_id": 20,
            "movie_id": 2,
            "show_date": "2025-10-21T07:00:00+07:00",
            "start_time": "19:00",
            "end_time": "21:00",
            "created_at": "2025-10-21T10:39:26.316268+07:00",
            "updated_at": "2025-10-21T10:41:45.26702+07:00"
        }
    ]
}
```

### 💸 Curl Create Schedule

```js
curl --location 'localhost:4040/v1/schedule/create' \
--header 'Authorization: Bearer your token' \
--header 'Content-Type: application/json' \
--data '{
    "cinema_id": 12,
    "movie_id": 2,
    "show_date": "2025-10-21 00:00:00",
    "start_time": "19:00",
    "end_time": "21:00"
}'
```

### 💸 Create Schedule Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Data Schedule successfully created",
    "error": null
}
```

### 💸 Curl Update Schedule

```js
curl --location --request PUT 'localhost:4040/v1/schedule/update/2' \
--header 'Authorization: Bearer your token' \
--header 'Content-Type: application/json' \
--data '{
    "cinema_id": 20,
    "movie_id": 2,
    "show_date": "2025-10-21 00:00:00",
    "start_time": "19:00",
    "end_time": "21:00"
}'
```

### 💸 Update Schedule Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Data Schedule successfully updated",
    "error": null
}
```

### 💸 Curl Delete Schedule

```js
curl --location --request DELETE 'localhost:4040/v1/schedule/delete/2' \
--header 'Authorization: Bearer your token' \
```

### 💸 Delete Schedule Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Data Schedule successfully deleted",
    "error": null
}
```

### 📖 Excersise Point A.2

```js
    1. Proses Pemilihan Tempat Duduk
        Tahap 1 – Load Denah Tempat Duduk
            1. Client membuka halaman pemilihan kursi.
            2. Sistem memuat denah (seats + seat_statuses dari schedule_id).
            3. Untuk performa:
                - Gunakan cache Redis dengan key seatmap:schedule_id.
                - Cache disinkronkan setiap update seat status.

        Tahap 2 – User Memilih Kursi (Hold)
            1 Saat user memilih kursi dan seatnya akan dilock menggunakan mutex lock untuk mencegah race condition dan sistem akan:
                - Mengecek seat_statuses.status = AVAILABLE.
                - Membuat hold sementara (status → HELD, isi hold_user_id, hold_expires_at).
            2. Jika berhasil → kursi dikunci 10 menit.
            3. Worker background memantau kursi yang hold_expires_at < NOW() untuk di-release.


        Tahap 3 – Pembayaran & Issued Ticket
            1. Setelah user klik Bayar, sistem membuat record di payments dengan status = PENDING.
            2. Setelah payment gateway callback sukses:
                - Update payments.status = SUCCESS.
                - Ubah seat_statuses.status = 'SOLD'.
                - mutex lock auto unlocked

    2. Refund / Pembatalan oleh Bioskop
        Kasus A — Refund karena Pembatalan Film (schedules.status = CANCELED)
            1. Admin mengubah schedules.status → CANCELED.
            2. Worker mendeteksi perubahan ini dan:
                - Menemukan semua tickets dengan status = ISSUED.
                - Membuat entry di refunds (status = PENDING).
                - Setelah dana dikembalikan → refunds.status = SUCCESS, tickets.status = REFUNDED, payments.status = REFUNDED.

        Kasus B — Refund Individual (User Request)
            1. User meminta refund → buat record refunds (status=PENDING).
            2. Setelah divalidasi admin:
                - Update refunds.status = SUCCESS, tickets.status = REFUNDED, payments.status = REFUNDED.
                - seat_statuses.status → AVAILABLE → kursi dapat dibeli kembali.

    3. Restock Tiket / Kursi
        Tiket bisa restock dalam dua kondisi:
            - Hold expired → kursi otomatis AVAILABLE.
            - Refund berhasil → kursi AVAILABLE.
                Worker task (tiap 1 menit):
```
