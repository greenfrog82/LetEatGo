# Build sample
# docker build --tag go:1

# Run sample
# docker run -it --name go-dev -v /home/seunghoryu/dev/golang:/root/go go:1 /bin/bash
FROM centos:7

# update centos
RUN yum -y update
RUN yum -y install wget

# volume
RUN mkdir /root/go

# install golang
RUN wget -nv https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.10.2.linux-amd64.tar.gz
ENV GOPATH /root/go
ENV PATH $PATH:/usr/local/go/bin

