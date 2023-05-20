```
docker compose up -d
cd terraform
terraform apply -auto-approve
cd ../app
go run ./cmd/main.go
```

Stablish websocket connection with ws://localhost:4510?charger=ch12345
