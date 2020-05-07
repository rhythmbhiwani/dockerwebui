FROM ubuntu
MAINTAINER rhythmbhiwani@gmail.com
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y git wget docker.io
RUN cd /tmp
RUN wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
RUN tar -xvf go1.11.linux-amd64.tar.gz
RUN mv go /usr/local
ENV GOROOT=/usr/local/go
ENV GOPATH=$HOME/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN /bin/bash -c "source ~/.profile"
RUN mkdir -p $GOPATH/src/webui/
COPY webui $GOPATH/src/webui/
WORKDIR $HOME/go/src/webui
RUN go build
RUN chmod +x webui
EXPOSE 80
ENTRYPOINT ./webui