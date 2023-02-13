How to install
go mod init
go mod tidy


How to run
$ go run .

$ curl http://localhost:8080/albums

SEE.
https://go.dev/doc/tutorial/web-service-gin

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

curl http://192.168.1.53:8081/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "5","title": "testets","artist":"atttt","price": 99.99}'    