FROM golang:1.6.3
ADD ./gotz /go/bin/gotz
CMD /go/bin/gotz
EXPOSE 8080
