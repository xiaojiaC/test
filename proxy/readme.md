## Run proxy

```
go run proxy/main.go
```

## Send request test

- POST http://localhost:8080/v1/example/hello

```json
{
	"name": "Bob"
}
```

or

- Click http://petstore.swagger.io/?url=http://localhost:8080/swagger/helloworld.swagger.json

Find service & Execute
