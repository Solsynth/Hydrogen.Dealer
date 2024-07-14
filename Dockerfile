# Building Backend
FROM golang:alpine as dealer-server

RUN apk add nodejs npm

WORKDIR /source
COPY . .

WORKDIR /source/web
RUN npm install
RUN npm run build

WORKDIR /source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs -o /dist ./pkg/main.go

# Runtime
FROM golang:alpine

COPY --from=dealer-server /dist /dealer/server
COPY --from=dealer-server /source/web/dist /dealer/web

EXPOSE 8444

CMD ["/dealer/server"]
