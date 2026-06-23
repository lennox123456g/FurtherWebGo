Sending Erro Messages 

we’re
still sending them a plain-text error message from the http.Error() and http.NotFound()
functions.
In this chapter we’ll fix that by creating some additional helpers to manage errors and send
the appropriate JSON responses to our clients.

//ROUTING ERROR MESSAGES httprouter AUtomatically sends when it cat find a matching route, http in this case is allowing us to use a cutom
	//allows us to set our own custom error handlers when we initialize
	//the router. These custom handlers must satisfy the http.Handler interface, which is good
	//news for us because it means we can easily re-use the notFoundResponse() and
	//methodNotAllowedResponse() helpers that we just made


	// Convert the notFoundResponse() helper to a http.Handler using the 
	// http.HandlerFunc() adapter, and then set it as the custom error handler for 404
	// Not Found responses.
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// Likewise, convert the methodNotAllowedResponse() helper to a http.Handler and set
	// it as the custom error handler for 405 Method Not Allowed responses.
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
    test routes 

curl -i localhost:4000/foo
curl -i localhost:4000/v1/movies/abc
curl -i -X PUT localhost:4000/v1/healthcheck

Additional Information
System-generated error responses
While we’re on the topic of errors, I’d like to mention that in certain scenarios Go’s
http.Server may still automatically generate and send plain-text HTTP responses. These
scenarios include when:
The HTTP request specifies an unsupported HTTP protocol version.
The HTTP request contains a missing or invalid Host header, or multiple Host headers.
The HTTP request contains an invalid header name or value.
The HTTP request contains an unsupported Transfer-Encoding header.
The size of the HTTP request headers exceeds the server’s MaxHeaderBytes setting.
The client makes a HTTP request to a HTTPS server