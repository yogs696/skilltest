<pre style="font-size: 1.4vw;">
<p align="center">
                _                         
 _             | |    _              _    
| |_  ____ ____| | _ | |_  ____  ___| |_  
|  _)/ _  ) ___) || \|  _)/ _  )/___)  _) 
| |_( (/ ( (___| | | | |_( (/ /|___ | |__ 
 \___)____)____)_| |_|\___)____|___/ \___)
                                          
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
<b>Go - Techtest</b> is a test preparation </b>
</p>

# Techtest API Guide

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

| NO  | Use                 | Endpoint                   | Example                                      | Action |
| --- | ------------------- | -------------------------- | -------------------------------------------- | ------ |
| 1   | register            | api/v1/user/register       | http://localhost:4040/v1/user/register       | POST   |
| 2   | Login               | api/v1/user/login          | http://localhost:4040/v1/user/login          | POST   |
| 3   | list Product        | api/v1/product/list        | http://localhost:4040/v1/product/list        | GET    |
| 3   | Create Product      | api/v1/product/create      | http://localhost:4040/v1/product/create      | POST   |
| 4   | Update Product      | api/v1/product/update/{id} | http://localhost:4040/v1/product/update/{id} | PUT    |
| 5   | Delete Product      | api/v1/product/delete/{id} | http://localhost:4040/v1/product/delete/{id} | DELETE |
| 6   | create item to cart | api/v1/cart/create         | http://localhost:4040/v1/cart/create         | POST   |
| 7   | create order        | api/v1/order/create        | http://localhost:4040/v1/order/create        | POST   |

---

## 📖 Compatible JSON Payload Techtest API

This is the JSON payload that's sended to Techtest API

### 💸 List Product Datatable Request

```js
curl --location --request GET 'localhost:4040/v1/product/list' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer your token' \
--data '{
    "draw": 0,
    "search": "test",
    "length": 3,
    "offset": 0
}'
```

### 💸 List Product Datatable Response

```js
{
    "draw": 1,
    "recordsTotal": 1,
    "filteredTotal": 1,
    "data": [
        {
            "id": 1,
            "name": "test Update 123",
            "description": "test",
            "price": 0
        }
    ]
}
```

### 💸 Create Product Request

```js
curl --location 'localhost:4040/v1/product/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer your token' \
--data '{
	"name": "test PROD",
    "description": "1",
    "price": 1000
}'
```

### 💸 Create Product Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Product [test] successfully created",
    "error": null
}
```

### 💸 Update Product Request

```js
curl --location --request PUT 'localhost:4040/v1/product/update/1' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer your token' \
--data '{
	"name": "test PROD",
    "description": "1",
    "price": 1000
}'
```

### 💸 Update Product Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Data product successfully updated",
    "error": null
}
```

### 💸 Delete Product Request

```js
curl --location --request DELETE 'localhost:4040/v1/product/delete/10' \
--header 'Authorization: Bearer your token' \
```

### 💸 Delete Product Response

```js
{
    "success": true,
    "code": 2400,
    "data": "product ID [10] successfully deleted",
    "error": null
}
```

### 💸 Add to Cart Request

```js
curl --location 'localhost:4040/v1/cart/create' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer your token' \
--data '{
    "product_id": 1
}'
```

### 💸 Add to Cart Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Item Added to Cart!",
    "error": null
}
```

### 💸 Create Order Request

```js
curl --location --request POST 'localhost:4040/v1/order/create' \
--header 'Authorization: Bearer your token' \
--data ''
```

### 💸 Create Order Response

```js
{
    "success": true,
    "code": 2400,
    "data": "Thank you, Your order successfully!",
    "error": null
}
```

### 📖 Excersise 2

```js
## Fetch a user by username.
    - we need do index with composite:username_created_at_key

## Fetch users who signed up after a certain date (created_at > "2023-01-01").
    - we need do only individual index on coloum created_at

## Fetch a user by email.
    - we need do index with composite:email_created_at_key

```

### 📖 Excersise 4

```js
    - sory for my answer is not giving the function only describing the logic, cause the time is limit, sorry for the inconvenience,
    if depisit and withdarw function running using concurrent we need handle race condition using sync.Mutex

```

### 📖 Excersise 4

```js
## Write an optimized SQL query to find the top 5 customers who spent the most money in the past month.
    SELECT
        o.customer_id,
        SUM(o.amount) AS total_spent
    FROM
        orders o
    WHERE
        o.order_date >= NOW() - INTERVAL 1 MONTH  -- filter orders from the last month
    GROUP BY
        o.customer_id
    ORDER BY
        total_spent DESC
    LIMIT 5;

## How would you improve the performance of this query in a production environment?
    - we need do index on order_date

```

### 📖 Excersise 5

```js
## What steps would you take to decompose the service into smaller, more manageable services?
    - we do decompose to be user service and management service (for all file)

## How would you ensure that the new system is backward compatible with the old one during the transition?
    - before going to production we need make sure running well on staging enviroment, if all good, we can go to production

```
