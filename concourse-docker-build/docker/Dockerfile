FROM golang 
ADD hello-world.go / 
RUN go mod init
RUN go build -o hello-world.go
CMD ["tail -f"]