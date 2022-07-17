FROM golang:1.18
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
# RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/main.go
RUN go build -o /bin/app ./cmd/main.go

# FROM scratch
# COPY --from=0 /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
