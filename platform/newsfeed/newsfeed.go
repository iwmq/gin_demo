package newsfeed

// Getter ...
type Getter interface {
	GetAll() []Item
}

// Adder ...
type Adder interface {
	Add(item Item)
}

// Item ...
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// Repo ...
type Repo struct {
	Items []Item
}

// New ...
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add ...
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll ...
func (r *Repo) GetAll() []Item {
	return r.Items
}
