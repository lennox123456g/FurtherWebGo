package data

import (
	"fmt"
	"strconv"

	"errors"  //New import
	"strings" //New import
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

// Dfeine and Errro r that our UnMarshalJSON() can return if we're unable
// to convert the 	JSON string successfuly
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

//implememnt the UnmarshalJSON method on the runtime type so that it asatisfies
//json.UnMarshaler interface . IMPORTANT; Because the UnmarshalJSON() needs to modify the
//receiver (our Runtime type)  we must use a pointer receiver for this to work
// correctly. Otherwise, we will only be modifying a copy (which is then discarded when
// this method returns).

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	//We expect that the ioncoming json value will be a string format
	// "<runtime> mins", and the first thing we need to do is remove the surrounding
	// double-quotes from this string. If we can't unquote it, then we return the
	// ErrInvalidRuntimeFormat error.
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	//Split the string to isolate the part containing the number
	parts := strings.Split(unquotedJSONValue, " ")

	//sanity check the parts of the string to make sure it was in the expected format
	//If it isn't, we return the ErrInvaalidRuntimeFormat Error again
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// Otherwise, parse the string containing the number into an int32. Again, if this
	// fails return the ErrInvalidRuntimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	//COnvert the int32 to a Runtime type and assign this to the receiver. Note that we
	//use  the *operator to deference the receiver ( which is  apointer to a Runtime)
	//tpe inordere to set the  underlying value of the pointer

	*r = Runtime(i)

	return nil
}
