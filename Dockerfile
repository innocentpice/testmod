FROM golang:latest

ARG APP_ENV=Production
ARG SERVICE_NAME=APP

ENV APP_ENV=${APP_ENV}

COPY ./ /go/src/${SERVICE_NAME}

WORKDIR /go/src/${SERVICE_NAME}

RUN go mod download

RUN if [ "$APP_ENV" = "Production" ];\
    then \
        go get github.com/githubnemo/CompileDaemon; \
    else \
        go get -u github.com/cosmtrek/air ;\
fi

ENTRYPOINT if [ "$APP_ENV" = "Production" ];\
    then \
        CompileDaemon --build="go build ./main.go" --command=./main; \
    else \
        air ;\
fi
