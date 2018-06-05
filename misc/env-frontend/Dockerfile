FROM centos:7

RUN yum -y update
RUN curl -sL https://rpm.nodesource.com/setup_8.x | bash -
RUN yum -y install nodejs

RUN yum -y install git
RUN npm install vue
RUN npm install vue-cli -g