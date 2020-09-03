FROM golang:alpine as builder


# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /jobs

COPY . ./

RUN go build -mod vendor -o /bin/jobs cmd/jobs.go

FROM alpine

RUN addgroup -S www && adduser -S client -G www
USER client

COPY --from=builder /jobs/config/ config
COPY --from=builder /bin/jobs jobs

CMD ["./jobs"]