# Build sample
# docker build --tag vuejs:1

# Run sample
# docker run -it --name vuejs-dev -v /home/seunghoryu/dev/vuejs:/root/vuejs vuejs:1 /bin/bash
FROM centos:7

# update centos
RUN yum -y update kernel
RUN yum -y update

# Install nodejs 8
# RUN yum -y install epel-release
RUN curl -sL https://rpm.nodesource.com/setup_8.x | bash -
RUN yum -y install nodejs

RUN yum -y install git
RUN npm install vue
RUN npm install vue-cli -g

# volume
RUN mkdir /root/vuejs
VOLUME ["/root/vuejs"]

