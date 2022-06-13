FROM alpine:latest

COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/
 
ENV PATH="/usr/local/go/bin:${PATH}"

# Variables that will setup a non-root user.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Copy my devcontainer setup script to the container
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# RUN go get -v golang.org/x/tools/gopls