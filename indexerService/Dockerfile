FROM golang:1.16.5-stretch
WORKDIR /serverHome
COPY ./go.mod /serverHome/
COPY ./go.sum /serverHome/
COPY ./main.go /serverHome/
RUN go mod tidy