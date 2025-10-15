# Step 1: Use official Go image
FROM golang:1.25-alpine


# Step 2: Set working directory
WORKDIR /app

# Step 3: Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy all project files
COPY . .

# Step 5: Build the Go binary
RUN go build -o file-analyser

# Step 6: Expose the port your app runs on
EXPOSE 8005

# Step 7: Run the app
CMD ["./file-analyser"]
