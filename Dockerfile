# sudo docker build -t skyspy:latest .
# sudo docker run -p 5000:5000 skyspy:latest
FROM golang:1.22.1-alpine3.19 as build
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN export GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-w -s" -o ./skyspy ./main.go

FROM redis:7.2.5-alpine
WORKDIR /app
COPY --from=build /app/skyspy /app
COPY .env /app/
COPY start.sh /app/
RUN chmod +x /app/start.sh
EXPOSE 5000
ENTRYPOINT ["/app/start.sh"]