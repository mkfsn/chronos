############################################
# Stage 0, based on golang, to build drone #
############################################
FROM golang:1.10 AS build-layer
MAINTAINER p408865@gmail.com

ENV GOSRC=/go/src
ENV REPO=github.com/mkfsn/chronos

# Install dependency
RUN go get -u github.com/kardianos/govendor

# copy cloud base stuff
RUN mkdir -p ${GOSRC}/${REPO}/vendor/
COPY ./vendor/vendor.json ${GOSRC}/${REPO}/vendor/
WORKDIR ${GOSRC}/${REPO}
RUN govendor sync

COPY . ${GOSRC}/${REPO}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app -v

ENTRYPOINT ["/go/src/github.com/mkfsn/chronos/app"]

###############################
# Final Stage, to run the app #
###############################
FROM broady/cacerts AS final-layer
WORKDIR /go/bin
COPY --from=build-layer /go/src/github.com/mkfsn/chronos/app /go/bin/
ENTRYPOINT ["/go/bin/app"]
