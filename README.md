# Install `wasmtime`

Following https://wasmtime.dev/ , 

```sh
curl https://wasmtime.dev/install.sh -sSf | bash
```

# Build main.go for Wasm runtime

```
GOOS=wasip1 GOARCH=wasm go build -o main
```

# execute it

```
wasmtime main
```

```
Hello, wasm
```
