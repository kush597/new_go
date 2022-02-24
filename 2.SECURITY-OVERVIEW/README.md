				UNIT 2: SECURITY OVERVIEW


Security risk overview is an accurate and thorough assessment of the potential risks and vulnerabilities to the confidentiality, integrity and availability of a program

Common Vulnerabilities and Exposures (CVE) is a database of publicly disclosed information security issues. ... CVE provides a convenient, reliable way for vendors, enterprises, academics, and all other interested parties to exchange information about cybersecurity issues.

Concurrency in Golang is the ability for functions to run independent of each other. A goroutine is a function that is capable of running concurrently with other functions. ... Popular programming languages such as Java and Python implement concurrency by using threads

Garbage collection refers to the process of managing heap memory allocation: free the memory allocations that are no longer in use and keep the memory allocations that are being used.

The OWASP Top 10 is a standard awareness document for developers and web application security. It represents a broad consensus about the most critical security risks to web applications. ... Companies should adopt this document and start the process of ensuring that their web applications minimize these risks.

Memory corruption:
 Data races are among the most common and hardest to debug types of bugs in concurrent systems. A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write. See the The Go Memory Model for details.

Here is an example of a data race that can lead to crashes and memory corruption:

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

CROSS-SITE SCRIPTING:
The server() function that handles HTTP GET requests reads the parameter param from the query string and returns it (as is) in the HTTP response. The default Content-Type response header is determined by the http.DetectContentType function which implements the algorithm described by the WhatWG spec.

package main

import "io"
import "net/http"

func server(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, r.URL.Query().Get("param"))
}

func main() {
    http.HandleFunc("/", server)
    http.ListenAndServe(":5000", nil)
}

By sending a payload with param=hello, the browser developer tool shows that the Content-Type is set to text/plain (which is not harmful and rendered as simple text by the browser).


By sending a request with param=<script>alert(1)</script>, the Content-Type of the response is set to text/html, resulting in an exposure to Cross-Site Scripting.

CSRF ATTACK:
Cross-site request forgery attacks (CSRF or XSRF for short) are used to send malicious requests from an authenticated user to a web application. The attacker can't see the responses to the forged requests, so CSRF attacks focus on state changes, not theft of data.

Here is an example program:

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.GET("/protected", func(c *gin.Context) {
		c.String(200, csrf.GetToken(c))
	})

	r.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	r.Run(":8080")
}





PACKAGES TO BE INSTALLED
USE -d ACCORDING TO REQUIREMENTS

$GO GET hash/fnv
$GO GET	log
$GO GET	math/rand
$GO GET	os
$GO GET	sync
$GO GET	time
$GO GET net/http
$GO GET text/template
$GO GET	github.com/gin-contrib/sessions
$GO GET	github.com/gin-contrib/sessions/cookie
$GO GET	github.com/gin-gonic/gin
$GO GET	github.com/utrack/gin-csrf