# Use a small Go base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy code
COPY . .

# Build the app
RUN go build -o purge-controller .

# Run the app
CMD ["./purge-controller"]
