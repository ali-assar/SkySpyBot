# sudo docker build -t skyspy:latest .
# sudo docker run -p 80:8080 -e TELEGRAM_BOT_TOKEN="7107379255:AAFVGFePM-735Z8gPqocT52IaekPkp9pJ_M" -e OWM_API_KEY="e738e445e75e7ce26e9f98d7f41299ed" -e REDIS_ADDRESS="localhost:6379" skyspy:latest
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
EXPOSE 8080
ENTRYPOINT ["/app/skyspy"]