sudo chown vscode /home/vscode/.local-kv

sudo chmod -R a+rwX /go/pkg

go install github.com/bokwoon95/wgo@latest
go install github.com/go-delve/delve/cmd/dlv@latest

cd api
go mod tidy

corepack enable