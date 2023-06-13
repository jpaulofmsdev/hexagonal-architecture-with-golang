### Hexagonal Architeture with GoLang

Just a simple project to introduce ports and adapters.

Need a lot of refactoring at some places because it was not the focus.

#### Running the project
```bash
docker-compose up -d

docker exec -it appproduct bash

# for cli usage
go run main.go cli --help

# for server usage
go run main.go http
```

### Notes while developing

#### Running mockgen

```bash
mockgen -destination=application/mocks/application.go -source=application/product.go
```

#### Running cobra
```bash
cobra-cli add cli
```

```bash
cobra-cli add http
```