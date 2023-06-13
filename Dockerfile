FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# First install go things
RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest && \
# Than install the rest
    apt-get update && \
    apt-get install sqlite3 -y && \
# Then do the permissions stuff
    usermod -u 1000 www-data && \
    mkdir -p /var/www/.cache && \
    chown -R www-data:www-data /go && \
    chown -R www-data:www-data /var/www/.cache

USER www-data

CMD ["tail", "-f", "/dev/null"]