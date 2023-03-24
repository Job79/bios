FROM alpine:edge as builder
RUN apk add go # prepare dependencies

# copy code into container
COPY . /build
WORKDIR /build/cmd/bios

# build binary
RUN GOOS=linux go build '.' 



FROM alpine:edge

# empty config file, custom config files can be mounted to container using -v
RUN echo '{}' > /app/conf.json 

# copy binary from builder into container
COPY --from=builder /build/cmd/bios/bios /app/bios

EXPOSE 8080/tcp
ENTRYPOINT ["./app/bios"]
