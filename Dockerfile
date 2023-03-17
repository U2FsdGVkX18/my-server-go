FROM golang:1.20
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download && go build -o main .
CMD ["./main"]


#FROM golang:1.20
 #WORKDIR /usr/src/app
 ## pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
 #COPY go.mod go.sum ./
 #RUN go env -w GOPROXY=https://goproxy.cn
 #RUN go mod download && go mod verify
 #COPY . .
 #RUN go build -v -o /usr/local/bin/app main.go
 #CMD ["app","&>","server.log"]