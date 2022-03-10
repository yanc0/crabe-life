FROM golang AS build
COPY main.go ./
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o crabe-life main.go

FROM scratch AS run
COPY --from=build /go/crabe-life /crabe-life
EXPOSE 8080
ENTRYPOINT ["/crabe-life"]