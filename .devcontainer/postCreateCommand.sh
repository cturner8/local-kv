sudo chmod -R a+rwX /go/pkg

go install github.com/bokwoon95/wgo@latest
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

cd api
go mod tidy
