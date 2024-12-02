## Intro

This is a POC of TigerBeetle [currency exchange](https://docs.tigerbeetle.com/coding/recipes/currency-exchange/). 
It is by no means complete, but is meant to serve as a solid starting point for a more robust product.

## Requirements

- Docker
- go 1.21
- node 


### Install Dependencies

```
make setup
```

### Start db

```
docker compose up db
```

### Start app

```
make dev
```

Now open your browser to `localhost:4000/` and start generating commissions for the service.