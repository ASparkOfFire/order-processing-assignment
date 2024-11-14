# Order Processing Service - Assignment

# Steps To Run
1. Clone the repository
2. Run Postgres Instance:`docker-compose up -d`
3. Run the application using `go run main.go`

### API Documentation:
1. `GET /api/orders/<7order_id>` - Returns a list of all orders
2. `GET /api/customers/<customer_id>` - Returns a customer
3. `GET /api/customers` - Returns a list of all customers

4. `POST /api/orders` - Creates an order
  
Request Schema
```json
  {
  "customer_id": 1,
  "products": [
    {"id": 2},
    {"id": 2},
    {"id": 3}
  ]
}  
```
