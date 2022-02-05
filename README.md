# Restaurant APP

A rest-api that works with golang

## Usage

Fist of all, clone the repo with the command below. You must have golang installed on your computer to run the project.

```shell
git clone https://github.com/bartukocakara/golang-mongodb
```

```shell
swag init
```

+ ### Locally


```shell
go run main.go
```

Now, you can send requests to the application either via postman, swagger or browser (GET requests only) according to
the following endpoint table.


Now go to `localhost:8000` in your browser. You will see the swagger page. Here you can send your requests to api.

or you can test the api on the URL below.

# Heroku link (example '/tables')

https://restaurant-go-app.herokuapp.com/tables

## Endpoint Table

| Endpoints  | Descriptions |  Methods | 
| :------:|  :-----------:| :-----------:|
| /user   |  | GET |
| /user/:user_id   |  | GET |
| /user/:user_id   |  | PATCH |
| /users/signup   |  | POST |
| /users/login   |  | POST |
| /users/delete   |  | DELETE |
| /foods   |  | GET |
| /foods   |  | POST |
| /foods/:food_id   |  | GET |
| /foods/:food_id   |  | PATCH |
| /foods/:food_id   |  | DELETE |
| /orders  |  | GET | 
| /orders  |  | POST | 
| /orders/:order_id  |  | GET | 
| /orders/:order_id  |  | PATCH | 
| /order/:order_id  |  | DELETE | 
| /order-items  |  | GET | 
| /order-items  |  | POST | 
| /order-items/:order_item_id  |  | GET | 
| /order-items/:order_item_id  |  | PATCH | 
| /tables  |  | GET | 
| /tables  |  | POST | 
| /tables/:table_id  |  | PATCH | 
| /tables/:table_id  |  | DELETE | 


