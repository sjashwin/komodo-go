![Komodo_Logo](logo.png?raw=true)

# Komodo RPC API Library Go Lang

RPC API-Library for GoLang. This library allows developers to communicate with the komodo client APIs.

## Description

Komodo-RPC library helps you integrate your Python Apps with Komodo asset-chains without having to setup/implement required RPC functions. Install this Python library and call Komodo-API RPCs as easily as calling a local function. Komodo-RPC library acts as a wrapper between your Python app and the Komodo-daemon running on a server.

## Getting Started

You will first need to get the library from the github repository by executing the following the command.

```
go get -u github.com/sjashwin/komodo-go
```

Make sure you are now running the komodo server, which is `komodod`.

Once the komodo server has successfully started you can then execute this main function. Please make sure you change the authentication properties.
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

Please feel free to report any bugs.
