# Build sample
# docker build --tag vuejs:1

# Run sample
# docker run -it --name vuejs-dev -v /home/seunghoryu/dev/vuejs:/root/vuejs vuejs:1 /bin/bash
FROM centos:7

# update centos
RUN yum -y update kernel
RUN yum -y update

# Install npm
RUN yum -y install epel-release
RUN yum -y install npm

# Install vuejs
RUN npm install vue
RUN npm install vue-cli -g

RUN yum clean all

# volume
RUN mkdir /root/vuejs
VOLUME ["/root/vuejs"]

