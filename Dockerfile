FROM golang:latest

ARG SERVICE_NAME

WORKDIR /workspace
ADD ./go.work .
ADD ./cartola-commons ./cartola-commons
ADD ./$SERVICE_NAME ./$SERVICE_NAME

WORKDIR /workspace/$SERVICE_NAME 

RUN go test -v ./...
RUN if [ $? -eq 1 ]; then exit 1; fi

RUN go build -o app

RUN chmod +x ./app

RUN ./app

EXPOSE 3000