# crossx Noé & Kilian

## Project install

### repository root

- `docker-compose -p crossx up -d`

### ./api

- `go get`
- `go install github.com/swaggo/swag/cmd/swag@latest`
- `go get -u github.com/swaggo/swag`
- `swag init`
- `~/go/bin/swag init & go run .`

### ./frontend

- `npm install`
- `npx expo start`
