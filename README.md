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

- [Architecture](#architecture)
    * [Di Container](#di-container)
    * [Providers](#providers)
    * [Defs](#defs)
    * [Services](#services)
    * [Di Scopes](#di-scopes)
- [Database](#database)
    * [Supported Databases](#supported-databases)
    * [Procedures](#procedures)
    * [Db Scopes](#db-scopes)
    * [Migrations](#migrations)
- [ORM](#orm)
- [Requests](#requests)
    * [Create Requests](#create-new-requests)
    * [Bind Request And Validate](#bind-request-and-validate)
    * [More Info For Validations Rules](#more-info-for-validations-rules)
    * [Custom Rules](#custom-rules)
- [Routes](#routes)
- [Auth](#auth)
    * [JWT](#jwt)
- [Controllers](#controllers)
- [Jobs](#jobs)
- [Middlewares](#conditional-middlewares)
    * [Create](#create)
    * [Conditional Middlewares](#conditional-middlewares)
- [Features To Be Added Soon](#features-to-be-added-soon)

## Setup
You can start using this repository by cloning it.

## Architecture


## Di Container

### Providers

### Di Scopes

### Services

### Defs



## Database

### Supported Databases

supports databases MySQL and PostgreSQL

### Procedures

Creating a file in the models/procedures folder,
create a type and create(db *gorm.DB), drop(db *gorm.DB), dropIfExist(db *gorm.DB) methods for that type.
create getter function for this procedure

You can look at the example below

#### Example

models/procedures/getUsersCount.go
```go
type UserCount struct {
    Count int  `json:"rate"`
}

func (UserCount) create(db *gorm.DB) error {
    sql := `CREATE PROCEDURE GetUsersCount()
    BEGIN
      SELECT COUNT(*) as count FROM users;
    END`
    
    return db.Exec(sql).Error
}

func (UserCount) drop(db *gorm.DB) error {
    sql := `DROP PROCEDURE GetUserCount;`
    return db.Exec(sql).Error
}

func (UserCount) dropIfExist(db *gorm.DB) error {
    sql := `DROP PROCEDURE IF EXISTS GetUserCount;`
    return db.Exec(sql).Error
}

func GetUserCount(db *gorm.DB) UserCount {
    var returnVal UserCount
    db.Raw("CALL GetUserCount()").Scan(&returnVal)
    return returnVal
}
```

#### Register Procedure

models/procedures/base.go

```go
func Initialize() {
    db := app.Application.Container.UnscopedGetDb()

    // UserCount Register
    _ = DropProcedureIfExist(UserCount{}, db)
    _ = CreateProcedure(UserCount{}, db)


    app.Application.Container.Clean()
}
```

### Db Scopes

#### Pagination Scope

In Controller Usage

controllers/userController.go index method
```go
request := new(requests.Pagination)

if err = c.Bind(request); err != nil {
   return
}

var count int64
dic.Db(c.Request()).Model(&models.User{}).Count(&count)

var users []models.User

if err := dic.Db(c.Request()).Scopes(scopes.Paginate(request, models.User{}, "name")).Find(&users).Error; err != nil {
   return echo.ErrInternalServerError
}

return c.JSON(http.StatusOK, helpers.SuccessResponse(accessories.Paginator{
    TotalRecord: int(count),
    Records:     users,
    Limit:       request.Limit,
    Page:        request.Page,
}))
```

You can add pagination to any request object
```go
type ExampleRequest struct {
    validation.Validatable `json:"-" form:"-" query:"-"`

    /**
     * PAGINATION
     */
    Pagination Pagination
  
    /**
    * BODY
    */
    Verified int `json:"verified" form:"verified" query:"verified"`
}
```

In Controller Usage
```go
 if err := dic.Db(c.Request()).Scopes(scopes.Paginate(&request.Pagination, models.User{}, "name")).Find(&users).Error; err != nil {
     return echo.ErrInternalServerError
 }
```

### Migrations

When you create a model, insert it into the Initialize() function of the database/migration/base.go.

#### Register Migration

models/procedures/base.go

#### Example
```go
func Initialize() {
    db := app.Application.Container.UnscopedGetDb()

    _ = db.AutoMigrate(&models.User{})
    
    app.Application.Container.Clean()
}
```

## ORM
Check out fantastic gorm library https://gorm.io/docs/


## Requests

### Create New Requests

Creating a file in the requests folder,
create a type and create a Validate() method for that type.
You can look at the examples below

### Bind Request And Validate

#### Example

Request Object

```go
type LoginRequest struct {
    validation.Validatable `json:"-" form:"-" query:"-"`
 
    /**
    * BODY
    */
     Email    string `json:"email" form:"email" query:"email"`
     Password string `json:"password" form:"password" query:"password"`
}

func (r LoginRequest) Validate() error {
    return validation.ValidateStruct(&r,
        validation.Field(&r.Email, validation.Required, validation.Length(4, 50), is.Email),
        validation.Field(&r.Password, validation.Required, validation.Length(8, 50)),
    )
}
```

In Controller Usage

```go
request := new(requests.LoginRequest)

if err = c.Bind(request); err != nil {
	return
}

v := request.Validate()

if v != nil {
    return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
         "errors": v,
    })
}
	
// you can access binded request object
fmt.println(request.Email)
```

### More Info For Validations Rules

Check out ozzo-validation library https://github.com/go-ozzo/ozzo-validation

### Custom Rules

#### Examples

rules/stringEquals.go

```go
func StringEquals(str string) validation.RuleFunc {
    return func(value interface{}) error {
        s, _ := value.(string)
            
        if s != str {
             return errors.New("unexpected string")
        }
            
        return nil
    }
}
```

#### Usage In Any Request Object Validate Method

```go
func (r ExampleRequest) Validate() error {
    return validation.ValidateStruct(&r,
        validation.Field(&r.Name, validation.By(rules.StringEquals("john"))),
    )
}
```

## Routes


## Controllers

## Jobs







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