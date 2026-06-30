VALIDATING JSON INPUT

In many cases, you’ll want to perform additional validation checks on the data from a client
to make sure it meets your specific business rules before processing it. In this chapter we’ll
illustrate how to do that in the context of a JSON API by updating our createMovieHandler
to check that:
The movie title provided by the client is not empty and is not more than 500 bytes long.
The movie year is not empty and is between 1888 and the current year.
The movie runtime is not empty and is a positive integer.
The movie has between one and five (unique) genres

If any of those checks fail, we want to send the client a 422 Unprocessable Entity response
along with error messages which clearly describe the validation failures.

We created ta validator package under inetrnals 

NOTE
In the code above we’ve defined a custom Validator type which contains a map of errors.
The Validator type provides a Check() method for conditionally adding errors to the map,
and a Valid() method which returns whether the errors map is empty or not. We’ve also
added PermittedValue(), Matches() and Unique() functions to help us perform some
specific validation checks.
Conceptually this Validator type is quite basic, but that’s not a bad thing. As we’ll see over
the course of this book, it’s surprisingly powerful in practice and gives us a lot of flexibility
and control over validation checks and how we perform them.


When you’re looking at this code, there might be a couple of questions in your head.
Firstly, you might be wondering why we’re initializing the Validator instance in our handler
and passing it to the ValidateMovie() function — rather than initializing it in
ValidateMovie() and passing it back as a return value.
This is because as our application gets more complex we will need call multiple validation
helpers from our handlers, rather than just one like we are above. So initializing the
Validator in the handler, and then passing it around, gives us more flexibility.
You might also be wondering why we’re decoding the JSON request into the input struct
and then copying the data across, rather than just decoding into the Movie struct directly.
The problem with decoding directly into a Movie struct is that a client could provide the
keys id and version in their JSON request, and the corresponding values would be
decoded without any error into the ID and Version fields of the Movie struct — even though
we don’t want them to be. We could check the necessary fields in the Movie struct after the
event to make sure that they are empty, but that feels a bit hacky, and decoding into an
intermediary struct (like we are in our handler) is a cleaner, simpler, and more robust
approach — albeit a little bit verbose

//play aroung with 
// BODY='{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}'
//curl -i -d "$BODY" localhost:4000/v1/movies