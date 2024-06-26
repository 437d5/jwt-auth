# JWT-Auth Microsevice

To start it you have to clone it with:
```git clone https://github.com/437d5/jwt-auth```

Then u have to create .env file in root directory with this structure:

![image](https://github.com/437d5/jwt-auth/assets/145232152/0da9e43e-00d0-4747-a9e6-afe5ef7ed7c8)


The next step is to start MongoDB and create database, collection and user also you need to specify
two new unique inedexes to escape user duplicates.
```
db.users.createIndex({"name": 1}, {"unique": true})
db.users.createIndex({"email": 1}, {"unique": true})
```
After it

```go mod tidy``` 

```go build cmd/main/main.go```

```./main```

You can check it out using evans gRPC client 

<https://github.com/ktr0731/evans>

![Screenshot from 2024-06-26 18-18-23](https://github.com/437d5/jwt-auth/assets/145232152/f8600d5f-0305-4dd6-90c5-186764fc8f05)

