package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnWrapper struct {
	Conn   *websocket.Conn
	Closed bool
}

type ConnMap struct {
	sync.Map
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conns = ConnMap{}
	once  sync.Once
)

// Send an error message and close the connection
func sendErrorAndClose(conn *websocket.Conn, status int, message string, id uuid.UUID) {
	err := conn.WriteJSON(map[string]interface{}{
		"status":  status,
		"message": message,
	})
	if err != nil {
		log.Println(err)
	}
	conn.Close()
	conns.Delete(id)
}

func handleSocket(ctx *octopus.Context) {
	conn, err := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	id := uuid.New()
	conns.Store(id, &ConnWrapper{Conn: conn, Closed: false})

	// Start a goroutine to send data to clients
	once.Do(func() {
		go func() {
			log.Println("Starting goroutine")
			for {
				value := <-models.Data
				key, ok := value["key"].(string)
				val, okval := value["data"].(map[string]interface{})

				if ok && okval {
					conns.Range(func(k, v interface{}) bool {
						kId, ok := k.(uuid.UUID)

						if ok {
							if v, ok := v.(*ConnWrapper); ok && !v.Closed {
								// Set a pong handler to check the connection status
								v.Conn.SetPongHandler(func(string) error {
									v.Conn.SetReadDeadline(time.Now().Add(time.Second * 60))
									return nil
								})
								// Send a ping message to the client
								if err := v.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
									v.Closed = true
									conns.Delete(kId)
									return true
								}
								// Send the actual data
								err := v.Conn.WriteJSON(map[string]interface{}{
									"data": val,
									"type": key,
								})
								if err != nil {
									v.Closed = true
									conns.Delete(kId)
								}
							}
						}

						return true
					})
				}
				// models.Data.Range(func(kk, value interface{}) bool {
				// 	key, ok := kk.(string)
				// 	val, okval := value.(map[string]interface{})
				// 	if ok && okval {
				// 		conns.Range(func(k, v interface{}) bool {
				// 			kId, ok := k.(uuid.UUID)

				// 			if ok {
				// 				if v, ok := v.(*ConnWrapper); ok && !v.Closed {
				// 					// Set a pong handler to check the connection status
				// 					v.Conn.SetPongHandler(func(string) error {
				// 						v.Conn.SetReadDeadline(time.Now().Add(time.Second * 60))
				// 						return nil
				// 					})
				// 					// Send a ping message to the client
				// 					if err := v.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				// 						v.Closed = true
				// 						conns.Delete(kId)
				// 						return true
				// 					}
				// 					// Send the actual data
				// 					err := v.Conn.WriteJSON(map[string]interface{}{
				// 						"data": val,
				// 						"type": key,
				// 					})
				// 					if err != nil {
				// 						v.Closed = true
				// 						conns.Delete(kId)
				// 					}
				// 				}
				// 			}

				// 			return true
				// 		})
				// 	}
				// 	models.Data.Delete(key)
				// 	return true
				// })
				// time.Sleep(time.Second)
			}
		}()
	})

	for {

		var data = map[string]interface{}{}
		err := conn.ReadJSON(&data)
		if err != nil {
			sendErrorAndClose(conn, http.StatusBadRequest, "Invalid message", id)
			return
		}

		if data["type"] == "private_message" {
			msg, ok := data["message"].(map[string]interface{})
			if !ok {
				sendErrorAndClose(conn, http.StatusBadRequest, "Invalid message", id)
				return
			}

			newMesssage := new(models.PrivateMessage)
			newMesssage.Content = msg["content"].(string)
			newMesssage.SenderID = uuid.MustParse(msg["sender_id"].(string))
			newMesssage.ReceiverID = uuid.MustParse(msg["receiver_id"].(string))

			if newMesssage.Content == "" || newMesssage.SenderID == uuid.Nil || newMesssage.ReceiverID == uuid.Nil {
				sendErrorAndClose(conn, http.StatusBadRequest, "Invalid message", id)
				return
			}

			user := models.User{}
			if user.Get(ctx.Db.Conn, newMesssage.ReceiverID) != nil || user.Get(ctx.Db.Conn, newMesssage.SenderID) != nil {
				sendErrorAndClose(conn, http.StatusNotFound, "User not found", id)
				return
			}

			if err := newMesssage.Create(ctx.Db.Conn); err != nil {
				sendErrorAndClose(conn, http.StatusInternalServerError, "Internal server error", id)
				return
			}

			newNotification := models.Notification{
				UserID:    newMesssage.SenderID,
				ConcernID: newMesssage.ReceiverID,
				Type:      models.TypeNewMessage,
				Message:   newMesssage.Content,
			}
			if err := newNotification.Create(ctx.Db.Conn); err != nil {
				sendErrorAndClose(conn, http.StatusInternalServerError, "Internal server error", id)
				return
			}
		} else if data["type"] == "group_message" {
			// Handle group messages
		}
	}
}

var handleSocketRoute = route{
	path:   "/socket",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AllowedSever,
		handleSocket,
	},
}

func init() {
	AllHandler[handleSocketRoute.path] = handleSocketRoute
}
