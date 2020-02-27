FROM golang:latest AS build
LABEL maintainer="nasa9084"

WORKDIR /go/src/github.com/nasa9084/reqbin
COPY . .
RUN CGO_ENABLED=0 go build -o reqbin main.go

FROM scratch

COPY --from=build /go/src/github.com/nasa9084/reqbin/reqbin /reqbin
CMD ["/reqbin"]
