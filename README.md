## **Fully functional ECOMMERCE API USING GIN FRAMEWORK AND MONGODB**


To Run project

```bash
# Start the mongodb container for local development
docker-compose up -d
go run main.go
```

## API FUNCTIONALIITIES

- Signup π
- Login π
- Product listing General View π
- Adding the products to DB
- Sorting the products from DB using regex π
- Adding the products to cart π
- Removing the Product from cartπ
- Viewing the items in cart with total priceππ°
- Adding the Home Address π 
- Adding the Work Address π’
- Editing the Address βοΈ
- Deleting the Adress ποΈ
- Checkout the Items from Cart
- Buy Now productsπ°

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.17 or newer and the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) installed.

```sh
$ git clone https://github.com/heroku/go-getting-started.git
$ cd go-getting-started
$ go build -o bin/golang-gin-poc -v . # or `go build -o bin/go-getting-started.exe -v .` 