# builder for svelte
FROM node:22-alpine AS svelte-builder
WORKDIR /client
COPY client .
RUN npm install --force
RUN npm run build

# builder for go
FROM golang:1.20-alpine AS go-builder
WORKDIR /server
COPY server .
COPY --from=svelte-builder /client/build/ build/
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o app
RUN apk update && apk add --no-cache ca-certificates && apk add --no-cache upx
RUN upx --best app

# production image
FROM scratch
EXPOSE 4000
COPY --from=go-builder /server/app /
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app"]
