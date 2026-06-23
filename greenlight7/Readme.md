Our createMovieHandler now works well when it receives a valid JSON request body with
the appropriate data. But at this point you might be wondering:
What if the client sends something that isn’t JSON, like XML or some random bytes?
What happens if the JSON is malformed or contains an error?
What if the JSON types don’t match the types we are trying to decode into?
What if the request doesn’t even contain a body?

COMMANDS TO PLAY WITH 
curl -d '<?xml version="1.0" encoding="UTF-8"?><note><to>Alice</to></note>' localhost:4000/v1/movies
2. url -d '{"title": "Moana", }' localhost:4000/v1/movies
3.curl -d '["foo", "bar"]' localhost:4000/v1/movies
4.  curl -d '{"title": 123}' localhost:4000/v1/movie
5. curl -X POST localhost:4000/v1/movies
THE ERROR MESSAGES ARE RETURNED BY THE DECODE() METHOD

For a private API which won’t be used by members of the public, then this behavior is
probably fine and you needn’t do anything else.
But for a public-facing API, the error messages themselves aren’t ideal. Some are too
detailed and expose information about the underlying API implementation. Others aren’t
descriptive enough (like "EOF"), and some are just plain confusing and difficult to
understand. There isn’t consistency in the formatting or language used either.
To improve this, we’re going to explain how to triage the errors returned by Decode() and
replace them with clearer, easy-to-action, error messages to help the client debug exactly
what is wrong with their JSON.

ERRORS AND REASON
1. json.SyntaxError  and io.ErrUnexpectedEOF
There is a syntax problem with the JSON being decoded.

2. son.UnmarshalTypeError

A JSON value is not appropriate for the destination Go type.

3.json.InvalidUnmarshalError 
The decode destination is not valid (usually because it is not a
pointer). This is actually a problem with our application code,
not the JSON itself.

4. io.EOF
The JSON being decoded is empty.

TO TRIAGE THIS ERRORS , WE USE 
Error.Is() and errors.As()

E edit the create a readJSON() hhelper  function in helpers file 

again:
$ curl -d '{"title": "Moana", "rating":"PG"}' localhost:4000/v1/movies
{
    "error": "body contains unknown key \"rating\""
}
$ curl -d '{"title": "Moana"}{"title": "Top Gun"}' localhost:4000/v1/movies
{
    "error": "body must only contain a single JSON value"
}
$ curl -d '{"title": "Moana"} :~()' localhost:4000/v1/movies
{
    "error": "body must only contain a single JSON value"
}
Those are working much better now — processing of the request is terminated and the
client receives a clear error message explaining exactly what the problem is.
Lastly, let’s try making a request with a very large JSON body.
To demonstrate this, I’ve created a 1.5MB JSON file that you can download into your /tmp
directory by running the following command:
$ wget -O /tmp/largefile.json https://www.alexedwards.net/static/largefile.json
If you try making a request to your POST /v1/movies endpoint with this file as the request
body, the http.MaxBytesReader() check will kick in and you should get a response similar
to this:

If you try making a request to your POST /v1/movies endpoint with this file as the request
body, the http.MaxBytesReader() check will kick in and you should get a response similar
to this:
$ curl -d @/tmp/largefile.json localhost:4000/v1/movies
{
    "error": "body must not be larger than 1048576 bytes"
}

I must admit that the code inside readJSON() isn’t the most beautiful-looking… there’s a lot
of error handling and logic that we’ve introduced for what is ultimately a one-line call to
Decode(). But now it’s written, it’s done. You don’t need to touch it again, and it’s
something that you can copy-and-paste into other projects easily

CUSTOM JSON DECODING 

Earlier on in this book we added some custom JSON encoding behavior to our API so that
movie runtime information was displayed in the format "<runtime> mins" in our JSON
responses.
In this chapter, we’re going to look at this from the other side and update our application so
that the createMovieHandler accepts runtime information in this format.
If you try sending a request with the movie runtime in this format right now, you’ll get a
400 Bad Request response (since it’s not possible to decode a JSON string into an int32
type). Like so:
$ curl -d '{"title": "Moana", "runtime": "107 mins"}' localhost:4000/v1/movies

WE DECODE 
o make this work, what we need to do is intercept the decoding process and manually
convert the "<runtime> mins" JSON string into an int32 instead.

//WE FOLLOW THE UNMASHALER JSON.INTERFACE 
type Unmarshaler interface {
UnmarshalJSON([]byte) error
}
When Go is decoding some JSON, it will check to see if the destination type satisfies the
json.Unmarshaler interface. If it does satisfy the interface, then Go will call it’s
UnmarshalJSON() method to determine how to decode the provided JSON into the target
type. This is basically the reverse of the json.Marshaler interface that we used earlier to
customize our JSON encoding behavior.