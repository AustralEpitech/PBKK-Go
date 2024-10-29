ARG VER=1.23

FROM "docker.io/golang:$VER" as build
WORKDIR /build/
ARG VER
COPY main.go .
RUN printf "module main\ngo $VER" > go.mod && CGO_ENABLED=0 go build -o /app

FROM scratch
COPY --from=build /app /app
COPY static/ /static/
COPY html/ /html/
EXPOSE 3000
CMD ["/app"]
