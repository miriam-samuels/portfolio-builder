FROM golang:1.18.3

WORKDIR /portfolio

COPY . .

# COPY ./src /home

# COPY ./cmd /home/cmd

# COPY go.mod /home

# COPY go.sum /home

RUN  go build -o portfolio ./cmd/main.go

CMD ["/portfolio/cmd/portfolio"]
