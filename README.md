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

**If want to run in docker**
```bash
$ make build
$ make start
```

### Routes
> Host
- in local:

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
- Example
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

## Run test
- Local Test
> Run all tests (unit, e2e)
```bash
$ make tests
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