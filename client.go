package main

import (
	"github.com/gorilla/websocket" // パッケージインストール
)

// clientはチャットを行っている一人のユーザを表します
type client struct {
	// socketはこのクライアントのためのWebSocketです
	socket *websocket.Conn
	// sendはメッセージが送られるチャネルです
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームです
	room *room
}

// データ読み込む
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			// 受け取ったメッセージはroomのforwardに送られる
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// データ書き出し
func (c *client) write() {
	// sendチャネルから断続的にメッセージを受け取る
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
