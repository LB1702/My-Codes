package Server

import (
	Accounts "Forum/platform/newsfeed"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("yaxsomo-seurity"))
var tpl *template.Template
var accountID = 0
var accountSession Accounts.Data
var accountRegData Accounts.Data
var temp string

func Registration(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	name := r.FormValue("name")
	familyName := r.FormValue("familyName")
	email := r.FormValue("email")
	pseudo := r.FormValue("pseudo")
	password := r.FormValue("password")
	fmt.Println(name, familyName, email, pseudo, password)
	clear := name != "" && familyName != "" && email != "" && pseudo != "" && password != ""
	if clear {
		accountToAdd := Accounts.Data{
			Name:       name,
			FamilyName: familyName,
			Email:      email,
			Pseudo:     pseudo,
			Pass:       password,
		}

		db.AddAccount(accountToAdd)
		accountRegData = accountToAdd
		fmt.Println("Successfully added to Database")
		db.DB.Close()
		http.Redirect(w, r, "/registrationSuccess", http.StatusFound)
	}

	t, _ := template.ParseFiles("./static/signin.html")
	t.Execute(w, nil)
}

func registrationSuccessHandler(w http.ResponseWriter, r *http.Request) {
	// MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	// db := Accounts.SetUpDBs(MasterDB)
	// accounts := db.GetAccounts()
	// accountSession = accounts[accountID]
	// fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/registrationSuccess.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, accountRegData.Name)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	session, _ := store.Get(r, "session")
	value, ok := session.Values["userID"]
	fmt.Println("ok: ", ok)
	if value != nil {
		http.Redirect(w, r, "/index", http.StatusFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	accounts := db.GetAccounts()
	email := r.FormValue("email")
	password := r.FormValue("password")
	success := false
	accountID = 0
	fmt.Println("values : " + email + " | " + password)
	if email != "" && password != "" {
		for i := 0; i < len(accounts); i++ {
			if accounts[i].Email == email && accounts[i].Pass == password {
				accountID = i
				success = true
				break
			}
		}

		if success {
			fmt.Println("Welcome " + accounts[accountID].Name + "(present into database)")
			session, err := store.Get(r, "session")
			if err != nil {
				panic(err)
			}
			fmt.Println(accountID)
			session.Values["userID"] = accountID
			session.Save(r, w)
			http.Redirect(w, r, "/index", http.StatusFound)
			return
			// t, err = template.ParseFiles("./static/loginAuth.html")
			// if err != nil {
			// 	panic(err)
			// }
			// t.Execute(w, nil)
			// tpl.ExecuteTemplate(w, "/index", "Logged in")
		} else {
			fmt.Println("Not present in database or anything typed.")
		}
		db.DB.Close()
	}
	t, err := template.ParseFiles("./static/login.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	accounts := db.GetAccounts()
	accountSession = accounts[accountID]
	// fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/contact.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, accountSession)
}

func discussionsHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	accounts := db.GetAccounts()
	accountSession = accounts[accountID]
	posts := db.GetPosts()
	comments := db.GetComments()
	content := r.FormValue("content")
	mediaBase64 := r.FormValue("imageConverted")
	pack := Accounts.Package{
		Posts:    posts,
		Comments: comments,
		Session:  accountSession,
	}

	// fmt.Println(content)
	// fmt.Println(mediaBase64)
	var checkContent = content != ""
	if checkContent {
		db.AddPost(Accounts.NewsFeed{
			ID:       0,
			UserID:   accountID + 1,
			Content:  content,
			Media:    mediaBase64,
			Likes:    0,
			Dislikes: 0,
		})
		fmt.Println("Post added successfully!")
		http.Redirect(w, r, "/discussions", http.StatusFound)
	}

	// fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/pagediscussion.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, pack)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	accounts := db.GetAccounts()
	accountSession = accounts[accountID]
	// fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/profile.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, accountSession)
}

func firstPageHandler(w http.ResponseWriter, r *http.Request) {
	// MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	// db := Accounts.SetUpDBs(MasterDB)
	// accounts := db.GetAccounts()
	// accountSession = accounts[accountID]
	// // fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/Accueil.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func deniedPageHandler(w http.ResponseWriter, r *http.Request) {
	// MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	// db := Accounts.SetUpDBs(MasterDB)
	// accounts := db.GetAccounts()
	// accountSession = accounts[accountID]
	// // fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/Denied.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexHandler(): Running...")
	session, _ := store.Get(r, "session")
	_, ok := session.Values["userID"]
	fmt.Println(session.Values) //
	fmt.Println("ok: ", ok)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	accounts := db.GetAccounts()
	accountSession = accounts[accountID]
	fmt.Println(accountSession.Name)
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, accountSession)

}

func postPageHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/Post.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/html")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	MasterDB, _ := sql.Open("sqlite3", "./MasterDB.db")
	db := Accounts.SetUpDBs(MasterDB)
	accounts := db.GetAccounts()
	// fmt.Println(count)
	currentPath := r.URL.Path
	pathSplitted := strings.Split(currentPath, "/")
	// fmt.Println(pathSplitted[2])
	if _, err := strconv.Atoi(pathSplitted[2]); err == nil {
		fmt.Println("trueee")
		postID, _ := strconv.Atoi(pathSplitted[2])
		fmt.Println("Account id: " + strconv.Itoa(accountID))
		post := db.GetSpecificPost(postID)
		// fmt.Println(post.ID)
		owner := accounts[post.UserID-1]
		comments := db.GetCommentsByPostID(postID)
		fmt.Println(comments)
		accounts := db.GetAccounts()
		pageContent := Accounts.Post{
			Content:  post,
			Owner:    owner,
			Comments: comments,
			Session:  accounts[accountID],
		}

		commentContent := r.FormValue("commentSection")
		if commentContent != "" {
			// fmt.Println(post.ID)
			db.AddComment(Accounts.Comment{
				PostID:   post.ID,
				UserID:   accountID,
				Media:    "",
				Content:  commentContent,
				Likes:    0,
				Dislikes: 0,
			})
			fmt.Println("Comment added successfully!")
			// http.Redirect(w, r, "/post/"+strconv.Itoa(pageContent.Content.ID), http.StatusAccepted)
		}

		// fmt.Println(commentContent)

		// fmt.Println(post.Content)
		t, err := template.ParseFiles("./static/Post.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, pageContent)
		return
	}
	// fmt.Println("No post.")

	t.Execute(w, nil)
	// fmt.Println(pathSplitted)
	// fmt.Println(currentPath == "/post")
	// fmt.Println(currentPath)

}

func StartServer() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/accueil", firstPageHandler)
	http.HandleFunc("/denied", deniedPageHandler)
	http.HandleFunc("/signin", Registration)
	http.HandleFunc("/registrationSuccess", registrationSuccessHandler)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/post", postPageHandler)
	http.HandleFunc("/post/", postPageHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/discussions", discussionsHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/index", indexHandler)
	err := http.ListenAndServe(":8000", context.ClearHandler(http.DefaultServeMux)) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

// func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
// 	expire := time.Now().Add(ttl)
// 	cookie := http.Cookie{
// 		Name:    name,
// 		Value:   value,
// 		Expires: expire,
// 	}
// 	http.SetCookie(w, &cookie)
// }
