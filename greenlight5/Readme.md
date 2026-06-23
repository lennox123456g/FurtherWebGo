DECODING JSON FROM CLIENT INTO STRUCT 

Just like JSON encoding, there are two approaches that you can take to decode JSON into a
native Go object: using a json.Decoder type or using the json.Unmarshal() function.
Both approaches have their pros and cons, but for the purpose of decoding JSON from a
HTTP request body, using json.Decoder is generally the best choice. It’s more efficient than
json.Unmarshal(), requires less code, and offers some helpful settings that you can use to
tweak its behavior

There are few important and interesting things about this code to point out:
When calling Decode() you must pass a non-nil pointer as the target decode destination.
If you don’t use a pointer, it will return a json.InvalidUnmarshalError error at runtime.
If the target decode destination is a struct — like in our case — the struct fields must be
exported (start with a capital letter). Just like with encoding, they need to be exported so
that they’re visible to the encoding/json package.
When decoding a JSON object into a struct, the key/value pairs in the JSON are mapped
to the struct fields based on the struct tag names. If there is no matching struct tag, Go
will attempt to decode the value into a field that matches the key name (exact matches
are preferred, but it will fall back to a case-insensitive match). Any JSON key/value pairs
which cannot be successfully mapped to the struct fields will be silently ignored.
There is no need to close r.Body after it has been read. This will be done automatically
by Go’s http.Server, so you don’t have 

Go to movies.go 

NOTE 

There are few important and interesting things about this code to point out:
When calling Decode() you must pass a non-nil pointer as the target decode destination.
If you don’t use a pointer, it will return a json.InvalidUnmarshalError error at runtime.
If the target decode destination is a struct — like in our case — the struct fields must be
exported (start with a capital letter). Just like with encoding, they need to be exported so
that they’re visible to the encoding/json package.
When decoding a JSON object into a struct, the key/value pairs in the JSON are mapped
to the struct fields based on the struct tag names. If there is no matching struct tag, Go
will attempt to decode the value into a field that matches the key name (exact matches
are preferred, but it will fall back to a case-insensitive match). Any JSON key/value pairs
which cannot be successfully mapped to the struct fields will be silently ignored.
There is no need to close r.Body after it has been read. This will be done automatically
by Go’s http.Server, so you don’t have too

TO TEST 
# Create a BODY variable containing the JSON data that we want to send.
$ BODY='{"title":"Moana","year":2016,"runtime":107, "genres":["animation","adventure"]}'
# Use the -d flag to send the contents of the BODY variable as the HTTP request body.
# Note that curl will default to sending a POST request when the -d flag is used.
$ curl -i -d "$BODY" localhost:4000/v1/movies

BETETR 
BODY='{"title":"Moana","year":2016,"runtime":107,"genres":["animation","adventure"]}'
curl -i -X POST -H "Content-Type: application/json" -d "$BODY" http://localhost:4000/v1/movies