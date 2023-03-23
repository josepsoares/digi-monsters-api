FROM golang:1.19

# Move to working directory (/build).
WORKDIR /app

# Copy and download dependency using go mod.
COPY . .
RUN go mod download

# Set necessary environment variables needed 
# for our image and build the consumer.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build ./main.go

EXPOSE 3002

CMD [ "./main" ]