FROM ubi8/go-toolset as build
COPY . .

RUN go mod init mikroserwis2.go && \
    go mod tidy
RUN pwd    
RUN go build mikroserwis2.go
RUN ls

FROM ubi8/ubi-micro
COPY --from=build /opt/app-root/src/mikroserwis2 .
CMD ./mikroserwis2
