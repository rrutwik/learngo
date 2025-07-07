package files

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func (p Post) String() string {
	return p.Title
}
