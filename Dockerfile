FROM golang

EXPOSE 8080 80 443

# Set pwd to the go folder
WORKDIR ${GOPATH}

RUN apt-get update -y && \
	apt-get install -y \
	nano

RUN go get github.com/gorilla/mux && go get github.com/go-sql-driver/mysql

COPY ${PROJECT_PATH} /go/cms

RUN cd /go/cms && go build

EXPOSE 8080

CMD ["./cms"]