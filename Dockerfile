# -----------------------------------------------------------------------------
# socle
FROM alpine:3.20.1 As socle

RUN apk add go

WORKDIR /src

# -----------------------------------------------------------------------------
# builder Env
FROM socle as builder

RUN apk add make

# Add go lang linter
WORKDIR /usr/lib/go
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.62.0
RUN ln -s /usr/lib/go/bin/golangci-lint /usr/bin/golangci-lint

WORKDIR /src

CMD ["make"]


# -----------------------------------------------------------------------------
# Execution
FROM builder As executor

ARG EXEC_NAME
ENV EXEC_NAME $EXEC_NAME

COPY $EXEC_NAME .

# Create a script to pass command line args
RUN echo "./"$EXEC_NAME" \$@" > /run.sh

ENTRYPOINT ["/bin/sh", "/run.sh"]