![gotham](https://user-images.githubusercontent.com/39353278/103892416-99f6c880-50fc-11eb-8869-af197ca81fd1.png)

## Repos and frameworks used

- Dependency Injection Container ( https://github.com/sarulabs/dingo )
- Handler ( https://echo.labstack.com )
- ORM (https://gorm.io)
- Validation (https://github.com/go-ozzo/ozzo-validation)
- Cron (https://github.com/jasonlvhit/gocron)

## Table of contents

![GitHub](https://img.shields.io/github/license/tolgaOzen/go-gotham)
![GitHub top language](https://img.shields.io/github/languages/top/tolgaozen/go-gotham)
![GitHub last commit](https://img.shields.io/github/last-commit/tolgaozen/go-gotham)

- [Basic Usage](#basic-usage)
- [Di Container](#di-container)
    * [Providers](#providers)
    * [Di Scopes](#di-scopes)
    * [Add New Service](#add-new-service)
- [Database](#database)
    * [Supported Databases](#supported-databases)
    * [Transactions](#transactions)
    * [Db Scopes](#db-scopes)
- [ORM](#orm)
- [Requests](#requests)
    * [Create Requests](#create-new-requests)
    * [Bind Requests](#bind-requests)
    * [Validations](#validations)
    * [Rules](#rules)
- [Routes](#routes)
- [Controllers](#controllers)
- [Jobs](#jobs)
- [Helpers](#helpers)
    * [Response](#response)
- [Middlewares](#conditional-middlewares)
    * [Create](#create)
    * [Conditional Middlewares](#conditional-middlewares)
- [Features To Be Added Soon](#features-to-be-added-soon)

## Basic Usage


## Di Container

### Providers

### Di Scopes

### Add New Service


## Database

### Supported Databases

### Transactions

### Db Scopes


## ORM

## Requests

### Create New Requests

### Bind Requests

### Validations

### Rules

## Routes


## Controllers

## Jobs

## Helpers

## Middlewares

### Create

Creating a file in the middleware folder

#### Example

```go
type Example struct{}

func (e Example) control(c echo.Context) (bool bool, err error) {
if c.IsWebSocket() {
return true, nil
}
return false, errors.New("is it not webSocket")
}
```

### Conditional Middlewares

#### OR

```go
r.GET("/users/:user", controllers.UserController{}.Show, GMiddleware.Or([]GMiddleware.MiddlewareI{GMiddleware.IsAdmin{}, GMiddleware.IsVerified{}}))
```

Authenticated user must be admin or verified

#### AND

```go
r.GET("/users/:user", controllers.UserController{}.Show, GMiddleware.And([]GMiddleware.MiddlewareI{GMiddleware.IsAdmin{}, GMiddleware.IsVerified{}}))
```

Authenticated user must be admin and verified

## Features To Be Added Soon

- Database seeder
- Unit testing

## Author

> Tolga Özen

> mtolgaozen@gmail.com

## License

MIT License

Copyright (c) 2021 Tolga Özen

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.