FROM centos:7

# update centos
RUN yum -y update
RUN yum -y install wget
RUN yum -y install git
RUN yum -y install vim

WORKDIR /root

# install golang
RUN wget -nv https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.10.2.linux-amd64.tar.gz

# Set Environment for go-lang. 
# Note that these environments have to be set before using go get command.
ENV GOPATH=/develop
ENV PATH $PATH:/usr/local/go/bin
ENV PATH $PATH:/develop/bin

WORKDIR /develop/src

# install go framework
RUN go get github.com/gin-gonic/gin
RUN go get github.com/codegangsta/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/derekparker/delve/cmd/dlv

