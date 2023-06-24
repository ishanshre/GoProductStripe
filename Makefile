buildAPI:
	go build -o api cmd/api/* && ./api
buildWEB:
	go build -o web cmd/web/*.go && ./web
