# Go-LibDucoHasher

Go version of libducohasher alternatives

> 1.6x faster based on my test with original Python DUCOS1 algorithms

## Usage

- Create shared lib

```bash
go build -o main.so -buildmode=c-shared main.go
```

- Import from python file
