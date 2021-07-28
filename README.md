# Interview Hash


## Tecnologies
- Golang
- Docker
- Sonarkube
- gRPC

## Design Pattern, Architecture
- Factory
- Dependency Injection
- Repository Pattern
- DDD
- Clean Architecture
- Microservices Architecture

## Applications
- checkout-service
> responsible for create checkout for ecommerce
- discount service
> responsible for getting product and return a discount

## Project
### Start Project
#### 1 - Create .env file
- Create .env file using .env.sample and set the variables there

#### 2 - Step
**If want to run local**
- API
```bash
$ go run ./my-ecommerce/cmd/main.go
```
> Remember, if you want to validate if discount service is running you must start the service run `docker-compose up -d discount-service`

**If want to run in docker**
```bash
$ make build
$ make start
```

**If you want to stop the application in docker**
```bash
$ make stop
```

### Routes
> Host
`http://localhost:4513/`

> Health Check
- GET `/health`
- Params
- Response
```json
"OK"
```

> POST checkout
- POST `/checkout`
- Params
    - products: array json
        - id : int
        - quantity: int
- Example Request
```json
{
  "products": [
    {
      "id": 1,
      "quantity": 1
    }
  ]
}
```
- Response
```json
{
    "total_amount": 20000, 
    "total_amount_with_discount": 19500,
    "total_discount": 500,
    "products": [
        {
            "id": 1,
            "quantity": 2,
            "unit_amount": 10000, 
            "total_amount": 20000,
            "discount": 500,
            "is_gift": false
        },
        {
            "id": 3,
            "quantity": 1,
            "unit_amount": 0,
            "total_amount": 0, 
            "discount": 0,
            "is_gift": true
        }
    ]
}
```

## Run test
- Local Test
> Run all tests (unit, e2e)
```bash
$ make tests
```

- Local Test in Docker
> Run all tests (unit, e2e)
```bash
$ make tests-docker
```

## Logs
**If you are using docker**
- Show full logs
```bash
make logs-checkout-service
```

- Show tail logs (last 100 lines)
```bash
make logs-tail
```