# Step 1: Use the golang official image as a base
FROM golang:1.22.3-alpine as builder

# Step 2: Set current directory inside the container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download all dependencies
RUN go mod download

# Step 5: Copy the source code into the container
COPY . .

# Step 6: Build the Go App targetting main.go inside cmd directory
RUN go build -o main ./cmd/main.go

# Step 7: Use a smaller base image for production
FROM alpine:3.18

# Step 8: Set the current directory inside the container
WORKDIR /app

# Step 9: Copy the pre-built binary file from previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Step 10: Expose service PORT
EXPOSE 8082

# Step 11: Command to run the executable file
CMD [ "./main" ]