FROM golang:1.7
ADD ./gotz /go/bin/gotz
CMD /go/bin/gotz
EXPOSE 8080
