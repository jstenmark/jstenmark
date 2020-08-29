# Use the Ubuntu image
FROM golang:1.20

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  fortune-mod \
  fortunes \
  && rm -rf /var/lib/apt/lists/*

ENV PATH="/usr/games:${PATH}"
WORKDIR /app

COPY scripts/update_fortune.go .
COPY README.md .

#RUN go mod tidy
CMD ["sh", "-c", "go run update_fortune.go -readme /app/README.md && cat README.md"]