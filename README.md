# Iris Web Framework
<img align="right" width="248" src="http://nodets.com/iris_logo.gif">
[![Build Status](https://travis-ci.org/kataras/iris.svg)](https://travis-ci.org/kataras/iris)
[![GoDoc](https://godoc.org/github.com/kataras/iris?status.svg)](https://godoc.org/github.com/kataras/iris)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kataras/iris?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Iris is a very minimal but flexible web framework written in go, providing a robust set of features for building single & multi-page, web applications.

## Table of Contents

- [Install](#install)
- [Principles](#principles-of-iris)
- [Introduction](#introduction)
- [Benchmarks](#benchmarks)
- [Alternatives](#alternatives)
- [Features](#features)
- [API](#api)
- [Named Parameters](#named-parameters)
- [Match anything and the Static serve handler](#match-anything-and-the-static-serve-handler)
- [Declaring routes](#declaring-routes)
- [Context](#context)
- [Renderer](#renderer)
- [Third Party Middleware](#third-party-middleware)
- [Contributors](#contributors)
- [Community](#community)
- [Todo](#todo)

## Install

```sh
$ go get github.com/kataras/iris
```
## Principles of iris

- Easy to use

- Robust

- Simplicity Equals Productivity. The best way to make something seem simple is to have it actually be simple. iris's main functionality has clean, classically beautiful APIs

## Introduction
The name of this framework came from **Greek mythology**, **Iris** was the name of the Greek goddess of the **rainbow**.
Iris is a very minimal but flexible golang http middleware & standalone web application framework, providing a robust set of features for building single & multi-page, web applications.

```go
package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/hello", func(r iris.Renderer) {
		r.HTML("<b> Hello </b>")
	})
	iris.Listen(8080)
}

```

## Benchmarks
Benchmark tests were written by 'the standar' way of benchmarking and comparing performance of other routers and frameworks, see [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/) .

In order to have safe results, this table was taken from [different source](https://raw.githubusercontent.com/gin-gonic/gin/develop/BENCHMARKS.md) than Iris. 

 1. Total Operations
 2. Nanoseconds per Operation (ns/op)  
 3. Heap Memory (B/op)  
 4. Average Allocations per Operation (allocs/op)

Benchmark name 					| 1 		| 2 		| 3 		| 4
--------------------------------|----------:|----------:|----------:|------:
BenchmarkAce_GithubAll 			| 10000 	| 109482 	| 13792 	| 167
BenchmarkBear_GithubAll 		| 10000 	| 287490 	| 79952 	| 943
BenchmarkBeego_GithubAll 		| 3000 		| 562184 	| 146272 	| 2092
BenchmarkBone_GithubAll 		| 500 		| 2578716 	| 648016 	| 8119
BenchmarkDenco_GithubAll 		| 20000 	| 94955 	| 20224 	| 167
BenchmarkEcho_GithubAll 		| 30000 	| 58705 	| 0 		| 0
BenchmarkGin_GithubAll 		| 30000 | 50991| 0 	| 0
BenchmarkGocraftWeb_GithubAll 	| 5000 		| 449648 	| 133280 	| 1889
BenchmarkGoji_GithubAll 		| 2000 		| 689748 	| 56113 	| 334
BenchmarkGoJsonRest_GithubAll 	| 5000 		| 537769 	| 135995 	| 2940
BenchmarkGoRestful_GithubAll 	| 100 		| 18410628 	| 797236 	| 7725
BenchmarkGorillaMux_GithubAll 	| 200 		| 8036360 	| 153137 	| 1791
BenchmarkHttpRouter_GithubAll 	| 20000 	| 63506 	| 13792 	| 167
BenchmarkHttpTreeMux_GithubAll 	| 10000 	| 165927 	| 56112 	| 334
**BenchmarkIris_GithubAll** 		| **30000** | **43069** | **0** 	| **0**
BenchmarkKocha_GithubAll 		| 10000 	| 171362 	| 23304 	| 843
BenchmarkMacaron_GithubAll 		| 2000 		| 817008 	| 224960 	| 2315
BenchmarkMartini_GithubAll 		| 100 		| 12609209 	| 237952 	| 2686
BenchmarkPat_GithubAll 			| 300 		| 4830398 	| 1504101 	| 32222
BenchmarkPossum_GithubAll 		| 10000 	| 301716 	| 97440 	| 812
BenchmarkR2router_GithubAll 	| 10000 	| 270691 	| 77328 	| 1182
BenchmarkRevel_GithubAll 		| 1000 		| 1491919 	| 345553 	| 5918
BenchmarkRivet_GithubAll 		| 10000 	| 283860 	| 84272 	| 1079
BenchmarkTango_GithubAll 		| 5000 		| 473821 	| 87078 	| 2470
BenchmarkTigerTonic_GithubAll 	| 2000 		| 1120131 	| 241088 	| 6052
BenchmarkTraffic_GithubAll 		| 200 		| 8708979 	| 2664762 	| 22390
BenchmarkVulcan_GithubAll 		| 5000 		| 353392 	| 19894 	| 609
BenchmarkZeus_GithubAll 		| 2000 		| 944234 	| 300688 	| 2648

With Intel(R) Core(TM) i7-4710HQ CPU @ 2.50GHz 2.50 HGz and 8GB Ram: 

![enter image description here](http://nodets.com/benchmarks_results_output.png)

* Sometimes it goes to 50169 ns/op but even then it's faster than all other.

* Note that the Iris framework does not have copied source (other than the benchmark test )  from other routers ( **I don't mean that is bad if someone do that, I love open source!**).

*  Also Iris framework doesn't uses the famous and fast enough [httprouter package](https://github.com/julienschmidt/httprouter),  Iris' approach seems to be simplier and faster. To be honesty, as I'm new to golang, I  learnt about this router a few minutes before publish this document.


## Alternatives 

Iris is not the only one framework which is fast and easy to use, [Gin](https://github.com/gin-gonic/gin) which is x40 times faster than [Martini](https://github.com/go-martini/martini)  is very good 'competitor' so I write the exact same benchmark test in order to compare Gin over Iris with Intel(R) Core(TM) i7-4710HQ CPU @ 2.50GHz 2.50 HGz and 8GB Ram, also note that Iris can use http.Handler and be more faster than it is with Context but Gin doesn't accept http.Handler as handler so both of them have their own Context as parameter to the handlers.
Let's take a look at the results: 

![enter image description here](http://nodets.com/iris_vs_gin.png)

 - Gin:   **54.636 ns/op**
 - Iris:  **50.969 ns/op**
 - Both of them with zero memory allocation!

So, Iris **is a bit faster than Gin**.
I wish Gin has compatibility with the Martini's middleware ecosystem, as Iris provides out of the box. 
**Gin is a complete web framework and the only good alternative over Iris** (that I know)  so if you don't care about performance so much (not a big difference, gin is ~4.000 nanoseconds slower only) or somehow you don't like Iris then you should get your self some [Gin](https://github.com/gin-gonic/gin). 

## Features 

**Only explicit matches:** With other routers, like http.ServeMux, a requested URL path could match multiple patterns. Therefore they have some awkward pattern priority rules, like longest match or first registered, first matched. By design of this router, a request can only match exactly one or no route. As a result, there are also no unintended matches, which makes it great for SEO and improves the user experience.

**Parameters in your routing pattern:** Stop parsing the requested URL path, just give the path segment a name and the router delivers the dynamic value to you. Because of the design of the router, path parameters are very cheap.

**Perfect for APIs:** The router design encourages to build sensible, hierarchical RESTful APIs. Moreover it has builtin native support for OPTIONS requests and 405 Method Not Allowed replies.

**Compatible:** At the end the iris is just a middleware which acts like router and a small simply web framework, this means that you can you use it side-by-side with your favorite big and well-tested web framework. Iris is fully compatible with the **net/http package.**

**Miltiple servers :** Besides the fact that iris has a default main server. You can declare a new iris using the iris.New() func. example: server1:= iris.New(); server1.Get(....); server1.Listen(9999)



## API
**Use of GET,  POST,  PUT,  DELETE, HEAD, PATCH & OPTIONS**

```go
package main

import (
	"github.com/kataras/iris"
	"net/http"
)

func main() {
	iris.Get("/home", testGet)
	iris.Post("/login",testPost)
	iris.Put("/add",testPut)
	iris.Delete("/remove",testDelete)
	iris.Head("/testHead",testHead)
	iris.Patch("/testPatch",testPatch)
	iris.Options("/testOptions",testOptions)
	
	iris.Listen(8080)	
}

//iris is fully compatible with net/http package
func testGet(res http.ResponseWriter, req *http.Request) {
	//...
}

func testPost(c iris.Context) {
	//...
}

func testPut(r iris.Renderer) {
	//...
}

func testDelete(c iris.Context, r iris.Renderer) {
	//...
}
//and so on....
```

## Named Parameters 

Named parameters are just custom paths to your routes, you can access them for each request using context's **c.Param("nameoftheparameter")**. Get all, as array (**{Key,Value}**) using **c.Params** property.

No limit on how long a path can be.

Usage: 


```go
package main

import "github.com/kataras/iris"

func main() {
	// MATCH to /hello/anywordhere
	// NOT match to /hello or /hello/ or /hello/anywordhere/something
	iris.Get("/hello/:name", func(c iris.Context) {
		name := c.Param("name")
		c.Write("Hello " + name)
	})
	
	// MATCH to /profile/kataras/friends/1
	// NOT match to /profile/ , /profile/kataras ,
	// NOT match to /profile/kataras/friends,  /profile/kataras/friends ,
	// NOT match to /profile/kataras/friends/2/something 
	iris.Get("/users/:fullname/friends/:friendId",
		func(c iris.Context, r iris.Renderer){
			name:= c.Param("fullname")
			friendId := c.ParamInt("friendId")
			r.HTML("<b> Hello </b>"+name)
		})

	iris.Listen(8080)
	//or log.Fatal(http.ListenAndServe(":8080", iris))
}

```

**Note:** Since this router has only explicit matches, you can not register static routes and parameters for the same path segment. For example you can not register the patterns /user/new and /user/:user for the same request method at the same time. The routing of different request methods is independent from each other.

## Match anything and the Static serve handler

Match everything/anything (symbol * (asterix))
```go
// Will match any request which url's preffix is "/anything/"
iris.Get("/anything/*", func(ctx iris.Context) { } )  
// Match: /anything/whateverhere , /anything/blablabla
// Not Match: /anything , /anything/ , /something
```
Pure http static  file server as handler using **iris.Static("./path/to/the/resources/directory/")**
```go
// Will match any request which url's preffix is "/public/" 
/* and continues with a file whith it's extension which exists inside the os.Gwd()(dot means working directory)+ /static/resources/
*/
iris.Any("/public/*", iris.Static("./static/resources/")) //or Get
//so simple
//Note: strip of the /public/ is handled so don't worry 
```
## Declaring routes
Iris framework has four (4) different forms of functions in order to declare a route's handler,  the typical http.Handler and one(1) annotated struct to declare a complete route.


 1. Typical classic handler function, compatible with net/http and other frameworks
	 *  **func(res http.ResponseWriter, req *http.Request)**
```go
	iris.Get("/profile/user/:userId", func(res http.ResponseWriter, req *http.Request) {
		
	})
```
 2. Context parameter in function-declaration
	 * **func(ctx iris.Context)**

```go
	iris.Get("/profile/user/:userId", func(ctx iris.Context) {
	
	})
```
 3. Renderer parameter in function-declaration
	 * **func(r iris.Renderer)**

```go
	iris.Get("/profile/user/:userId", func(r iris.Renderer) {
	
	})
```
 4. Context & Renderer parameters in function-declaration
	 * **func(c iris.Context, r iris.Renderer)**

```go
	iris.Get("/profile/user/:userId", func(ctx iris.Context, r iris.Renderer) {
	
	})
```
 5. http.Handler
	 * **http.Handler**

```go
	iris.Get("/profile/user/:userId", http.HandlerFunc(func(res http.Response, req *req.Request) {
	
	}))
```
 6. **'External' annotated struct** which directly implements the Iris Annotated interface



```go
///file: userhandler.go
import "github.com/kataras/iris"

type UserRoute struct {
	iris.Annotated `get:"/profile/user/:userId"`
}

func (u *UserHandler) Handle(ctx iris.Context, r iris.Renderer) {
	defer ctx.Close()
	userId, err := ctx.ParamInt("userId")
	//or just userId := ctx.Param("userId") and use it as string
} 


///file: main.go

//...
	iris.Handle(&UserRoute{})
//...

```
Personally I use the external struct and the **func(ctx iris.Context, r iris.Renderer)** form .
 At the next chapter you will learn what are the benefits of having the  **Context**  and the  **Renderer**  as arguments/parameters to the Route handlers.


## Context

> Variables

 1. **ResponseWriter**
	 - The ResponseWriter is the exactly the same as you used to use with the standar http library.
 2. **Request**
	 - The Request is the pointer of the *Request, is the exactly the same as you used to use with the standar http library.
 3. **Params**
	 - Contains the Named path Parameters, imagine it as a map[string]string which contains all parameters of a request.
 
>Functions

 1. **Write(contents string)**
	 - Writes a pure string to the ResponseWriter and sends to the client.
 2. **Param(key string)** returns string
	 - Returns the string representation of the key's  named parameter's value. Registed path= /profile/:name) Requested url is /profile/something where the key argument is the named parameter's key, returns the value  which is 'something' here.
 3. **ParamInt(key string)** returns integer, error
	 - Returns the int representation of the key's  named parameter's value, if something goes wrong the second return value, the error is not nil.
 4. **URLParam(key string)** returns string
	 - Returns the string representation of a requested url parameter (?key=something) where the key argument is the name of, something is the returned value.
 5. **URLParamInt(key string)** returns integer, error
	 - Returns the int representation of  a requested url parameter
 6. **SetCookie(name string, value string)**
	 - Adds a cookie to the request.
 7. **GetCookie(name string)** returns string
	 - Get the cookie value, as string, of a cookie.
 8. **ServeFile(path string)**
	 - This just calls the http.ServeFile, which serves a file given by the path argument  to the client.
 9. **NotFound()**
	 - Sends a http.StatusNotFound with a custom template you defined (if any otherwise the default template is there) to the client.
	 --- *Note: We will learn all about Custom Error Handlers later*.
 10. **Close()**
	 - Calls the Request.Body.Close().

## Renderer
>Functions

1. **WriteHTML(status int, contents string) & HTML(contents string)**
	- WriteHTML: Writes html string with a given http status to the client, it sets the Header with the correct content-type.
	- HTML: Same as WriteHTML but you don't have to pass a status, it's defaulted to http.StatusOK (200). 
2. **WriteData(status int, binaryData []byte) & Data(binaryData []byte)**
	- WriteData: Writes binary data with a given http status to the client, it sets the Header with the correct content-type.
	- Data : Same as WriteData but you don't have to pass a status, it's defaulted to http.StatusOK (200). 
3. **WriteText(status int, contents string) & Text(contents string)**
	- WriteText: Writes plain text with a given http status to the client, it sets the Header with the correct content-type.
	- Text: Same as WriteTextbut you don't have to pass a status, it's defaulted to http.StatusOK (200). 
4. **WriteJSON(status int, jsonStructs ...interface{}) & JSON(jsonStructs ...interface{}) returns error**
	- WriteJSON: Writes json which is converted from struct(s) with a given http status to the client, it sets the Header with the correct content-type. If something goes wrong then it's returned value which is an error type is not nil.
	- JSON: Same as WriteJSON but you don't have to pass a status, it's defaulted to http.StatusOK (200). 
5. **WriteXML(status int, xmlStructs ...interface{}) & XML(xmlStructs ...interface{}) returns error**
	- WriteXML: Writes writes xml which is converted from struct(s) with a given http status to the client, it sets the Header with the correct content-type. If something goes wrong then it's returned value which is an error type is not nil.
	- XML: Same as WriteXML but you don't have to pass a status, it's defaulted to http.StatusOK (200). 
6. **RenderFile(file string, pageContext interface{}) returns error**
	- RenderFile: Renders a file by its name (which a file is saved to the template cache) and a context passed to the function, default http status is http.StatusOK(200) if the template was found, otherwise http.StatusNotFound(404),  If something goes wrong then it's returned value which is an error type is not nil.
7. **Render(pageContext interface{})  returns error**
	- Render: Renders the registed and cached by the template cache file template from this particular route (if it's has children  it will render them too) file  and a context passed to the function, default http status is http.StatusOK(200) if the template was found, otherwise http.StatusNotFound(404),  If something goes wrong then it's returned value which is an error type is not nil.
	--- *Note:  We will learn how to add a template to the template cache for a route at the next chapters*.


**The next chapters are being written this time, they will be published soon, check the docs later [[TODO chapters: Register custom error handlers, Add templates to the route, Declare middlewares]]**


## Third Party Middleware
*The iris is re-written in order to support all middlewares that are already exists for [Negroni](https://github.com/codegangsta/negroni) middleware*
 
Here is a current list of compatible middlware.


| Middleware | Author | Description |
| -----------|--------|-------------|
| [RestGate](https://github.com/pjebs/restgate) | [Prasanga Siripala](https://github.com/pjebs) | Secure authentication for REST API endpoints |
| [Graceful](https://github.com/stretchr/graceful) | [Tyler Bunnell](https://github.com/tylerb) | Graceful HTTP Shutdown |
| [secure](https://github.com/unrolled/secure) | [Cory Jacobsen](https://github.com/unrolled) | Middleware that implements a few quick security wins |
| [JWT Middleware](https://github.com/auth0/go-jwt-middleware) | [Auth0](https://github.com/auth0) | Middleware checks for a JWT on the `Authorization` header on incoming requests and decodes it|
| [binding](https://github.com/mholt/binding) | [Matt Holt](https://github.com/mholt) | Data binding from HTTP requests into structs |
| [logrus](https://github.com/meatballhat/negroni-logrus) | [Dan Buch](https://github.com/meatballhat) | Logrus-based logger |
| [render](https://github.com/unrolled/render) | [Cory Jacobsen](https://github.com/unrolled) | Render JSON, XML and HTML templates |
| [gorelic](https://github.com/jingweno/negroni-gorelic) | [Jingwen Owen Ou](https://github.com/jingweno) | New Relic agent for Go runtime |
| [gzip](https://github.com/phyber/negroni-gzip) | [phyber](https://github.com/phyber) | GZIP response compression |
| [oauth2](https://github.com/goincremental/negroni-oauth2) | [David Bochenski](https://github.com/bochenski) | oAuth2 middleware |
| [sessions](https://github.com/goincremental/negroni-sessions) | [David Bochenski](https://github.com/bochenski) | Session Management |
| [permissions2](https://github.com/xyproto/permissions2) | [Alexander Rødseth](https://github.com/xyproto) | Cookies, users and permissions |
| [onthefly](https://github.com/xyproto/onthefly) | [Alexander Rødseth](https://github.com/xyproto) | Generate TinySVG, HTML and CSS on the fly |
| [cors](https://github.com/rs/cors) | [Olivier Poitrey](https://github.com/rs) | [Cross Origin Resource Sharing](http://www.w3.org/TR/cors/) (CORS) support |
| [xrequestid](https://github.com/pilu/xrequestid) | [Andrea Franz](https://github.com/pilu) | Middleware that assigns a random X-Request-Id header to each request |
| [VanGoH](https://github.com/auroratechnologies/vangoh) | [Taylor Wrobel](https://github.com/twrobel3) | Configurable [AWS-Style](http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html) HMAC authentication middleware |
| [stats](https://github.com/thoas/stats) | [Florent Messa](https://github.com/thoas) | Store information about your web application (response time, etc.) |

## Contributors

Thanks goes to the people who have contributed code to this package, see the
[GitHub Contributors page][].

[GitHub Contributors page]: https://github.com/kataras/iris/graphs/contributors



## Community

If you'd like to discuss this package, or ask questions about it, please use one
of the following:

* **Chat**: https://gitter.im/kataras/iris


## Todo
*  Complete the documents
*  Query parameters
*  Create examples in this repository

## Licence

This project is licensed under the MIT license.

