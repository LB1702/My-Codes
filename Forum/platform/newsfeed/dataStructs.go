package Accounts

type Item struct {
	ID      int
	Content string
}

type Data struct {
	ID             int
	Name           string
	FamilyName     string
	Email          string
	Pseudo         string
	Pass           string
	ProfilePicture string
}

type NewsFeed struct {
	ID       int
	UserID   int
	Content  string
	Media    string
	Likes    int
	Dislikes int
}

type Comment struct {
	ID       int
	PostID   int
	UserID   int
	Content  string
	Media    string
	Likes    int
	Dislikes int
}

type Package struct {
	Posts    []NewsFeed
	Comments []Comment
	Session  Data
}

type Post struct {
	Content  NewsFeed
	Owner    Data
	Comments []Comment
	Session  Data
}
