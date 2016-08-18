FROM ubuntu:14.04

RUN apt-get update && apt-get install --no-install-recommends -y \
    apt-transport-https \
    ca-certificates \
    curl \
    git-core

RUN curl -s https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz | tar -C /usr/local -xz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:/usr/local/bin:$PATH

RUN go get github.com/djangoxv/gotz
RUN go install github.com/djangoxv/gotz

CMD /go/bin/gotz
EXPOSE 80
