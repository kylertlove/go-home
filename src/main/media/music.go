package media

//Music struct
type Music struct {
	Albums     []string
	ArtistName string
}

//CreateRecord builds new music struct
func CreateRecord(name string, albums []string) Music {
	return Music{albums, name}
}
