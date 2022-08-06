package main

import (
	Server "Forum/platform/newsfeed/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	// db := Accounts.SetUpDBs(MasterDB)
	// res := db.GetSpecificAccount(1)
	// fmt.Println(res)
	// post := db.GetSpecificPost(1)
	// fmt.Println(post)

	// MasterDB.AddAccount(Accounts.Data{
	// 	ID:         1,
	// 	Name:       "Luca",
	// 	FamilyName: "Bourry",
	// 	Email:      "luca.bourry@icloud.com",
	// 	Pseudo:     "lb1702",
	// 	Pass:       "azerty",
	// })

	// MasterDB.AddPost(Accounts.NewsFeed{
	// 	PostID:   1,
	// 	ID:       1,
	// 	Content:  "Salut, ceci est un post d'essai!",
	// 	Media:    "./pathToFiles",
	// 	Likes:    0,
	// 	Dislikes: 0,
	// })

	// posts := MasterDB.GetPosts()
	// fmt.Println(posts)

	Server.StartServer()
}
