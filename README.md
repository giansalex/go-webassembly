# Go WebAssembly Calculator

### Generate wasm
```
GOOS=js GOARCH=wasm go build -o main.wasm
```

### Run local http server

```
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

Go to http://localhost:8080
