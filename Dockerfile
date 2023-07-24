# builder for svelte
FROM node:18-slim AS svelte-builder
WORKDIR /app
COPY ./client .
RUN npm install
RUN npm run build

# builder for go
FROM golang:1.20.6-alpine AS go-builder
WORKDIR /app
COPY ./server .
RUN apk add --no-cache ca-certificates
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -trimpath -o app

# production image
FROM scratch
COPY --from=go-builder /app/app /app
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=svelte-builder /app/public /app
ENTRYPOINT ["/app"]
