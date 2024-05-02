package main

import (
	"fmt"
	"net/http"
	"time"
)

type comment struct {
	username   string
	text       string
	dateString string
}

var comments []comment

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	body := `
        <html>
            <head>
                <title>Comments</title>
                <style type="text/css">
                    body {
                        width: 500px;
                        margin: 0 auto;
                    }
                    h1 {
                        margin: 0;
                        padding: 0;
                    }
                    div {
                        padding: 20px 0;
                    }
                    textarea, input[type="text"] {
                        width: 100%;
                    }
                    textarea {
                        height: 200px;
                    }
                    .comment {
                        padding: 10px;
                        border: 1px solid dd;
                        margin-bottom: 4px;
                    }
                </style>
            </head>
        <body>`
	commentBody := ""
	for i := range comments {
		commentBody += fmt.Sprintf("<div class='comment'>%s (%s) - @%s</div>", comments[i].text, comments[i].dateString, comments[i].username)
	}
	body += fmt.Sprintf(`
        <h1>Comments</h1>
        %s
        <form method="POST" action="/post">
            <div><input type="text" placeholder="Username" name="username" /></div>
            <textarea placeholder="Comment text" name="comment"></textarea>
            <div><input type="submit" value="Submit" /></div>
        </form>
        </body>
        </html>
    `, commentBody)
	w.Write([]byte(body))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	commentText := r.Form.Get("comment")
	comments = append(comments, comment{username: username, text: commentText, dateString: time.Now().Format(time.RFC3339)})

	http.Redirect(w, r, "/comments", http.StatusFound)
}

func main() {
	http.HandleFunc("/comments", commentsHandler)
	http.HandleFunc("/post", postHandler)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}
