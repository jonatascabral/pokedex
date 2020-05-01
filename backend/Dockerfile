FROM golang:1.14-alpine AS builder

RUN apk --no-cache add tzdata

ARG APPPATH=$GOPATH/src/github.com/jonatascabral/pokedex/backend

COPY . $APPPATH

WORKDIR $APPPATH

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -o /app .

FROM alpine

ARG APP_PORT=8000
ARG APPPATH=/go/src/github.com/jonatascabral/pokedex/backend
ARG TZ=America/Sao_Paulo

WORKDIR /api/

COPY --from=builder /app ./app
COPY --from=builder /usr/share/zoneinfo/$TZ /etc/localtime
COPY --from=builder $APPPATH/.env ./.env
COPY --from=builder $APPPATH/swagger.json ./swagger.json
COPY --from=builder $APPPATH/migrations ./migrations

RUN echo $TZ > /etc/timezone && \
    chmod +x /api/app

CMD ["/api/app"]

EXPOSE $APPPORT
