# docker build --tag myimage:1
# docker run -i -t --name dev myimage:1 /bin/bash
FROM centos:7

# update centos
RUN yum -y update kernel
RUN yum -y update
RUN yum -y install wget
RUN yum clean all

# volume
RUN mkdir /root/go
VOLUME ["/root/go"]

# install golang
RUN wget -nv https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.10.2.linux-amd64.tar.gz
ENV GOPATH /root/go
ENV PATH $PATH:/usr/local/go/bin

