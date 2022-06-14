## **Fully functional ECOMMERCE API USING GIN FRAMEWORK AND MONGODB**


To Run project

```bash
# Start the mongodb container for local development
docker-compose up -d
go run main.go
```

## API FUNCTIONALIITIES

- Signup 🔒
- Login 🔒
- Product listing General View 👀
- Adding the products to DB
- Sorting the products from DB using regex 👀
- Adding the products to cart 🛒
- Removing the Product from cart🛒
- Viewing the items in cart with total price🛒💰
- Adding the Home Address 🏠
- Adding the Work Address 🏢
- Editing the Address ✂️
- Deleting the Adress 🗑️
- Checkout the Items from Cart
- Buy Now products💰

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.17 or newer and the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) installed.

```sh
$ git clone https://github.com/heroku/go-getting-started.git
$ cd go-getting-started
$ go build -o bin/golang-gin-poc -v . # or `go build -o bin/go-getting-started.exe -v .` 