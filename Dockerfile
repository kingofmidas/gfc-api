FROM golang:1.14

RUN apt-get update; \
    apt-get -y install postgresql-client; \
    apt-get -y install curl; \
    apt-get -y install netcat

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.13.0/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin/ --transform=s/migrate.linux-amd64/migrate/

WORKDIR /src/gfc-api/

COPY . /src/gfc-api/

RUN chmod +x ./wait-for-postgres.sh

RUN cd ./app && go build -o apiserver cmd/main.go

ENTRYPOINT [ "./wait-for-postgres.sh" ]