FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go build -o privateness-mcp-app ./cmd
CMD ["./privateness-mcp-app"]
