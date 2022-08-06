package Accounts

import (
	"database/sql"
	"fmt"
)

type Feed struct {
	DB *sql.DB
}

// ACCOUNT SECTION
func (feed *Feed) AddAccount(data Data) {
	statement, err := feed.DB.Prepare(`
	INSERT INTO 'Accounts' ('name', 'familyName','email', 'pseudo', 'pass', profilePicture) VALUES
  (?,?,?,?,?,?);
	`)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("OK")
	}

	statement.Exec(data.Name, data.FamilyName, data.Email, data.Pseudo, data.Pass, data.ProfilePicture)
}

func (feed *Feed) GetAccounts() []Data {
	datas := []Data{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM Accounts
	`)
	var id int
	var name string
	var familyName string
	var email string
	var pseudo string
	var pass string
	var profilePicture string
	for rows.Next() {
		rows.Scan(&id, &name, &familyName, &email, &pseudo, &pass, &profilePicture)
		payload := Data{
			ID:             id,
			Name:           name,
			FamilyName:     familyName,
			Email:          email,
			Pseudo:         pseudo,
			Pass:           pass,
			ProfilePicture: profilePicture,
		}
		datas = append(datas, payload)
	}
	return datas
}

func (feed *Feed) GetSpecificAccount(id int) Data {
	data := Data{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM Accounts where ID == ?
	`, id)
	var name string
	var familyName string
	var email string
	var pseudo string
	var pass string
	var profilePicture string
	for rows.Next() {
		rows.Scan(&id, &name, &familyName, &email, &pseudo, &pass, &profilePicture)
		payload := Data{
			ID:             id,
			Name:           name,
			FamilyName:     familyName,
			Email:          email,
			Pseudo:         pseudo,
			Pass:           pass,
			ProfilePicture: profilePicture,
		}
		data = payload
	}
	return data
}

// POST SECTION
func (feed *Feed) AddPost(post NewsFeed) {
	statement, err := feed.DB.Prepare(`
	INSERT INTO 'NewsFeed' ('UserID', 'Content', 'Media', 'Likes','Dislikes') VALUES
  (?,?,?,?,?);
	`)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("OK")
	}

	statement.Exec(post.UserID, post.Content, post.Media, post.Likes, post.Dislikes)
}

func (feed *Feed) GetPosts() []NewsFeed {
	datas := []NewsFeed{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM NewsFeed
	`)
	var id int
	var userId int
	var content string
	var media string
	var likes int
	var dislikes int
	for rows.Next() {
		rows.Scan(&id, &userId, &content, &media, &likes, &dislikes)
		payload := NewsFeed{
			ID:       id,
			UserID:   userId,
			Content:  content,
			Media:    media,
			Likes:    likes,
			Dislikes: dislikes,
		}
		datas = append(datas, payload)
	}
	return datas
}

func (feed *Feed) GetSpecificPost(postId int) NewsFeed {
	data := NewsFeed{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM NewsFeed where ID == ?
	`, postId)
	var id int
	var userId int
	var content string
	var media string
	var likes int
	var dislikes int
	for rows.Next() {
		rows.Scan(&id, &userId, &content, &media, &likes, &dislikes)
		payload := NewsFeed{
			ID:       id,
			UserID:   userId,
			Content:  content,
			Media:    media,
			Likes:    likes,
			Dislikes: dislikes,
		}
		data = payload
	}
	return data
}

// COMMENTS SECTION
func (feed *Feed) AddComment(comment Comment) {
	statement, err := feed.DB.Prepare(`
	INSERT INTO 'Comments' ('PostID','UserID', 'Media','Content', 'Likes','Dislikes') VALUES
  (?,?,?,?,?,?);
	`)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("OK")
	}

	statement.Exec(comment.PostID, comment.UserID, comment.Media, comment.Content, comment.Likes, comment.Dislikes)
}

func (feed *Feed) GetComments() []Comment {
	comments := []Comment{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM Comments
	`)
	var id int
	var postId int
	var userId int
	var content string
	var media string
	var likes int
	var dislikes int
	for rows.Next() {
		rows.Scan(&id, &postId, &userId, &media, &content, &likes, &dislikes)
		payload := Comment{
			ID:       id,
			PostID:   postId,
			UserID:   userId,
			Media:    media,
			Content:  content,
			Likes:    likes,
			Dislikes: dislikes,
		}
		comments = append(comments, payload)
	}
	return comments
}

func (feed *Feed) GetCommentsByPostID(postId int) []Comment {
	comments := []Comment{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM Comments WHERE "PostID" == ?
	`, postId)
	var id int
	var userId int
	var content string
	var media string
	var likes int
	var dislikes int
	for rows.Next() {
		rows.Scan(&id, &postId, &userId, &media, &content, &likes, &dislikes)
		payload := Comment{
			ID:       id,
			PostID:   postId,
			UserID:   userId,
			Media:    media,
			Content:  content,
			Likes:    likes,
			Dislikes: dislikes,
		}
		comments = append(comments, payload)
	}
	return comments
}

func (feed *Feed) GetSpecificComment(commentId int) Comment {
	comment := Comment{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM Comments where ID == ?
	`, commentId)
	var postId int
	var userId int
	var media string
	var content string
	var likes int
	var dislikes int
	for rows.Next() {
		rows.Scan(&commentId, &postId, &userId, &media, &content, &likes, &dislikes)
		payload := Comment{
			ID:       commentId,
			PostID:   postId,
			UserID:   userId,
			Media:    media,
			Content:  content,
			Likes:    likes,
			Dislikes: dislikes,
		}
		comment = payload
	}
	return comment
}

func SetUpDBs(db *sql.DB) *Feed {
	statement, _ := db.Prepare(`

	CREATE TABLE IF NOT EXISTS 'Accounts' (
			"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"name"	TEXT,
			"familyName"	TEXT,
			"email"	TEXT,
			"pseudo"	TEXT,
			"pass"	TEXT,
			"profilePicture"	TEXT
		);
	
	`)

	statement1, _ := db.Prepare(`

	CREATE TABLE IF NOT EXISTS 'NewsFeed' (
			"ID"	INTEGER PRIMARY KEY AUTOINCREMENT,
			"UserID"	INTEGER,
			"Content"	TEXT,
			"Media"	TEXT,
			"Likes"	INTEGER,
			"Dislikes"	INTEGER
		);
	
	`)

	statement2, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS 'Comments' (
		"ID"	INTEGER PRIMARY KEY AUTOINCREMENT,
		"PostID"	INTEGER,
		"UserID"	INTEGER,
		"Media"	TEXT,
		"Content"  TEXT,
		"Likes"	INTEGER,
		"Dislikes"	INTEGER
	);
	
	`)

	statement.Exec()
	statement1.Exec()
	statement2.Exec()

	return &Feed{
		DB: db,
	}
}

func initDBs() *Feed {

	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := SetUpDBs(MasterDB)

	return db

}
