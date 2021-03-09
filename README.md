# challenge_quick_wallet_api

## Setup and run

1. Clone the repo
   ```git clone https://github.com/niki4/challenge_quick_wallet_api```
2. Start the containers (make sure you have Docker/docker-compose installed, as well as VPN turned off):

```
docker-compose up --build
```

By default, app will run in debug mode ON (`DEBUG: "true"` in `docker-compose.yml`) so the app will create some test
wallets for convenience of the testing.

3. Once containers started, you can try following requests (assuming you have `curl` installed, otherwise use your
   favourite client):

* **Retrieve the balance** of a given wallet ID (1)

```curl
curl -i -X GET -H "Accept: application/json" http://localhost:8080/api/v1/wallets/1/balance
```

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{"balance":"100","id":1} 

```

* **Credit wallet** with ID 1 with amount 0.15 so that user balance will be increased on that sum.

```curl
curl -i --header "Content-Type: application/json" --request POST --data '{"id":1,"balance":0.15}' http://localhost:8080/api/v1/wallets/1/credit 
```

The route returns updated balance for given wallet.

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{"balance":"100.15","id":1} 
```

* **Debit wallet** with ID 1 with amount 0.55 so that user balance will be decreased on that sum.

```curl
curl -i --header "Content-Type: application/json" --request POST --data '{"id":1,"balance":0.55}' http://localhost:8080/api/v1/wallets/1/debit 
```

The route returns updated balance for given wallet.

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8

{"balance":"99.6","id":1} 
```

It will return error in case requested debit amount exceed wallet balance.

### Context

You are responsible for managing the wallets of the players of an online casino, and you need to provide an API for
getting and updating their account balances.

### Endpoints

* _balance_: retrieves the balance of a given wallet id
  ```GET /api/v1/wallets/{wallet_id}/balance```
* _credit_: credits money on a given wallet id
  ```POST / api/v1/wallets/{wallet_id}/credit```
* _debit_: debits money from a given wallet id
  ```POST / api/v1/wallets/{wallet_id}/debit```
