# Microservices with Go

The Microservices Architecture contains

-   Broker Service
-   Authentication Service
-   Logger Service
-   Mail Service
-   Listener Service AMQP with RabbitMQ

## GNU Make

-   [www.gnu.org/software/make](https://www.gnu.org/software/make/) is a tool which controls the generation of executables and other non-source files of a program from the program's source files
-   Make gets its knowledge of how to build your program from a file called the `makefile`, which lists each of the non-source files and how to compute it from other files. When you write a program, you should write a makefile for it, so that it is possible to use Make to build and install the program
-   All the `Docker` images control for

    -   `Linux/MacOS` are in the `project/Makefile` file
    -   `Windows` are in the `project/Makefile.windows` file

## Frontend Test

-   The Test Page is used for test the Services in this Architecture
-   The Test Page will send an example request to the specific Service that user choose and then receive the response from that Service
-   Technologies
    -   [bootstrap](https://getbootstrap.com/docs/5.2/getting-started/introduction/)
    -   [javaScript](https://www.javascript.com) for fetching data from the services
    -   [html](https://developer.mozilla.org/en-US/docs/Learn/HTML)

## Broker Service

-   <strong>Broker Service</strong> is an entry point for redirecting the request to appropriate service and receiving the response from that service
-   Package index
    -   [github.com/go-chi/chi](https://github.com/go-chi/chi) is a lightweight, idiomatic and composable router for building `Go` HTTP services
        -   Installation
            ```sh
            go get github.com/go-chi/chi/v5
            go get github.com/go-chi/chi/middleware
            go get github.com/go-chi/cors
            ```
    -   [pkg.go.dev/net/rpc](https://pkg.go.dev/net/rpc) provides access to the exported methods of an object across a network or other I/O connection. A server registers an object, making it visible as a service with the name of the type of the object. After registration, exported methods of the object will be accessible remotely. A server may register multiple objects (services) of different types but it is an error to register multiple objects of the same type.
-   File `broker-service/broker-service.dockerfile` is the `dockerfile` for the service

## Authentication Service with PostgreSQL

-   <strong>Authentication Service</strong> is an API for determining if the credentials from the request body is matching the data in the database. All the credentials are stored in a PostgreSQL image
-   Package index
    -   [github.com/go-chi/chi](https://github.com/go-chi/chi) is a lightweight, idiomatic and composable router for building `Go` HTTP services
        -   Installation
            ```sh
            go get github.com/go-chi/chi/v5
            go get github.com/go-chi/chi/middleware
            go get github.com/go-chi/cors
            ```
    -   [github.com/jackc/pgconn](https://github.com/jackc/pgconn) is a low-level `PostgreSQL` database driver. It operates at nearly the same level as the `C` library `libpq`
        -   Installation
            ```sh
            go get github.com/jackc/pgconn
            ```
    -   [github.com/jackc/pgx](https://github.com/jackc/pgx) is a higher level libraries, high performance interface that exposes PostgreSQL-specific features such as `LISTEN` / `NOTIFY` and `COPY`. It also includes an adapter for the standard `database/sql` interface. The toolkit component is a related set of packages that implement `PostgreSQL` functionality such as parsing the wire protocol and type mapping between `PostgreSQL` and `Go`
        -   Installation
            ```sh
            go get github.com/jackc/pgx/v4
            go get github.com/jackc/pgx/v4/stdlib
            ```
    -   [github.com/golang/crypto/bcrypt](https://github.com/golang/crypto/blob/master/bcrypt/bcrypt.go) is used for hashing passwords
    -   [pkg.go.dev/database/sql](https://pkg.go.dev/database/sql) provides a generic interface around SQL (or SQL-like) databases. The sql package must be used in conjunction with a database driver. See [golang.org/s/sqldrivers](https://golang.org/s/sqldrivers) for a list of drivers
-   File `authentication-service/authentication-service.dockerfile` is the `dockerfile` for the service

## Logger Service with MongoDB

-   <strong>Logger Service</strong> is an API for saving logs whenever one Service receives and processes a response. All the `LogEntry` struct contains 2 fields Name and Data. The collections are stored in a MongoDB image
-   Package index
    -   [github.com/go-chi/chi](https://github.com/go-chi/chi) is a lightweight, idiomatic and composable router for building `Go` HTTP services
        -   Installation
            ```sh
            go get github.com/go-chi/chi/v5
            go get github.com/go-chi/chi/middleware
            go get github.com/go-chi/cors
            ```
    -   [github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver) is the `MongoDB` supported driver for `Go`
        -   Installation
            ```sh
            go get go.mongodb.org/mongo-driver/mongo
            go get go.mongodb.org/mongo-driver/mongo/options
            ```
    -   [pkg.go.dev/net/rpc](https://pkg.go.dev/net/rpc) provides access to the exported methods of an object across a network or other I/O connection. A server registers an object, making it visible as a service with the name of the type of the object. After registration, exported methods of the object will be accessible remotely. A server may register multiple objects (services) of different types but it is an error to register multiple objects of the same type.
-   File `logger-service/logger-service.dockerfile` is the `dockerfile` for the service

## Mail Service

-   <strong>Mail Service</strong> connects directly with <strong>Broker Service</strong> in the development version which you shouldn't do in production. In production, <strong>Mail Service</strong> can not be connected by User, just be connected to other Service except Broker Service. The Service sends tested mails to the MailHog server
-   Package index
    -   [github.com/go-chi/chi](https://github.com/go-chi/chi) is a lightweight, idiomatic and composable router for building `Go` HTTP services
        -   Installation
            ```sh
            go get github.com/go-chi/chi/v5
            go get github.com/go-chi/chi/middleware
            go get github.com/go-chi/cors
            ```
    -   [github.com/mailhog/MailHog](https://github.com/mailhog/MailHog) is an email testing tool for developers
        -   Overview
            -   Configure your application to use `MailHog` for SMTP delivery
            -   View messages in the web UI, or retrieve them with the JSON API
            -   Optionally release messages to real SMTP servers for delivery
        -   Installation
            ```sh
            go get github.com/mailhog/MailHog
            ```
        -   The local server of `MailHog` runs on `http://localhost:1025`, check this server for web UI and see all the tested mails in the `Inbox` section
        -   A `Docker` image for `MailHog` is configured in file `project/docker-compose.yml`
    -   [github.com/vanng822/go-premailer](github.com/vanng822/go-premailer) is an inline styling for HTML mail in `Go`
        -   Styling mail with both `HTML` and `plain` formats before sending
        -   Installation
            ```sh
            go get github.com/vanng822/go-premailer/premailer
            ```
    -   [github.com/xhit/go-simple-mail](https://github.com/xhit/go-simple-mail) is the best way to send emails in `Go` with SMTP Keep Alive and Timeout for Connect and Send
        -   You can find more information in the [documentation](https://pkg.go.dev/github.com/xhit/go-simple-mail/v2)
        -   Installation
            ```sh
            go get github.com/xhit/go-simple-mail/v2
            ```
-   File `mail-service/mail-service.dockerfile` is the `dockerfile` for the service

## Listener Service AMQP

-   <strong>Listener Service AMQP</strong> listens for any data (requests) pushed to RabbitMQ server to consume it as soon as possible
-   Any requests after being routed by <strong>Broker Service</strong> is pushed directly to RabbitMQ server. The <strong>Listener Service AMQP</strong> connects to the RabbitMQ server and listens to any requests in the message queue and then sends requests to the appropriate Services
-   Package index
    -   [github.com/rabbitmq/amqp091-go](https://github.com/rabbitmq/amqp091-go) is a `Go AMQP 0.9.1` client maintained by the [RabbitMQ core team](https://github.com/rabbitmq). The package provides a functional interface that closely represents the `AMQP 0.9.1` model targeted to `RabbitMQ` as a server
        -   Installation
            ```sh
            go get github.com/rabbitmq/amqp091-go
            ```
-   File `listener-service/listener-service.dockerfile` is the `dockerfile` for the service

## RabbitMQ Server

-   Image for the `RabbitMQ` server is [`rabbitmq:3.9-alpine`](https://github.com/docker-library/rabbitmq/blob/6889979f517c7ea3a7bd54bb88864dc8c29d327c/3.9/alpine/Dockerfile), runs on port `5672` on `docker` server
-   Uses the `topics` as exchange type
-   Does not delete data until it is consumed successfully
-   References
    -   Find more information about `RabbitMQ` at [RabbitMQ website](https://www.rabbitmq.com)
    -   [Godoc API reference](https://pkg.go.dev/github.com/rabbitmq/amqp091-go)
    -   [RabbitMQ tutorials in Go](https://github.com/rabbitmq/rabbitmq-tutorials/tree/main/go)

## Communicating between Services using RPC

-   <strong>Broker Service</strong> has 2 options to communicate with <strong>Logger Service</strong>
    -   `logEventViaRabbit(w http.ResponseWriter, l LogPayload)` just pushes requests from the client to the `RabbitMQ` Server for <strong>Listener Service</strong> to consume
    -   `logItemViaRPC(w http.ResponseWriter, l LogPayload)` connects to the `RPC` Server of <strong>Logger Service</strong> on port `5001` and calls the function of the `RPC` Server using [pkg.go.dev/net/rpc](https://pkg.go.dev/net/rpc) package
        -   When it comes to the `RPC` option, the <strong>Logger Service</strong> has to always listen to the RPC requests
        -   <strong>Broker Service</strong> and <strong>Logger Service</strong> use [pkg.go.dev/net/rpc](https://pkg.go.dev/net/rpc) package to communicate. The <strong>Broker Service</strong> is the client and the <strong>Logger Service</strong> is the server

## Docker usage

-   All the `Docker` configurations are in the `project/docker-compose.yml` file

-   Build all the `Docker` images and services's binaries
    ```sh
    make up-build
    ```
-   Build one `Docker` image of one specific service
    ```sh
    make service-build
    ```
-   Pull and Start all the `Docker` images
    ```sh
    make up
    ```
-   Stop docker compose
    ```sh
    make down
    ```
-   Start the `frontend` binary listening on `http://localhost:80`
    ```sh
    make start
    ```
-   Stop the `frontend` binary listening on `http://localhost:80`
    ```sh
    make stop
    ```

## Contributor

-   Minh Tran (Me)
