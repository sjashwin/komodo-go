![Komodo_Logo](logo.png?raw=true)

# Komodo RPC API Library Go Lang

RPC API-Library for GoLang developers

## Description

Still Under construction

## Getting Started

```go
func main(){
    authentication := &komodo.Auth{
        Host: "127.0.0.1",
        User: "myusername",
        Password: "mypassword",
        Port: 7771,
    }
    client := komodo.Client{
        Auth: authentication
    }
    if j, err := client.GetWalletInfo(); err != nil{
        panic(err)
    } else {
        fmt.Println(j)
    }
}
```
