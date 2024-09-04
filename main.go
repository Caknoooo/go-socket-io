package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

var users = make(map[string]string) 

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	router := gin.New()
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "joinChat", func(s socketio.Conn, username string) {
		users[s.ID()] = username
		updateUserList(server)
	})

	server.OnEvent("/", "startChat", func(s socketio.Conn, data map[string]string) {
		room := createRoom(data["from"], data["to"])
		s.Join(room)
		log.Println(data["from"], "started chat with", data["to"], "in room", room)
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, data map[string]string) {
		room := createRoom(users[s.ID()], data["to"])
		server.ForEach("/", room, func(c socketio.Conn) {
			if c.ID() != s.ID() {
				c.Emit("msg", data["message"])
			}
		})
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		delete(users, s.ID())
		updateUserList(server)
		log.Println("closed", msg)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router.Use(GinMiddleware("http://localhost:3000"))
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.StaticFS("/public", http.Dir("./asset"))

	if err := router.Run("127.0.0.1:8223"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

func updateUserList(server *socketio.Server) {
	var userList []string
	for _, username := range users {
		userList = append(userList, username)
	}
	server.BroadcastToNamespace("/", "updateUserList", userList)
}

func createRoom(user1, user2 string) string {
	if user1 < user2 {
		return user1 + "_" + user2
	}
	return user2 + "_" + user1
}
