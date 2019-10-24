package media

//Movie struct
type Movie struct {
	name string
	year int16
	good bool
}

// CreateMovie -
// takes in params, returns new Movie obj.
// return type prepends * to explain it is a pointer to obj
// return the pointer value (memory address) of the movie obj
func CreateMovie(name string, year int16) *Movie {
	return &Movie{name, year, false}
}

//DoILikeThis -
//because CreateMovie returns a pointer
//you can update params
func DoILikeThis(m *Movie, b bool) {
	m.good = b
}
