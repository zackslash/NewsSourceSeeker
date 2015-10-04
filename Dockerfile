# Dockerfile for News Source Seeker
# https://github.com/zackslash/NewsSourceSeeker

FROM google/golang:1.4

MAINTAINER Luke Hines <lukehines1+opensource@gmail.com>

ADD . /proj

RUN cd proj && \
	make build

WORKDIR /proj/bin

ENTRYPOINT ["/proj/bin/News_Source_Seeker"]