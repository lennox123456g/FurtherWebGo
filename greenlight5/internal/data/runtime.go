package data

import (
	"fmt"
	"strconv"
)

//Declare a custom Runtime type, which has the underlying type int32 (the same as )
//our movie struct field

type Runtime int32

//Impolement a MarshalJSON() method on Runtime type so that it satifies the
//JSON.Marsaler interface. This should return the JSON-encoded value for the movie
//runtime (in our case, it will return a string in the ormat "<runtime> mins")

func (r Runtime) MarshalJSON() ([]byte, error) {
	//Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	//Use the stringconv.Quote() function on the string to wrap it in double quotes.It
	//needs to be surrounded by double quotes in order to be valid *JSON string
	quotedJSONValue := strconv.Quote(jsonValue)

	//convert the quoted string value to a byte slice and return it
	return []byte(quotedJSONValue), nil
}

//Go pdate the MOVIE STRUCT

//1.If your MarshalJSON() method returns a JSON string value, like ours does, then you must
//wrap the string in double quotes before returning it. Otherwise it won’t be interpreted as
//a JSON string and you’ll receive a runtime error similar to this:
//json: error calling MarshalJSON for type data.Runtime: invalid character 'm' after top-level value
//2.We’re deliberately using a value receiver for our MarshalJSON() method rather than a
//pointer receiver like func (r *Runtime) MarshalJSON(). This gives us more flexibility
//because it means that our custom JSON encoding will work on both Runtime values and
//pointers to Runtime values. As Effective Go mentions:
//The rule about pointers vs. values for receivers is that value methods can be invoked on
//pointers and values, but pointer methods can only be invoked on pointers.
