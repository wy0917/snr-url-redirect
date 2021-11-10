FROM golang:alpine as builder
WORKDIR /go/src/snr-url-redirect
COPY . .
ARG REDIRECT_URL
RUN GOOS=linux go build -ldflags "-X main.RedirectURL=${REDIRECT_URL}" -o /go/bin/server .

FROM alpine:latest
# copy the binary to this container
COPY --from=builder /go/bin/server /server
EXPOSE 8080
CMD ["/server"]
