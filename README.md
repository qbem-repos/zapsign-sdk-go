# Zapsign SDK - Go

```go 
    var token = "klasjadslkjasdkladsjdaskljasdklasdjlajsas"     
    zcl, err := zapsign.NewClient(token)
    if err != nil {
        log.Fatal(err.Error())
    }

    var doc = new(requestbody.Docs)
    
    openid, err := zapsign.CreateDocs(doc)
    if err != nil {
        log.Fatal(err.Error())
    }

    println(openid) //42
```
