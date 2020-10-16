[comment]: <> (This is a generated file please edit source in ./templates)
[comment]: <> (All modification will be lost, you have been warned)
[comment]: <> ()
### Sample CRUD API for the mysql database default:default@/laravel

## Example
The project is a RESTful api for accessing the mysql database default:default@/laravel.

## Project Files
The generated project will contain the following code under the `./example` directory.
* Makefile
  * useful Makefile for installing tools building project etc. Issue `make` to display help
* .gitignore
  * git ignore for go project
* go.mod
  * go module setup, pass `--module` flag for setting the project module default `example.com/example`
* README.md
  * Project readme
* app/server/main.go
  * Sample Gin Server, with swagger init and comments
* api/*.go
  * REST crud controllers
* dao/*.go
  * DAO functions providing CRUD access to database
* model/*.go
  * Structs representing a row for each database table

The REST api server utilizes the Gin framework, GORM db api and Swag for providing swagger documentation
* [Gin](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo/swag)
* [Gorm](https://github.com/jinzhu/gorm)

## Building
```.bash
make example
```
Will create a binary `./bin/example`

## Running
```.bash
./bin/example
```
This will launch the web server on localhost:8080

## Swagger
The swagger web ui contains the documentation for the http server, it also provides an interactive interface to exercise the api and view results.
http://localhost:8080/swagger/index.html

## REST urls for fetching data


* http://localhost:8080/addresstypes
* http://localhost:8080/addresses
* http://localhost:8080/apikeytypes
* http://localhost:8080/apikeys
* http://localhost:8080/apps
* http://localhost:8080/blocktypes
* http://localhost:8080/blocks
* http://localhost:8080/campaings
* http://localhost:8080/clients
* http://localhost:8080/clientsapps
* http://localhost:8080/clientsusers
* http://localhost:8080/coupons
* http://localhost:8080/customertracking
* http://localhost:8080/failedjobs
* http://localhost:8080/jobs
* http://localhost:8080/larametricslogs
* http://localhost:8080/larametricsmodels
* http://localhost:8080/larametricsnotifications
* http://localhost:8080/larametricsrequests
* http://localhost:8080/migrations
* http://localhost:8080/orderitems
* http://localhost:8080/orders
* http://localhost:8080/pagetaxonomy
* http://localhost:8080/pagetypes
* http://localhost:8080/pages
* http://localhost:8080/passwordresets
* http://localhost:8080/paymentstatus
* http://localhost:8080/paymentstatushistory
* http://localhost:8080/people
* http://localhost:8080/peopletaxonomy
* http://localhost:8080/personactions
* http://localhost:8080/personhistory
* http://localhost:8080/prices
* http://localhost:8080/productimages
* http://localhost:8080/producttaxonomy
* http://localhost:8080/producttypes
* http://localhost:8080/products
* http://localhost:8080/progress
* http://localhost:8080/settings
* http://localhost:8080/shippingstatus
* http://localhost:8080/shippingstatushistory
* http://localhost:8080/tasks
* http://localhost:8080/termtypes
* http://localhost:8080/terms
* http://localhost:8080/users
* http://localhost:8080/usersinapps

## Project Generated Details
```.bash
gen \
    [-v] \
    --sqltype=mysql \
    --connstr \
    default:default@/laravel \
    --database \
    laravel \
    --module=github.com/pedrokoblitz/comm \
    --json \
    --gorm \
    --guregu \
    --generate-dao \
    --generate-proj
```











