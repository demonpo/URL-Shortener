FROM golang:1.22 as common-build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy

# Utils
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go install github.com/cosmtrek/air@latest

# Build
RUN task build
RUN chmod +x ./app

EXPOSE 3000



FROM common-build-stage as development-build-stage

ENV NODE_ENV development

CMD ["air", "--build.cmd", "go build -o app ./src/app.go", "--build.bin", "./app"]

FROM common-build-stage as production-build-stage

ENV NODE_ENV production

CMD ["./app"]