![gotham](https://user-images.githubusercontent.com/39353278/103892416-99f6c880-50fc-11eb-8869-af197ca81fd1.png)

I have designed go-gotham boilerplate adhering to SOLID principles for developers to help them create RESTful API. I take advantage of some other libraries however, I did not neglect to add my codes into it. The aim of this boilerplate is to provide developers a common functionality which will make writing APIs efficient, convenient and testable.

## Check out the documentation of supporting projects 

- Di ( https://github.com/sarulabs/di )
- Echo ( https://echo.labstack.com )
- Gorm (https://gorm.io)
- Ozzo-Validation (https://github.com/go-ozzo/ozzo-validation)
- Echo Swagger (https://github.com/swaggo/echo-swagger)
- GoCron (https://github.com/jasonlvhit/gocron)


## Install

You can start using this repository by cloning it.

```
git clone https://github.com/tolgaOzen/go-gotham
```
```
  FOLDER STRUCTURE
  /
  |- app
    |- container
    |- defs
    |- provider
    app.go
  |- config
  |- controllers
  |- database
    |- migrations
    |-seeds
  |- docs
  |- helpers
  |- infrastructures
  |- mails
  |- middlewares
  |- models
  |- policies
  |- repositories
  |- requests
  |- routers
  |- rules
  |- services
  |- utils
  |- viewModels
  |- views - (for mails)
  main.go
  .env
```

  ### App - DI Container
  The container part is the part that all of our objects are injected through interfaces, as we specified in definitions.
  
  ### Controllers
  Controllers are the handlers of all requests coming to the route.
  The controller can implement interfaces to many services to meet the needs of the request. The controller must be unaware of the logic of the services.

  ### Middlewares
  Middlewares work before the request reaches the controller. These are the parts where you can perform authorization check, record requests, limit the number of requests etc. Middlewares can implement service interfaces. This way, they can check the data.

  ### Policies
  Policies folder consists of sections that check if the authorized user is eligible to perform this action.

  ### Services
  Services folder is where the business logic is based. It is responsible for processing the request from the controller. It takes data from the data layer (repositories) and works to meet what the controller expects.

  ### Repositories
  Repositories folder is the data access layer. All database queries made must be performed in the repositories.

  ### Models
  Models folder hosts all structs under models namespace, model is a struct reflecting our data object from / to database. models should only define data structs, no other functionalities should be included here.

  ### ViewModels
  ViewModels folder hosts all the structs under viewmodels namespace, viewmodels are model to be use as a response return of REST API call
  
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