FROM golang:1.16.15-alpine3.15 AS fabconnect-builder
RUN apk add make git
ENV GODEBUG='x509ignoreCN=0'
RUN echo 'GODEBUG=x509ignoreCN=0' > ~/.profile
WORKDIR /fabconnect
ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN make build

FROM alpine:latest
WORKDIR /fabconnect
COPY --from=fabconnect-builder /fabconnect/fabconnect ./
ADD ./openapi ./openapi/
RUN ln -s /fabconnect/fabconnect /usr/bin/fabconnect
ENV GODEBUG='x509ignoreCN=0'
RUN echo 'GODEBUG=x509ignoreCN=0' > ~/.profile
ENTRYPOINT [ "fabconnect" ]
