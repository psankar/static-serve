# static-serve

```
$ go run cmd/main.go &
[1] 81301

$ curl http://localhost:8080/public/static-file.html
<html>
    <body>
        <h1>Static File</h1>
    </body>
</html>

$ curl http://localhost:8080/embedded/static-file.html
404 page not found
```
