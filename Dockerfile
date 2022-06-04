FROM  golang AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /yads

##
## Deploy
##

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /yads /yads

EXPOSE 80

USER nonroot:nonroot
ENTRYPOINT [ "/yads" ]