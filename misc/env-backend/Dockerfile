FROM centos:7

# update centos
RUN yum -y update
RUN yum -y install wget
RUN yum -y install git
RUN yum -y install vim

WORKDIR /develop

# install golang
RUN wget -nv https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.10.2.linux-amd64.tar.gz

# Set Environment for go-lang. 
# Note that these environments have to be set before using go get command.
ENV GOPATH=/develop
ENV PATH $PATH:/usr/local/go/bin