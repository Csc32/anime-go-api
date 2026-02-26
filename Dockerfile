FROM golang:1.25-alpine

# 1. Set the working directory inside the container
WORKDIR /app

# 2. Install Air for hot-reloading
RUN go install github.com/air-verse/air@latest

# 3. Copy only dependency files first (for faster builds)
COPY go.mod go.sum* ./
RUN go mod download

# 4. Copy the rest of the project
COPY . .

# 5. Run Air by default
CMD ["air"]
