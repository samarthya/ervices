FROM golang:latest as build
WORKDIR /service
COPY src/ws/go.mod src/ws/go.mod
COPY src/ws/ws.go src/ws/ws.go

COPY src/main.go src/main.go
COPY src/go.mod src/go.mod

ENV WS_GO_PORT 8181
ENV VERSION_IMAGE 2.0
WORKDIR src/
# RUN ls
RUN CGO_ENABLED=0 go build -o /bin/main

FROM scratch as bin
COPY --from=build /bin/main /bin/server
EXPOSE $WS_GO_PORT
ENTRYPOINT [ "/bin/server" ] 
