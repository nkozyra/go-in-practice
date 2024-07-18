package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

var directory string
var watcher *fsnotify.Watcher

type FileUpdateInfo struct { // # A
	Name      string `json:"name"`
	Op        string `json:"operation"`
	SizeBytes int    `json:"size"`
}

func init() {
	if osdir := os.Getenv("SSE_DIRECTORY"); osdir == "" {
		panic("SSE_DIRECTORY environment variable not set")
	} else {
		directory = osdir
	}
}

func main() {
	var err error
	watcher, err = fsnotify.NewWatcher() // # B
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	watcher.Add(directory)

	http.HandleFunc("/sse", sseHandler)
	http.HandleFunc("/files", filesHandler)
	http.ListenAndServe(":8080", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher) // # C
	if !ok {
		http.Error(w, "byte streams not supported by your client", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream") // # D

	changes := make(chan FileUpdateInfo)
	go fileListener(r.Context(), changes) // # E
	for change := range changes {
		changeJSON, err := json.Marshal(change)
		if err != nil {
			log.Println(err)
			continue
		}
		response := fmt.Sprintf("event: file-update\ndata: %s\n\n", changeJSON) // # F
		fmt.Fprint(w, response)                                                 // # F
		flusher.Flush()
	}
}

func filesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head><title>File changes</title></head>
<style>
    body > div {
        margin: auto;
        max-width: 800px;
    }
    .message {
        padding: 1rem;
        background-color: grey;
        margin-bottom: 0.5rem;
    }
    .message.create {
        background-color: palegreen;
    }
    .message.remove {
        background-color: tomato;
    }
</style>
<body>
    <div>
        <h1>File changes</h1>
        <div id="files"></div>
    </div>
    <script>
        const sse = new EventSource('/sse');
        const files = document.getElementById('files');
        const createMessage = (message) => {
            const div = document.createElement('div');
            div.classList.add('message');
            div.classList.add(message.operation.toLowerCase());
            div.innerText = message.operation + ':' + message.name + ' (' + message.size + ' bytes)';
            return div;
        }
        sse.onmessage = (e) => {
            alert(e.data);
        }
        sse.addEventListener('file-update', (e) => {
            console.log(e, e.data);
            files.prepend(createMessage(JSON.parse(e.data)));

        });
    </script>
</body>
</html>`) // # G
}

func fileListener(ctx context.Context, changes chan<- FileUpdateInfo) {
	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-watcher.Events:
			if !ok {
				break
			}
			size := 0
			if stat, err := os.Stat(event.Name); err == nil {
				size = int(stat.Size())
			}
			changes <- FileUpdateInfo{event.Name, event.Op.String(), size} // # H
		}
	}
}
