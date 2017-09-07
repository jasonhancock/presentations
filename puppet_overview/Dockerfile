FROM centos:7

# Install the RPM to get the GPG key securely, then overwrite the epel repo config
RUN rpm -ivh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm

RUN yum -y install \
  git \
  graphviz \
  make \
  psmisc \
  supervisor \
  wget \
  && yum clean all

ENV goversion 1.9
ENV gofile go${goversion}.linux-amd64.tar.gz
ENV gourl https://storage.googleapis.com/golang/${gofile}
ENV GOPATH /go

RUN wget -q -O /usr/local/${gofile} ${gourl} \
  && mkdir /usr/local/go \
  && tar -xzf /usr/local/${gofile} -C /usr/local/go --strip 1

# supervisor configs
ADD supervisord.conf /etc/supervisord.conf
ADD *.ini /etc/supervisord.d/
ADD run.sh /
RUN chmod 0755 /run.sh

EXPOSE 3999
ENV PATH /usr/local/go/bin:${GOPATH}/bin:${PATH}

RUN go get golang.org/x/tools/cmd/present
RUN go get github.com/jmhodges/justrun

CMD /run.sh
