package data

import (
	"time"
)

type Movie struct {
	ID        int64     //unique integer ID or the  ovie
	CreatedAt time.Time //Tiimestamp for when the movie is added toi the adatabase
	Title     string    //Movie Title
	Year      int32     //Movie release Year
	Runtime   int32     //Movie runtime in minutes
	Genres    []string  //slice of genres for the movies
	Version   int32     //version number  that starts at 1 an sdincerses with update
}
