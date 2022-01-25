# Restaurant APP

A rest-api that works with golang

## Usage

Fist of all, clone the repo with the command below. You must have golang installed on your computer to run the project.

```shell
git clone https://github.com/bartukocakara/golang-mongodb
```

+ ### Locally


```shell
go run main.go
```

Now, you can send requests to the application either via postman, swagger or browser (GET requests only) according to
the following endpoint table.


Now go to `localhost:8000` in your browser. You will see the swagger page. Here you can send your requests to api.

or you can test the api on the URL below.

## Endpoint Table

| Endpoints  | Descriptions |  Methods | 
| :------:|  :-----------:| :-----------:|
| /user/:user_id   |  | GET |
| /food   |  | POST |
| /food   |  | POST |
| /order  |  | GET | 
| /order  |  | POST | 
| /order/:order_id  |  | GET | 
| /flush  |  | DELETE |


