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