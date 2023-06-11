FROM golang:1.18.3

WORKDIR /portfolio

COPY . .

# COPY ./src /home

# COPY ./cmd /home/cmd

# COPY go.mod /home

# COPY go.sum /home

RUN --mount=type=secret,id=_env,dst=/etc/secrets/.env cat /etc/secrets/.env

RUN  go build -o portfolio ./cmd/main.go

CMD ["/portfolio/cmd/portfolio"]

