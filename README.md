# GoLang Komodo API 

Still Under Construction

## Getting Started
```go
func main(){
    authentication := &komodo.Auth{
        Password: "",
        User: "",
        Port: 9009,
    }
    client:=komodo.Client{Auth: authentication}
    parameter := komodo.Params{
        "addresses": []string{"RTTg3izdeVnqkTTxjzsPFrdUQexgqCy1qb"}
        "start": 1,
        "end": 200,
        "chaininfo": true,
    }
    client.GenerateAddressDeltas(parameter)
}
```