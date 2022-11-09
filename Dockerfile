FROM golang:buster as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder /app/maximum-portage /bin/
ENTRYPOINT [ "/bin/maximum-portage" ]