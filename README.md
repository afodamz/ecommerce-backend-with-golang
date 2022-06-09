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


- **SIGNUP FUNCTION API CALL (POST REQUEST)**

http://localhost:8000/users/signup

```json
{
  "first_name": "oluwadamilola",
  "last_name": "afolabi",
  "email": "afodamz@gmail.com",
  "password": "testpass",
  "phone": "+1558426655"
}
```

Response :"Successfully Signed Up!!"

- **LOGIN FUNCTION API CALL (POST REQUEST)**

  http://localhost:8000/users/login


```json
{
  "email": "afodamz@gmail.com",
  "password": "testpass"
}
```

response will be like this

```json
{
  "_id": "***********************",
  "first_name": "oluwadamilola",
  "last_name": "afolabi",
  "email": "josephhermis@protonomail.com",
  "phone": "+1558921455",
  "token": "eyJc0Bwcm90b25vbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiam9zZXBoIiwiTGFzdF9OYW1lIjoiaGVybWlzIiwiVWlkIjoiNjE2MTRmNTM5ZjI5YmU5NDJiZDlkZjhlIiwiZXhwIjoxNjMzODUzNjUxfQ.NbcpVtPLJJqRF44OLwoanynoejsjdJb5_v2qB41SmB8",
  "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnLCJVaWQiOiIiLCJleHAiOjE2MzQzNzIwNTF9.ocpU8-0gCJsejmCeeEiL8DXhFcZsW7Z3OCN34HgIf2c",
  "created_at": "2021-10-09T08:14:11Z",
  "updtaed_at": "2021-10-09T08:14:11Z",
  "user_id": "61614f539f29be942bd9df8e",
  "usercart": [],
  "address": [],
  "orders": []
}
```

Login Function call create an outlayer for our collection

- **Admin add Product Function POST REQUEST**

  **note this function is not seperated from normal user fixed soon for admin**

  http://localhost:8000/admin/addproduct

```json
{
  "product_name": "laptop",
  "price": 300,
  "rating": 10,
  "image": "1.jpg"
}
```

Response : "Successfully added our Product Admin!!"

- **View all the Products in db GET REQUEST**

  pagination added soon in next release

  http://localhost:8000/users/productview

Response

```json
[
  {
    "Product_ID": "6153ff8edef2c3c0a02ae39a",
    "product_name": "notepad",
    "price": 50,
    "rating": 10,
    "image": "penc.jpg"
  },
  {
    "Product_ID": "616152679f29be942bd9df8f",
    "product_name": "laptop",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "top",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "table",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "apple",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
```

- **Search Product by regex function (GET REQUEST)**

defines the word search sorting
http://localhost:8000/users/search?name=le

response:

```json
[
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "table",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "apple",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
```

The corresponding Query to mongodb is **ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": queryParam}})**

- **Adding the Products to the Cart (GET REQUEST)**

  http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

  Corresponding mongodb query

```go
filter := bson.D{primitive.E{Key: "_id", Value: id}}
update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productcart}}}}}}
_, err = UserCollection.UpdateOne(ctx, filter, update)
```

- **Removing Item From the Cart (GET REQUEST)**

  http://localhost:8000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx

  Corresponding mongodb query

```go
filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}

update := bson.M{"$pull": bson.M{"usercart": bson.M{"_id": removed_id}}}
_, err = UserCollection.UpdateMany(ctx, filter, update)
```

- **Listing the item in the users cart (GET REQUEST) and total price**

  http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx

  Corresponding Mongodb Query (WE are using the aggrgate operation to find sum)

       filter_match := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}}
      unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
      grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
      pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind, grouping})

- **Addding the Address (POST REQUEST)**

  http://localhost:8000/addadress?id=user_id**\*\***\***\*\***

  The Address array is limited to two values home and work address more than two address is not acceptable

```json
{
  "house_name": "jupyterlab",
  "street_name": "notebook",
  "city_name": "mars",
  "pin_code": "685607"
}
```

- **Editing the Home Address(PUT REQUEST)**

  http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

- **Editing the Work Address(PUT REQUEST)**

  http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

- **Delete Addresses(GET REQUEST)**

  http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

  delete both addresses

- **Cart Checkout Function and placing the order(GET REQUEST)**

  After placing the order the items have to be deleted from cart functonality added

  http://localhost:8000?id=xxuser_idxxx

- **Instantly Buying the Products(GET REQUEST)**
  http://localhost:8000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx

