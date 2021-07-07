FROM golang:1.16.5-stretch
WORKDIR /serverHome
COPY ./go.mod /serverHome/
COPY ./go.sum /serverHome/
COPY ./main.go /serverHome/
COPY ./controllers/ /serverHome/controllers/
COPY ./models/ /serverHome/models/
COPY ./routes/ /serverHome/routes/
RUN go mod tidy