package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"golang.org/x/net/websocket"
)

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var clients map[string]*websocket.Conn //

func init() {
	clients = make(map[string]*websocket.Conn)
}

func generateId() string { //
	r := make([]rune, 16)
	for i := range r {
		r[i] = chars[rand.Intn(len(chars))]
	}
	return string(r)
}

func chatHandler(w http.ResponseWriter, r *http.Request) { //
	fmt.Fprintf(w, `
	<!DOCTYPE html>
		<html>
		<head>
			<title>Let's Chat</title>
			<style>
			 {
				max-width: 400px;
				margin: auto;
				font-family: system-ui, sans-serif;
			}
			.message {
				padding: 1rem 0.25rem;
				border: 1px solid black;
				margin-bottom: 0.5rem;
			}
			</style>
		</head>
		<body>
			<div id="chat">
			<h1>Chat</h1>
			<div id="messages"></div>
			<input id="message" autofocus type="text" placeholder="Enter message ..." />
			<div>
				<p>Chat members:</p>
				<ul id="chat-members"></ul>
			</div>
		</div>
		<script>
			const text = document.getElementById('message');
			const messages = document.getElementById('messages');
			const members = document.getElementById('chat-members');
			const ws = new WebSocket('ws://localhost:8081/ws');
			ws.onmessage = e => {
				const msg = JSON.parse(e.data);
				if (msg.message_type == 'joinleave') {
					members.innerHTML = '';
					msg.chat_members.forEach(member => {
						const li = document.createElement('li');
						li.innerHTML = member;
						members.appendChild(li);
					});
					return;
				}
				if (msg.message_type === 'message') {
					const message = document.createElement('div');
					message.classList.add('message');
					message.innerHTML = msg.sender_id + " said: " + msg.message;
					messages.appendChild(message);
					return;
				}
			}
			document.getElementById('message').addEventListener('keyup', e => {
				if (e.key == 'Enter') {
					ws.send(e.target.value);
					message.value = '';
				}
			});
		</script>
		</body>
		</html>
	`)
}

type servermsg struct {
	MessageType string   `json:"message_type"`
	Message     string   `json:"message,omitempty"`
	Id          string   `json:"id,omitempty"`
	SenderId    string   `json:"sender_id,omitempty"`
	ChatMembers []string `json:"chat_members"`
}

func compileChatMembers() []string { //
	var chatMembers []string
	for k, _ := range clients {
		chatMembers = append(chatMembers, k)
	}
	return chatMembers
}

func sendToClients(msg servermsg) error { //
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	log.Println(string(msgJSON))
	for k := range clients {
		if err := websocket.Message.Send(clients[k], string(msgJSON)); err != nil {
			return err
		}
	}
	return nil
}

func disconnecClient(id string) error { //
	delete(clients, id)
	if err := sendToClients(servermsg{
		MessageType: "joinleave",
		Message:     "",
		Id:          id,
		SenderId:    "",
		ChatMembers: compileChatMembers(),
	}); err != nil {
		return err
	}
	return nil
}

func ws(ws *websocket.Conn) { //
	id := generateId()
	clients[id] = ws

	join := servermsg{
		MessageType: "joinleave",
		Message:     "",
		Id:          id,
		SenderId:    "",
		ChatMembers: compileChatMembers(),
	}

	sendToClients(join)

	for { //
		var incoming string
		if err := websocket.Message.Receive(ws, &incoming); err != nil { //
			if err := disconnecClient(id); err != nil {
				log.Println(err)
			}
			break
		}
		if err := sendToClients(servermsg{ //
			MessageType: "message",
			Message:     incoming,
			Id:          "",
			SenderId:    id,
			ChatMembers: compileChatMembers(),
		}); err != nil {
			if err := disconnecClient(id); err != nil {
				log.Println(err)
			}
			break
		}
	}
}

func main() {

	http.HandleFunc("/chat", chatHandler)
	http.Handle("/ws", websocket.Handler(ws)) //
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
