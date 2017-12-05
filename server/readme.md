## Run server

```
go run server/main.go
```

## Use grpcc test

```
grpcc -d ./proto -p example/helloworld.proto -a localhost:50051 -i

Greeter@localhost:50051> client.sayHello({name: "Bob"}, pr)
```

**Notice**: You can also run proxy.go and test the service using http.