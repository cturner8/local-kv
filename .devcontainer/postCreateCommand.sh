sudo chmod -R a+rwX /go/pkg

go install github.com/bokwoon95/wgo@latest

cd api
go mod tidy
