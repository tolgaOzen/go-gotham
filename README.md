![gotham](https://user-images.githubusercontent.com/39353278/103892416-99f6c880-50fc-11eb-8869-af197ca81fd1.png)

I have designed go-gotham boilerplate adhering to SOLID principles for developers to help them create RESTful API. I take advantage of some other libraries however, I did not neglect to add my codes into it. The aim of this boilerplate is to provide developers a common functionality which will make writing APIs efficient, convenient and testable.

## Check out the documentation of supporting projects 

- Di ( https://github.com/sarulabs/di )
- Echo ( https://echo.labstack.com )
- Gorm (https://gorm.io)
- Ozzo-Validation (https://github.com/go-ozzo/ozzo-validation)
- GoCron (https://github.com/jasonlvhit/gocron)

## Table of contents

![GitHub](https://img.shields.io/github/license/tolgaOzen/go-gotham)
![GitHub top language](https://img.shields.io/github/languages/top/tolgaozen/go-gotham)
![GitHub last commit](https://img.shields.io/github/last-commit/tolgaozen/go-gotham)

- [Install](#install)
- [Controllers](#controllers)
- [Services](#services)
- [Repositories](#repositories)
- [Middlewares](#conditional-middlewares):
  * [Conditional Middlewares](#conditional-middlewares)
- [Definitions](#definitions)
- [Provider](#provider)
- [Container](#container)
- [Models](#models)
- [ViewModels](#viewModels)
- [Routers](#routers)
- [Database](#database)
    * [Migrations](#migrations)
    * [Db Scopes](#db-scopes)
    * [Procedures](#procedures)
- [ORM](#orm)
- [Requests](#requests)
    * [Create Requests](#create-new-requests)
    * [Bind Request And Validate](#bind-request-and-validate)
    * [More Info For Validations Rules](#more-info-for-validations-rules)
    * [Custom Rules](#custom-rules)
- [Auth](#auth)
    * [JWT](#jwt)
- [Jobs](#jobs)
- [Features To Be Added Soon](#features-to-be-added-soon)

## Install

You can start using this repository by cloning it.

```
git clone https://github.com/tolgaOzen/go-gotham
```

## Controllers

Controllers are the handlers of all requests coming to the route.

The controller can implement interfaces to many services to meet the needs of the request. The controller must be unaware of the logic of the services.

#### Example

controllers/userController.go
```go
type UserController struct {
    UserService services.IUserService
}

/**
 * index
 *
 * @param echo.Context
 * @return error
 */
func (u UserController) Index(c echo.Context) (err error) {

    request := new(scopes.Pagination)

    if err = c.Bind(request); err != nil {
        return
    }

    users, err := u.UserService.GetUsers(request, "name")
    if err != nil {
        return echo.ErrInternalServerError
    }

    count, err := u.UserService.GetUsersCount()
    if err != nil {
        return echo.ErrInternalServerError
    }

    return c.JSON(http.StatusOK, helpers.SuccessResponse(viewModels.Paginator{
        TotalRecord: int(count),
        Records:     users,
        Limit:       request.Limit,
        Page:        request.Page,
    }))
}
```
#### Route Example

routes/api.go
```go
r.GET("/users", app.Application.Container.GetUserController().Index,GMiddleware.And([]GMiddleware.IMiddleware{app.Application.Container.GetIsAdminMiddleware(),app.Application.Container.GetIsVerifiedMiddleware()}))
```

#### Injection
The service interface dependency of the controllers is injected in the defs folder.

defs/controllers.go
```go
var ControllerDefs = []dingo.Def{
    {
        Name:  "user-controller",
        Scope: di.App,
        Build: func(service services.IUserService) (controllers.UserController, error) {
            return controllers.UserController{
        UserService: service,
        }, nil
    },
        Params: dingo.Params{
        "0": dingo.Service("user-service"),
        },
    },
    .
    .
    .
}
```

## Services
The services folder is where the business logic is based. It is responsible for processing the request from the controller. It takes data from the data layer (repositories) and works to meet what the controller expects.

### Example

services/userService.go
```go
type IUserService interface {
    GetUsers(pagination *scopes.Pagination, orderDefault string) ([]models.User, error)
    GetUserByID(id int) (models.User, error)
    GetUserByEmail(email string) (models.User, error)
    GetUsersCount() (int64, error)
}

type UserService struct {
    UserRepository repositories.IUserRepository
}

func (service *UserService) GetUserByID(id int) (user models.User, err error) {
    return service.UserRepository.GetUserByID(id)
}

func (service *UserService) GetUserByEmail(email string) (user models.User, err error) {
    return service.UserRepository.GetUserByEmail(email)
}

func (service *UserService) GetUsers(pagination *scopes.Pagination, orderDefault string) (users []models.User, err error) {
    return service.UserRepository.GetUsers(pagination, orderDefault)
}

func (service *UserService) GetUsersCount() (count int64, err error) {
    return service.UserRepository.GetUsersCount()
}
```


### Injection
The data layer interface dependency of the services is injected in the defs folder.

defs/userService.go
```go
var UserServiceDefs = []dingo.Def{
    .
    .
    .
    {
        Name:  "user-service",
        Scope: di.App,
        Build: func(repository repositories.IUserRepository) (s services.IUserService , err error) {
            return &services.UserService{UserRepository: repository}, nil
        },
        Params: dingo.Params{
            "0": dingo.Service("user-repository"),
        },
    },  
}
```

## Repositories

The repositories folder is the data access layer. All database queries made must be performed in the repositories.

### Examples

repositories/userRepository.go
```go
type IUserRepository interface {
	GetUserByID(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUsers(pagination *scopes.Pagination, orderDefault string) ([]models.User, error)
	GetUsersCount() (int64, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (repository *UserRepository) GetUserByID(id int) (user models.User, err error) {
	err = repository.DB.First(&user, id).Error
	return
}

func (repository *UserRepository) GetUserByEmail(email string) (user models.User, err error) {
	err = repository.DB.Where("email = ?", email).First(&user).Error
	return
}

func (repository *UserRepository) GetUsers(pagination *scopes.Pagination, orderDefault string) (users []models.User, err error) {
	err = repository.DB.Scopes(pagination.Paginate(models.User{} , orderDefault)).Find(&users).Error
	return
}

func (repository *UserRepository) GetUsersCount() (count int64, err error) {
	// you can user getUsersCount procedure here
	err = repository.DB.Model(&models.User{}).Count(&count).Error
	return
}
```

### Injection
The database dependency of the repositories is injected into the defs folder.

defs/userService.go
```go
var UserServiceDefs = []dingo.Def{
    {
        Name:  "user-repository",
        Scope: di.App,
        Build: func(db *gorm.DB) (s repositories.IUserRepository, err error) {
          return &repositories.UserRepository{DB: db}, nil
        },
        Params: dingo.Params{
          "0": dingo.Service("db"),
       },
    },
    . 
    .
    .
}
```

## Middlewares

Middlewares work before the request reaches the controller.
These are the parts where you can perform authorization check, record requests, limit the number of requests, etc.
Middlewares can implement service interfaces. This way, they can check the data.

#### Injection
The service interface dependency of the middlewares is injected in the defs folder.

defs/middlewares.go
```go
var MiddlewareDefs = []dingo.Def{
    {
        Name:  "is-admin-middleware",
        Scope: di.App,
        Build: func(repository services.IUserService) (s GMiddleware.IsAdmin, err error) {
            return GMiddleware.IsAdmin{UserService: repository}, nil
        },
        Params: dingo.Params{
        "0": dingo.Service("user-service"),
        },
    },
    {
        Name:  "is-verified-middleware",
        Scope: di.App,
        Build: func(repository services.IUserService) (s GMiddleware.IsVerified , err error) {
            return GMiddleware.IsVerified{UserService: repository}, nil
        },
        Params: dingo.Params{
        "0": dingo.Service("user-service"),
        },
    },
}
```

### Conditional Middlewares
The purpose of the conditional middlewares is to decrease the redundant code.
If we want authenticated user to be admin or verified user we supposed to have written a code with middleware such as isAdminOrIsVerified. In another scenario, we could have wanted authenticated user to be an admin and verified. For this reason we should have written isAdminAndVerified middleware.
If we only write the isAdmin middleware and isVerified middleware, we will reduce the code repetition in all scenarios.
You can take a look at the example below.


#### Example

middlewares/isAdmin.go

```go
type IsAdmin struct {
    services.IUserService
}

func (i IsAdmin) control(c echo.Context) (bool bool, err error) {
    u := c.Get("user").(*jwt.Token)
    claims := u.Claims.(*config.JwtCustomClaims)

    user, err := i.FirstUserByID(int(claims.Id))

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, echo.ErrUnauthorized
        }
        return false, echo.ErrInternalServerError
    }

    if user.IsAdmin() {
        return true, nil
    }

    return false, errors.New("you are not admin")
}

```

middlewares/isVerified.go
```go
type IsVerified struct{
    services.IUserService
}

func (i IsVerified) control(c echo.Context) (bool bool, err error) {
    u := c.Get("user").(*jwt.Token)
    claims := u.Claims.(*config.JwtCustomClaims)

    user, err := i.FirstUserByID(int(claims.Id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return false, echo.ErrUnauthorized
        }
        return false, echo.ErrInternalServerError
    }

    if user.IsVerified() {
        return true, nil
    }

    return false, errors.New("your email not verified")
}

```

#### OR

```go
r.GET("/users/:user", app.Application.Container.GetUserController().Show,GMiddleware.Or([]GMiddleware.IMiddleware{app.Application.Container.GetIsAdminMiddleware(),app.Application.Container.GetIsVerifiedMiddleware()}))
```
Authenticated user must be admin or verified

#### AND

```go
r.GET("/users", app.Application.Container.GetUserController().Index,GMiddleware.And([]GMiddleware.IMiddleware{app.Application.Container.GetIsAdminMiddleware(),app.Application.Container.GetIsVerifiedMiddleware()}))
```
Authenticated user must be admin and verified

## Definitions
The definition consists of parts where we write the dependencies required to create the object and where we can determine the life cycles of objects.

#### Examples

app/defs/database.go
```go
var DatabaseServiceDefs = []dingo.Def{
	{
		Name:  "db-pool",
		Scope: di.App,
		Build: func() (gorm.Dialector, error) {
			return services.NewDatabaseService(config.GetDbConfig()).OpenDatabase(), nil
		},
	},
	{
		Name:  "db",
		Scope: di.App,
		Build: func(dia gorm.Dialector) (db *gorm.DB,err error) {
			return services.DatabaseService{}.ConnectDatabase(dia)
		},
		Params: dingo.Params{
			"0": dingo.Service("db-pool"),
		},
		Close: func(db *gorm.DB) error {
			sqlDB, _ := db.DB()
			return sqlDB.Close()
		},
	},
}
```
Like the example above, the db object is dependent on the dp-pool object. While calling the db object, the db-pool object is injected into the db object, and the  db object is created.

app/defs/controllers.go
```go
var ControllerDefs = []dingo.Def{
	{
		Name:  "user-controller",
		Scope: di.App,
		Build: func(service services.IUserService) (controllers.UserController, error) {
			return controllers.UserController{
				IUserService: service,
			}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-service"),
		},
	},
	{
		Name:  "auth-controller",
		Scope: di.App,
		Build: func(service services.IAuthService) (controllers.AuthController, error) {
			return controllers.AuthController{
				IAuthService: service,
			}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("auth-service"),
		},
	},
}
```

app/defs/middlewares.go
```go
var MiddlewareDefs = []dingo.Def{
        {
                Name:  "is-admin-middleware",
                Scope: di.App,
                Build: func(repository services.IUserService) (s GMiddleware.IsAdmin, err error) {
                       return GMiddleware.IsAdmin{IUserService: repository}, nil
                },
                Params: dingo.Params{
                     "0": dingo.Service("user-service"),
                },
        },
        {
                Name:  "is-verified-middleware",
                Scope: di.App,
                Build: func(repository services.IUserService) (s GMiddleware.IsVerified , err error) {
                        return GMiddleware.IsVerified{IUserService: repository}, nil
                },
                Params: dingo.Params{
                        "0": dingo.Service("user-service"),
                },
        },
}
```


app/defs/userService.go
```go
var UserServiceDefs = []dingo.Def{
	{
		Name:  "user-repository",
		Scope: di.App,
		Build: func(db *gorm.DB) (s repositories.IUserRepository, err error) {
			return repositories.UserRepository{DB: db}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("db"),
		},
	},
	{
		Name:  "user-service",
		Scope: di.App,
		Build: func(repository repositories.IUserRepository) (s services.IUserService , err error) {
			return &services.UserService{IUserRepository: repository}, nil
		},
		Params: dingo.Params{
			"0": dingo.Service("user-repository"),
		},
	},
}
```

## Provider

You will have to write the service definitions and register them in a Provider.

app/provider/appServiceProvider.go
```go
func (p *Provider) Load() error {
    if err := p.AddDefSlice(defs.DatabaseServiceDefs); err != nil {
        return err
    }

    if err := p.AddDefSlice(defs.UserServiceDefs); err != nil {
        return err
    }

    if err := p.AddDefSlice(defs.AuthServiceDefs); err != nil {
        return err
    }

    if err := p.AddDefSlice(defs.ControllerDefs); err != nil {
        return err
    }

    if err := p.AddDefSlice(defs.MiddlewareDefs); err != nil {
        return err
    }
    
    return nil
}
```

## Container
The container part is the part that all of our objects are injected through interfaces, as we specified in definations. With the app.Application.Container call, we can access the objects we have created.

check out for more;

- [What is a dependency injection container and why use one ?](https://www.sarulabs.com/post/2/2018-06-12/what-is-a-dependency-injection-container-and-why-use-one.html)

## Models

The models must be a reflection of our data objects. Only a few facilitating methods should be written to models.

#### Example

models/user.go
```go
type User struct {
    ID                uint    `gorm:"primaryKey;auto_increment" json:"id"`
    Name              string  `gorm:"size:255;not null" json:"name"`
    Email             string  `gorm:"size:100;not null;unique;unique_index" json:"email"`
    Password          string  `gorm:"size:100" json:"-"`
    Verified          uint8   `gorm:"type:boolean" json:"verified"`
    VerificationToken *string `gorm:"size:50;" json:"-"`
    Image             *string `gorm:"size:500;" json:"image"`
    Admin             uint8   `gorm:"type:boolean;not null;default:0" json:"admin"`

    // Time
    CreatedAt time.Time      `gorm:"type:datetime(0)" json:"created_at"`
    UpdatedAt time.Time      `gorm:"type:datetime(0)"  json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

/**
 * VerifyPassword
 *
 * @param string , string
 * @return error
 */
func (u *User) VerifyPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

/**
 * IsVerified
 *
 * @return bool
 */
func (u *User) IsVerified() bool {
    return u.Verified == 1
}

/**
 * IsAdmin
 *
 * @return bool
 */
func (u *User) IsAdmin() bool {
    return u.Admin == 1
}

/**
 * Scopes
 *
 * @return *gorm.DB
 */
func (User) VerifiedScope(db *gorm.DB) *gorm.DB {
    return db.Where("users.verified = 1")
}

```

## ViewModels

View models are the model to be used as the response return of the API call

#### Example

viewModels/paginator.go
```go
type Paginator struct {
    TotalRecord int         `json:"total_record"`
    Records     interface{} `json:"records"`
    Limit       int         `json:"limit"`
    Page        int         `json:"page"`
}
```

## Routers
routers/api.go

This is where we connect the appropriate route to process the http request.

Check out echo https://echo.labstack.com/guide

## Database

### Migrations

When you create a model, insert it into the Initialize() function of the database/migration/base.go.

#### Register Migration

models/procedures/base.go

#### Example

```go
func Initialize() {
    db := app.Application.Container.GetDb()
    
    _ = db.AutoMigrate(&models.User{})
}
```

### Db Scopes

#### Pagination Scope

repositories/userRepository.go
```go
func (repository *UserRepository) GetUsers(pagination *scopes.Pagination, orderDefault string) (users []models.User, err error) {
    err = repository.DB.Scopes(pagination.Paginate(models.User{} , orderDefault)).Find(&users).Error
    return
}
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

### Procedures

To create a procedures we need to create a type. We need to add 4 different methodologies;
create(db *gorm.DB),
dropIfExist(db *gorm.DB),
drop(db *gorm.DB) and lastly 
getter method

You can take a look at the example below

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
In order to use the procedure that we have created, we need to initialize the procedure

models/procedures/base.go
```go
func Initialize() {
    db := app.Application.Container.GetDb()

    _ = DropProcedureIfExist(UserCount{}, db)
    _ = CreateProcedure(UserCount{}, db)
}
```

## ORM

Check out fantastic gorm library https://gorm.io/docs/

## Requests

### Create New Requests
In order to create a request we need to create a type first and then, we should add a Validate method for our type.
You can take a look at the example below.

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

#### Example

rules/stringEquals.go
```go
func StringEquals(str string) validation.RuleFunc {
return func (value interface{}) error {
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

## Auth

### JWT

#### Config

config/jwt.go
```go
type JwtCustomClaims struct {
    Id               uint   `json:"id"`
    Name             string `json:"name"`
    Email            string `json:"email"`
    jwt.StandardClaims
}
```

#### Middleware

routers/api.go
```go
r := e.Group("/restricted")

c := middleware.JWTConfig{
	Claims:     &config.JwtCustomClaims{},
    SigningKey: []byte(app.Application.Config.SecretKey),
}

r.Use(middleware.JWTWithConfig(c))
```

#### LoginController

controllers/authController.go
```go
accessTokenExp := time.Now().Add(time.Minute * 15).Unix()

claims := &config.JwtCustomClaims{
    Id:    user.ID,
    Name:  user.Name,
    Email: user.Email,
    StandardClaims: jwt.StandardClaims{
        ExpiresAt: accessTokenExp,
    },
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

accessToken, err := token.SignedString([]byte(config.Conf.SecretKey))

if err != nil {
    return
}

return c.JSON(http.StatusOK, helpers.SuccessResponse(viewModels.Login{
    AccessToken: accessToken,
    AccessTokenExp: accessTokenExp,
    User: user,
}))
```

You can find the information about who owns the token in any controllers or middleware.
```go
u := c.Get("user").(*jwt.Token)
claims := u.Claims.(*config.JwtCustomClaims)
```

## Jobs
Check out GoCron https://github.com/jasonlvhit/gocron

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