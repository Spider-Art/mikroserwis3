FROM ubi8/go-toolset as build
COPY . .

RUN go mod init mikroserwis3.go && \
    go mod tidy
RUN pwd    
RUN go build mikroserwis3.go
RUN ls

FROM ubi8/ubi-micro
COPY --from=build /opt/app-root/src/mikroserwis3 .
CMD ./mikroserwis3
