# radiant [![Test](https://github.com/W3-Engineers-Ltd/Radiant/actions/workflows/test.yml/badge.svg?branch=develop)](https://github.com/W3-Engineers-Ltd/Radiant/actions/workflows/test.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/W3-Engineers-Ltd/Radiant)](https://goreportcard.com/report/github.com/W3-Engineers-Ltd/Radiant) [![Go Reference](https://pkg.go.dev/badge/github.com/W3-Engineers-Ltd/Radiant.svg)](https://pkg.go.dev/github.com/W3-Engineers-Ltd/Radiant)

radiant is used for rapid development of enterprise application in Go, including RESTful APIs, web apps and backend services.

It is inspired by Tornado, Sinatra and Flask. radiant has some Go-specific features such as interfaces and struct embedding.

![architecture](https://cdn.nlark.com/yuque/0/2020/png/755700/1607857489109-1e267fce-d65f-4c5e-b915-5c475df33c58.png)

radiant is composed of four parts:

1. Base modules: including log module, config module, governor module;
2. Task: is used for running timed tasks or periodic tasks;
3. Client: including ORM module, httplib module, cache module;
4. Server: including web module. We will support gRPC in the future;

**Please use RELEASE version, or master branch which contains the latest bug fix**

## Quick Start

[Official website](http://radiant.me)

[Example](https://github.com/W3-Engineers-Ltd/Radiant-example)

> If you could not open official website, go to [beedoc](https://github.com/radiant/beedoc)

### Web Application

![Http Request](https://cdn.nlark.com/yuque/0/2020/png/755700/1607857462507-855ec543-7ce3-402d-a0cb-b2524d5a4b60.png)

#### Create `hello` directory, cd `hello` directory

    mkdir hello
    cd hello

#### Init module

    go mod init

#### Download and install

    go get github.com/W3-Engineers-Ltd/Radiant

#### Create file `hello.go`

```go
package main

import "github.com/W3-Engineers-Ltd/Radiant/server/web"

func main() {
	
	web.Run()
}
```

#### Build and run

    go build hello.go
    ./hello

#### Go to [http://localhost:8080](http://localhost:8080)

Congratulations! You've just built your first **radiant** app.

## Features

* RESTful support
* [MVC architecture](https://github.com/radiant/beedoc/tree/master/en-US/mvc)
* Modularity
* [Auto API documents](https://github.com/radiant/beedoc/blob/master/en-US/advantage/docs.md)
* [Annotation router](https://github.com/radiant/beedoc/blob/master/en-US/mvc/controller/router.md)
* [Namespace](https://github.com/radiant/beedoc/blob/master/en-US/mvc/controller/router.md#namespace)
* [Powerful development tools](https://github.com/radiant/bee)
* Full stack for Web & API

## Modules

* [orm](https://github.com/radiant/beedoc/tree/master/en-US/mvc/model)
* [session](https://github.com/radiant/beedoc/blob/master/en-US/module/session.md)
* [logs](https://github.com/radiant/beedoc/blob/master/en-US/module/logs.md)
* [config](https://github.com/radiant/beedoc/blob/master/en-US/module/config.md)
* [cache](https://github.com/radiant/beedoc/blob/master/en-US/module/cache.md)
* [context](https://github.com/radiant/beedoc/blob/master/en-US/module/context.md)
* [admin](https://github.com/radiant/beedoc/blob/master/en-US/module/admin.md)
* [httplib](https://github.com/radiant/beedoc/blob/master/en-US/module/httplib.md)
* [task](https://github.com/radiant/beedoc/blob/master/en-US/module/task.md)
* [i18n](https://github.com/radiant/beedoc/blob/master/en-US/module/i18n.md)

## Community

* [http://radiant.me/community](http://radiant.me/community)
* Welcome to join us in Slack: [https://radiant.slack.com invite](https://join.slack.com/t/radiant/shared_invite/zt-fqlfjaxs-_CRmiITCSbEqQG9NeBqXKA),
* QQ Group Group ID:523992905
* [Contribution Guide](https://github.com/radiant/beedoc/blob/master/en-US/intro/contributing.md).

## License

Radiant source code is licensed under the GNU LESSER GENERAL PUBLIC LICENSE, Version 3.0
