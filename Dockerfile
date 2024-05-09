FROM golang:latest as buildImage
        
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

FROM apline:latest

WORKDIR /app

COPY --from=buildImage /app/main .

EXPOSE 3000

CMD ["./main"]