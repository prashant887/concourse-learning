FROM golang 
ADD hello-world.go / 
ENV GO111MODULE=on
RUN go mod download
RUN go build -o hello-world.go
CMD ["tail -f"]