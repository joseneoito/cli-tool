FROM golang:latest
WORKDIR  /app

COPY . .
RUN go build -o cli-tool . 
CMD ["./cli-tool"]
