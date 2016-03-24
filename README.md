# Iris Web Framework
<img align="right" width="132" src="http://kataras.github.io/iris/assets/56e4b048f1ee49764ddd78fe_iris_favicon.ico">
[![Build Status](https://travis-ci.org/kataras/iris.svg?branch=development)](https://travis-ci.org/kataras/iris)
[![Go Report Card](https://goreportcard.com/badge/github.com/kataras/iris)](https://goreportcard.com/report/github.com/kataras/iris)
[![GoDoc](https://godoc.org/github.com/kataras/iris?status.svg)](https://godoc.org/github.com/kataras/iris)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kataras/iris?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Iris is a very minimal but flexible go web framework  providing arobust set of features for building shiny web applications.


### V0.0.1
>This project is under extreme development


![Benchmark Wizzard Iris vs gin vs martini](http://kataras.github.io/iris/assets/benchmarks_all.png)

[See all benchmarks](https://github.com/kataras/go-http-routing-benchmark)

```go
package main

import "github.com/kataras/iris"
	iris.Get("/hello", func(c *iris.Context) {
		c.HTML("<b> Hello </b>")
	})
	iris.Listen(":8080")
}

```

## Table of Contents

- [Install](#install)
- [Benchmark](#benchmark)
- [Principles](#principles-of-iris)
- [Features](#features)
- [Introduction](#introduction)
- [API](#api)
- [Declaration & Options](#declaration)
- [Party](#party)
- [Named Parameters](#named-parameters)
- [Match anything and the Static serve handler](#match-anything-and-the-static-serve-handler)
- [Declaring routes](#declaring-routes)
- [Context](#context)
- [Examples](https://github.com/kataras/iris/tree/examples)
- [Third Party Middleware](#third-party-middleware)
- [Contributors](#contributors)
- [Community](#community)
- [Todo](#todo)

## Install

```sh
$ go get github.com/kataras/iris
```

### Update
Iris is still in development status, in order to have the latest version update the package every 2-3 days
```sh
$ go get -u github.com/kataras/iris
```

## Benchmark

With Intel(R) Core(TM) i7-4710HQ CPU @ 2.50GHz 2.50 HGz and 8GB Ram:

[See all benchmarks](https://github.com/kataras/go-http-routing-benchmark)

1. Total Operations (Executions)
2. Nanoseconds per operation ns/op
3. Heap Memory B/op
4. Allocations per operation allocs/op



Benchmark name 					| Total Operations 		| ns/op 		| B/op 		| allocs/op
--------------------------------|----------:|----------:|----------:|------:
BenchmarkAce_GithubAll 			| 10000 	| 121206 	| 13792 	| 167
BenchmarkBear_GithubAll 		| 10000 	| 348919 	| 86448 	| 943
BenchmarkBeego_GithubAll 		| 5000 		| 296816 	| 16608 	| 524
BenchmarkBone_GithubAll 		| 500 		| 2502143 	| 548736 	| 7241
BenchmarkDenco_GithubAll 		| 20000 	| 99705 	| 20224 	| 167
BenchmarkEcho_GithubAll 		| 30000 	| 45469 	| 0 		| 0
BenchmarkGin_GithubAll 		| 50000     | 39402     | 0 	    | 0
BenchmarkGocraftWeb_GithubAll 	| 5000 		| 446025 	| 131656 	| 1686
BenchmarkGoji_GithubAll 		| 2000 		| 547698 	| 56112 	| 334
BenchmarkGojiv2_GithubAll 		| 2000 		| 763043 	| 118864 	| 3103
BenchmarkGoJsonRest_GithubAll 	| 5000 		| 538030 	| 134371 	| 2737
BenchmarkGoRestful_GithubAll 	| 100 		| 14870850 	| 837832 	| 6913
BenchmarkGorillaMux_GithubAll 	| 200 		| 6690383 	| 144464 	| 1588
BenchmarkHttpRouter_GithubAll 	| 20000 	| 65653 	| 13792 	| 167
BenchmarkHttpTreeMux_GithubAll 	| 10000 	| 215312 	| 65856 	| 671
**BenchmarkIris_GithubAll** 	| **100000** 	| **20731** 	| **0** 	| **0**
BenchmarkKocha_GithubAll 		| 10000 	| 167209 	| 23304 	| 843
BenchmarkLARS_GithubAll 		| 30000 		| 41069 	| 0 	| 0
BenchmarkMacaron_GithubAll 		| 2000 	| 665038 	| 201138 	| 1803
BenchmarkMartini_GithubAll 		| 100 		| 5433644 	| 228213 	| 2483
BenchmarkPat_GithubAll 			| 300 		| 4210240 	| 1499569 	| 27435
BenchmarkPossum_GithubAll 		| 10000 	| 255114 	| 84448 	| 609
BenchmarkR2router_GithubAll 	| 10000 	| 237113 	| 77328 	| 979
BenchmarkRevel_GithubAll 		| 2000 		| 1150565 	| 337424 	| 5512
BenchmarkRivet_GithubAll 		| 20000 	| 96555 	| 16272 	| 167
BenchmarkTango_GithubAll 		| 5000 		| 417423 	| 87075 	| 2267
BenchmarkTigerTonic_GithubAll 	| 2000 		| 994556 	| 233680 	| 5035
BenchmarkTraffic_GithubAll 		| 200 		| 7770444 	| 2659331 	| 21848
BenchmarkVulcan_GithubAll 		| 5000 		| 292216 	| 19894 	| 609


## Principles of iris

- Easy to use

- Robust

- Simplicity Equals Productivity. The best way to make something seem simple is to have it actually be simple. iris's main functionality has clean, classically beautiful APIs


## Features

**Parameters in your routing pattern:** give meaming to your routes, give them a path segment a name and the iris' will provide the dynamic value to you.

**Party of routes:** Combine routes where have same prefix, provide a middleware to this Party, a Party can have other Party too.

**Compatible:** At the end the Iris is just a middleware which acts like router and a small simply web framework, this means that you can you use it side-by-side with your favorite big and well-tested web framework. Iris is fully compatible with the **net/http package.**

**Multi server instances:** Besides the fact that iris has a default main server. You can declare a new iris using the iris.New() func. example: server1:= iris.New(); server1.Get(....); server1.Listen(":9999")


## Introduction
The name of this framework came from **Greek mythology**, **Iris** was the name of the Greek goddess of the **rainbow**.
Iris is a very minimal but flexible golang http middleware & standalone web application framework, providing a robust set of features for building single & multi-page, web applications.

```go
package main

import "github.com/kataras/iris"
	iris.Get("/hello", func(c *iris.Context) {
		c.HTML("<b> Hello </b>")
	})
	iris.Listen(":8080")
}

```
>Note: for macOS, If you are having problems on .Listen then pass only the port "8080" without ':'



## API
**Use of GET,  POST,  PUT,  DELETE, HEAD, PATCH & OPTIONS**

```go
package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	iris.Get("/home", iris.ToHandlerFunc(testGet))
	iris.Post("/login",testPost)
	iris.Put("/add",testPut)
	iris.Delete("/remove",testDelete)
	iris.Head("/testHead",testHead)
	iris.Patch("/testPatch",testPatch)
	iris.Options("/testOptions",testOptions)
	iris.Listen(":8080")
}

func testGet(res http.ResponseWriter, req *http.Request) {
	//...
}

//iris.Context gives more information and control of the route, named parameters, redirect, error handling and render.
func testPost(c *iris.Context) {
	//...
}

//and so on....
```
> Iris is compatible with net/http package over iris.ToHandlerFunc(...) or iris.ToHandler(...) if you wanna use a whole iris.Handler interface. You can use any method you like but, believe me it's easier to pass just a func(c *Context).

## Declaration

Let's make a pause,

- Q: Why you use iris package declaration? other frameworks needs more lines to start a server
- A: Iris gives you the freedom to choose between three methods/ways to use Iris

 1. global **iris.**
 2. set a new iris with variable  = iris**.New()**
 3. set a new iris with custom options with variable = iris**.Custom(options)**


```go
import "github.com/kataras/iris"

// 1.
func methodFirst() {

	iris.Get("/home",func(c *iris.Context){})
	iris.Listen(":8080")
	//iris.ListenTLS(":8080","yourcertfile.cert","yourkeyfile.key"
}
// 2.
func methodSecond() {

	api := iris.New()
	api.Get("/home",func(c *iris.Context){})
	api.Listen(":8080")
}
// 3.
func methodThree() {
	//these are the default options' values
	options := iris.StationOptions{
		Profile:            false,
		ProfilePath:        iris.DefaultProfilePath,
		Cache:              true,
		CacheMaxItems:      0,
		CacheResetDuration: 5 * time.Minute,
		PathCorrection: 	true, //explanation at the end of this chapter
	}//these are the default values that you can change
	//DefaultProfilePath = "/debug/pprof"

	api := iris.Custom(options)
	api.Get("/home",func(c *iris.Context){})
	api.Listen(":8080")
}

```

> Note that with 2. & 3. you **can define and use more than one Iris container** in the
> same app, when it's necessary.

As you can see there are some options that you can chage at your iris declaration, you cannot change them after.
**If an option value not passed then it considers to be false if bool or the default if string.**

For example if we do that...
```go
import "github.com/kataras/iris"
func main() {
	options := iris.StationOptions{
		Cache:				true,
		Profile:            true,
		ProfilePath:        "/mypath/debug",
	}

	api := iris.Custom(options)
	api.Listen(":8080")
}
```
run it, then you can open your browser, type '**localhost:8080/mypath/debug/profile**' at the location input field and you should see a webpage  shows you informations about CPU.

For profiling & debug there are seven (7) generated pages ('/debug/pprof/' is the default profile path, which on previous example we changed it to '/mypath/debug'):

 1. /debug/pprof/cmdline
 2. /debug/pprof/profile
 3. /debug/pprof/symbol
 4. /debug/pprof/goroutine
 5. /debug/pprof/heap
 6. /debug/pprof/threadcreate
 7. /debug/pprof/pprof/block


**PathCorrection**
corrects and redirects the requested path to the registed path
for example, if /home/ path is requested but no handler for this Route found,
then the Router checks if /home handler exists, if yes, redirects the client to the correct path /home
and VICE - VERSA if /home/ is registed but /home is requested then it redirects to /home/ (Default is true)

## Party

Let's party with Iris web framework!

```go
func main() {

    //log everything middleware

    iris.UseFunc(func(c *iris.Context) {
		println("[Global log] the requested url path is: ", c.Request.URL.Path)
		c.Next()
	})

    // manage all /users
    users := iris.Party("/users")
    {
  	    // provide a  middleware
		users.UseFunc(func(c *iris.Context) {
			println("LOG [/users...] This is the middleware for: ", c.Request.URL.Path)
			c.Next()
		})
		users.Post("/login", loginHandler)
        users.Get("/:userId", singleUserHandler)
        users.Delete("/:userId", userAccountRemoveUserHandler)
    }



    // Party inside an existing Party example:

    beta:= iris.Party("/beta")

    admin := beta.Party("/admin")
    {
		/// GET: /beta/admin/
		admin.Get("/", func(c *iris.Context){})
		/// POST: /beta/admin/signin
        admin.Post("/signin", func(c *iris.Context){})
		/// GET: /beta/admin/dashboard
        admin.Get("/dashboard", func(c *iris.Context){})
		/// PUT: /beta/admin/users/add
        admin.Put("/users/add", func(c *iris.Context){})
    }



    iris.Listen(":8080")
}
```


## Named Parameters

Named parameters are just custom paths to your routes, you can access them for each request using context's **c.Param("nameoftheparameter")**. Get all, as array (**{Key,Value}**) using **c.Params** property.

No limit on how long a path can be.

Usage:


```go
package main

import "github.com/kataras/iris"

func main() {
	// MATCH to /hello/anywordhere  (if PathCorrection:true match also /hello/anywordhere/)
	// NOT match to /hello or /hello/ or /hello/anywordhere/something
	iris.Get("/hello/:name", func(c *iris.Context) {
		name := c.Param("name")
		c.Write("Hello %s", name)
	})

	// MATCH to /profile/iris/friends/42  (if PathCorrection:true matches also /profile/iris/friends/42/ ,otherwise not match)
	// NOT match to /profile/ , /profile/something ,
	// NOT match to /profile/something/friends,  /profile/something/friends ,
	// NOT match to /profile/anything/friends/42/something
	iris.Get("/profile/:fullname/friends/:friendId",
		func(c *iris.Context){
			name:= c.Param("fullname")
			//friendId := c.ParamInt("friendId")
			c.HTML("<b> Hello </b>"+name)
		})

	iris.Listen(":8080")
	//or 
	//log.Fatal(http.ListenAndServe(":8080", iris.Serve()))
}

```

## Match anything and the Static serve handler

Match everything/anything (symbol *withAKeyLikeParameters)
```go
// Will match any request which url's preffix is "/anything/" and has content after that
iris.Get("/anything/*randomName", func(c *iris.Context) { } )
// Match: /anything/whateverhere/whateveragain , /anything/blablabla
// c.Params("randomName") will be /whateverhere/whateveragain, blablabla
// Not Match: /anything , /anything/ , /something
```
Pure http static  file server as handler using **iris.Static("./path/to/the/resources/directory/","path_to_strip_or_nothing")**
```go
// Will match any request which url's preffix is "/public/"
/* and continues with a file whith it's extension which exists inside the os.Gwd()(dot means working directory)+ /static/resources/
*/
iris.Get("/public/*assets", iris.Static("./static/resources/","/public/"))
//Note: strip of the /public/ is handled  by passing the last argument to "/public/"
//you can pass only the first two arguments for no strip path.
```
## Declaring routes
Iris framework has three (3) different forms of functions in order to declare a route's handler and one(1) annotated struct to declare a complete route.


 1. Typical classic handler function, compatible with net/http and other frameworks using iris.ToHandlerFunc
	 *  **func(res http.ResponseWriter, req *http.Request)**
```go
	iris.Get("/user/add", iris.ToHandlerFunc(func(res http.ResponseWriter, req *http.Request)) {

	})
```
 2. Context parameter in function-declaration
	 * **func(c *iris.Context)**

```go
	iris.Get("/user/:userId", func(c *iris.Context) {

	})
```

 3. http.Handler again it can be converted by ToHandlerFunc
	 * **http.Handler**

```go
	iris.Get("/about", iris.ToHandlerFunc(http.HandlerFunc(func(res http.Response, req *req.Request)) {

	}))
```
> Note that all .Get,.Post takes a func(c *Context) as parameter, to pass an iris.Handler use the iris.Handle("/path",handler,"GET")
 4. **'External' annotated struct** which directly implements the Iris Handler interface



```go
///file: userhandler.go
import "github.com/kataras/iris"

type UserRoute struct {
	iris.Handler `get:"/profile/user/:userId"`
}

func (u *UserRoute) Serve(c *iris.Context) {
	defer c.Close()
	userId := c.Param("userId")
	c.RenderFile("user.html", struct{ Message string }{Message: "Hello User with ID: " + userId})
}


///file: main.go

	//...cache the html files
	iris.Templates("src/iristests/templates/**/*.html")
	//...register the handler
	iris.HandleAnnotated(&UserRoute{})
	//...continue writing your wonderful API

```
Personally I use the external struct and the **func(c *iris.Context)** form .
 At the next chapter you will learn what are the benefits of having the  **Context** as parameter to the handler.


## Context

> Variables

 1. **ResponseWriter**
	 - The ResponseWriter is the exactly the same as you used to use with the standar http library.
 2. **Request**
	 - The Request is the pointer of the *Request, is the exactly the same as you used to use with the standar http library.
 3. **Params**
	 - Contains the Named path Parameters, imagine it as a map[string]string which contains all parameters of a request.

>Functions

 1. **Clone()**
	 - Returns a clone of the Context, useful when you want to use the context outscoped for example in goroutines.
 2. **Write(contents string)**
	 - Writes a pure string to the ResponseWriter and sends to the client.
 3. **Param(key string)** returns string
	 - Returns the string representation of the key's  named parameter's value. Registed path= /profile/:name) Requested url is /profile/something where the key argument is the named parameter's key, returns the value  which is 'something' here.
 4. **ParamInt(key string)** returns integer, error
	 - Returns the int representation of the key's  named parameter's value, if something goes wrong the second return value, the error is not nil.
 5. **URLParam(key string)** returns string
	 - Returns the string representation of a requested url parameter (?key=something) where the key argument is the name of, something is the returned value.
 6. **URLParamInt(key string)** returns integer, error
	 - Returns the int representation of  a requested url parameter
 7. **SetCookie(name string, value string)**
	 - Adds a cookie to the request.
 8. **GetCookie(name string)** returns string
	 - Get the cookie value, as string, of a cookie.
 9. **ServeFile(path string)**
	 - This just calls the http.ServeFile, which serves a file given by the path argument  to the client.
 10. **NotFound()**
	 - Sends a http.StatusNotFound with a custom template you defined (if any otherwise the default template is there) to the client.
	 --- *Note: We will learn all about Custom Error Handlers later*.
 11. **Close()**
	 - Calls the Request.Body.Close().

 12. **WriteHTML(status int, contents string) & HTML(contents string)**
	 - WriteHTML: Writes html string with a given http status to the client, it sets the Header with the correct content-type.
	 - HTML: Same as WriteHTML but you don't have to pass a status, it's defaulted to http.StatusOK (200).
 13. **WriteData(status int, binaryData []byte) & Data(binaryData []byte)**
	 - WriteData: Writes binary data with a given http status to the client, it sets the Header with the correct content-type.
	 - Data : Same as WriteData but you don't have to pass a status, it's defaulted to http.StatusOK (200).
 14. **WriteText(status int, contents string) & Text(contents string)**
	 - WriteText: Writes plain text with a given http status to the client, it sets the Header with the correct content-type.
	 - Text: Same as WriteTextbut you don't have to pass a status, it's defaulted to http.StatusOK (200).
 15. **ReadJSON(jsonObject interface{}) error**
 	 - ReadJSON: reads the request's body content and parses it, assigning the result into jsonObject passed by argument. Don't forget to pass the argument as reference.
 16. **WriteJSON(status int, jsonObject interface{}) & JSON(jsonObject interface{}) returns error**
	 - WriteJSON: Writes json which is converted from structed object(s) with a given http status to the client, it sets the Header with the correct content-type. If something goes wrong then it's returned value which is an error type is not nil. No indent.
 17.  **RenderJSON(jsonObjects ...interface{}) returns error**
	 - RenderJSON: Same as WriteJSON & JSON but with Indent (formated json)
	 - JSON: Same as WriteJSON but you don't have to pass a status, it's defaulted to http.StatusOK (200).
 18. **ReadXML(xmlObject interface{}) error**
 	 - ReadXML: reads the request's body and parses it, assigin the result into xmlObject passed by argument.
 19. **WriteXML(status int, xmlStructs ...interface{}) & XML(xmlStructs ...interface{}) returns error**
	 - WriteXML: Writes writes xml which is converted from struct(s) with a given http status to the client, it sets the Header with the correct content-type. If something goes wrong then it's returned value which is an error type is not nil.
	 - XML: Same as WriteXML but you don't have to pass a status, it's defaulted to http.StatusOK (200).
 20. **RenderFile(file string, pageContext interface{}) returns error**
	 - RenderFile: Renders a file by its name (which a file is saved to the template cache) and a page context passed to the function, default http status is http.StatusOK(200) if the template was found, otherwise http.StatusNotFound(404). If something goes wrong then it's returned value which is an error type is not nil.
 21. **Render(pageContext interface{})  returns error**
	 - Render: Renders the root file template and a context passed to the function, default http status is http.StatusOK(200) if the template was found, otherwise http.StatusNotFound(404). If something goes wrong then it's returned value which is an error type is not nil.
--- *Note:  We will learn how to add templates at the next chapters*.

 22. **Next()**
	 - Next: calls all the next handler from the middleware stack, it used inside a middleware

 23. **SendStatus(statusCode int, message string)**
	 - SendStatus:  writes a http statusCode with a text/plain message


**[[TODO chapters: Register custom error handlers, cache templates , create & use middleware]]**

Inside the [examples](https://github.com/kataras/iris/tree/examples) branch you will find practical examples

.


## Third Party Middleware

*The iris tries to supports a lot of middleware out there, you can use them by parsing their handlers, for example: *
```go
iris.UseFunc(func(c *iris.Context) {
		//run the middleware here
		c.Next()
	})
```

>Note: Some of these, may not be work, a lot of them are especially for Negroni and nothing more.

Iris has a middleware system to create it's own middleware and is at a state which tries to find person who are be willing to convert them to Iris middleware or create new. Contact or open an issue if you are interesting.


| Middleware | Author | Description | Tested |
| -----------|--------|-------------|--------|
| [sessions](https://github.com/kataras/iris/tree/development/sessions) | [Ported to Iris](https://github.com/kataras/iris/tree/development/sessions) | Session Management | [Yes](https://github.com/kataras/iris/tree/development/sessions) |
| [Graceful](https://github.com/tylerb/graceful) | [Tyler Bunnell](https://github.com/tylerb) | Graceful HTTP Shutdown | [Yes](https://github.com/kataras/iris/tree/examples/thirdparty_graceful) |
| [gzip](https://github.com/kataras/iris/tree/development/middleware/gzip/) | [Iris](https://github.com/kataras/iris) | GZIP response compression | [Yes](https://github.com/kataras/iris/tree/examples/middleware_compression_gzip) |
| [RestGate](https://github.com/pjebs/restgate) | [Prasanga Siripala](https://github.com/pjebs) | Secure authentication for REST API endpoints | No |
| [secure](https://github.com/unrolled/secure) | [Cory Jacobsen](https://github.com/unrolled) | Middleware that implements a few quick security wins | [Yes](https://github.com/kataras/iris/tree/examples/thirdparty_secure) |
| [JWT Middleware](https://github.com/auth0/go-jwt-middleware) | [Auth0](https://github.com/auth0) | Middleware checks for a JWT on the `Authorization` header on incoming requests and decodes it| No |
| [binding](https://github.com/mholt/binding) | [Matt Holt](https://github.com/mholt) | Data binding from HTTP requests into structs | No |
| [logrus](https://github.com/meatballhat/negroni-logrus) | [Dan Buch](https://github.com/meatballhat) | Logrus-based logger | No |
| [render](https://github.com/unrolled/render) | [Cory Jacobsen](https://github.com/unrolled) | Render JSON, XML and HTML templates | No |
| [gorelic](https://github.com/jingweno/negroni-gorelic) | [Jingwen Owen Ou](https://github.com/jingweno) | New Relic agent for Go runtime | No |
| [pongo2](https://github.com/kataras/iris/middleware/pongo2) | [Iris](https://github.com/kataras/iris) | Middleware for [pongo2 templates](https://github.com/flosch/pongo2)| [Yes](https://github.com/kataras/iris/middleware/pongo2/README.md) |
| [oauth2](https://github.com/goincremental/negroni-oauth2) | [David Bochenski](https://github.com/bochenski) | oAuth2 middleware | No |
| [permissions2](https://github.com/xyproto/permissions2) | [Alexander Rødseth](https://github.com/xyproto) | Cookies, users and permissions | No |
| [onthefly](https://github.com/xyproto/onthefly) | [Alexander Rødseth](https://github.com/xyproto) | Generate TinySVG, HTML and CSS on the fly | No |
| [cors](https://github.com/kataras/iris/tree/development/middleware/cors) | [Keuller Magalhaes](https://github.com/keuller) | [Cross Origin Resource Sharing](http://www.w3.org/TR/cors/) (CORS) support | [Yes](https://github.com/kataras/iris/tree/development/middleware/cors) |
| [xrequestid](https://github.com/pilu/xrequestid) | [Andrea Franz](https://github.com/pilu) | Middleware that assigns a random X-Request-Id header to each request | No |
| [VanGoH](https://github.com/auroratechnologies/vangoh) | [Taylor Wrobel](https://github.com/twrobel3) | Configurable [AWS-Style](http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html) HMAC authentication middleware | No |
| [stats](https://github.com/thoas/stats) | [Florent Messa](https://github.com/thoas) | Store information about your web application (response time, etc.) | No |

## Contributors

Thanks goes to the people who have contributed code to this package, see the
[GitHub Contributors page][].

[GitHub Contributors page]: https://github.com/kataras/iris/graphs/contributors



## Community

If you'd like to discuss this package, or ask questions about it, feel free to

* **Chat**: https://gitter.im/kataras/iris

## Guidelines
- Never stop writing the docs.
- Provide full README for examples branch and thirdparty middleware examples.
- Before any commit run -count 50 -benchtime 30s , if performance stays on top then commit else find other way to do the same thing.
- Notice the author of any thirdparty package before I try to port into iris myself, maybe they can do it better.
- Notice author's of middleware, which I'm writing examples for,to take a look, if they don't want to exists in the Iris community, I have to respect them.

## Todo
- [x] [Provide a lighter, with less using bytes,  to save middleware for a route.](https://github.com/kataras/iris/tree/development/handler.go)
- [x] [Create examples.](https://github.com/kataras/iris/tree/examples)
- [x] [Subdomains supports with the same syntax as iris.Get, iris.Post ...](https://github.com/kataras/iris/tree/examples/subdomains_simple)
- [x] [Provide a more detailed benchmark table to the README with all go web frameworks that I know, no just the 6 most famous](https://github.com/kataras/iris/tree/benchmark)
- [ ] Convert useful middlewares out there into Iris middlewares, or contact with their authors to do so.
- [ ] Create an easy websocket api.
- [ ] Create a mechanism that scan for Typescript files, compile them on server startup and serve them.

## Licence

This project is licensed under the [BSD 3-Clause License](https://opensource.org/licenses/BSD-3-Clause).
License can be found [here](https://github.com/kataras/iris/blob/master/LICENSE). 


