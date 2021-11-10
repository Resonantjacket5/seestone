

  seeing stones

  cli for the one api using urfave/cli



https://the-one-api.dev/





```
  go run main.go
```

build
```
  go install
  $GOPATH/bin/seestone

  # or for windows
  $GOPATH\\bin\\seestone
```

Run
```
  seestone books list --api-token <apitoken>
```


* books
  * id
  * id/chapter
* chapter
  * id
* movie
  * id
  * id/quote
* character
  * id
  * id/quote
* quote
  * id


For bash

    curl --insecure -H "Authorization: Bearer APIKEY"  https://the-one-api.dev/v2/movie

For powershell

    # curl is alias for Invoke-WebRequest
     curl -H @{"Authorization"="Bearer APIKEY"} -uri https://the-one-api.dev/v2/movie