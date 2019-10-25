package media

//Movie struct
type Movie struct {
	Name string
	Year int16
	Good bool
}

// CreateMovie -
// takes in params, returns new Movie obj.
// return type prepends * to explain it is a pointer to obj
// return the pointer value (memory address) of the movie obj
func CreateMovie(name string, year int16) Movie {
	m := Movie{
		name,
		year,
		false,
	}
	return m
}

//DoILikeThis -
//because CreateMovie returns a pointer
//you can update params
func DoILikeThis(m *Movie, b bool) {
	m.Good = b
}
