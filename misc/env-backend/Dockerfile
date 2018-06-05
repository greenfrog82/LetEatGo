FROM centos:7

# update centos
RUN yum -y update
RUN yum -y install wget
RUN yum -y install git

# install golang
RUN wget -nv https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.10.2.linux-amd64.tar.gz
ENV GOPATH /root/go
ENV PATH $PATH:/usr/local/go/bin

# install gin gonic
RUN go get github.com/gin-gonic/gin
