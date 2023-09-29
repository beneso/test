FROM golang:1.21 as build
WORKDIR /src
COPY server.go .
# https://mt165.co.uk/blog/static-link-go/
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /server ./server.go

FROM scratch
COPY --from=build /server /server
COPY content/ /content
CMD ["/server"]
