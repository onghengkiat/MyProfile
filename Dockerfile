# Base image for building the go project
FROM golang:1.14-alpine AS build

# Updates the repository and installs git
RUN apk update && apk upgrade && apk add --no-cache git

## If you have a go.mod and go.sum file in your project, uncomment lines 13, 14, 15

# Switches to /tmp/app as the working directory, similar to 'cd'
WORKDIR /tmp/app
COPY . .

# Builds the current project to a binary file called api
# The location of the binary file is /tmp/app/out/api
RUN GOOS=linux go build -o ./out/api .

#########################################################

# The project has been successfully built and we will use a
# lightweight alpine image to run the server 
FROM alpine:latest

# Adds CA Certificates to the image
RUN apk add ca-certificates

# Copies the binary file from the BUILD container
COPY --from=build /tmp/app/out/api /app
COPY --from=build /tmp/app .

EXPOSE 8080

# Runs the binary once the container starts
ENTRYPOINT ["/app"]