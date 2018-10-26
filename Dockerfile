FROM golang

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/go-sketch

RUN go get -u github.com/gorilla/mux
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app;
	
EXPOSE 8080