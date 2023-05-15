package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"

	"docker-echo-template/api/interfaces/database"
	"docker-echo-template/api/usecase"
)

type ChatController struct {
	Interactor usecase.ChatInteractor
}

var (
	roomChat = make(map[string]map[*websocket.Conn]bool)
	roomChatUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func NewChatController(sqlHandler database.SqlHandler) *ChatController {
	return &ChatController{
		Interactor: usecase.ChatInteractor{
			ChatRepository: &database.ChatRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ChatController) RoomChat(c echo.Context) (err error) {
	room := c.Param("roomId")
	conn, err := roomChatUpgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	if roomChat[room] == nil {
		roomChat[room] = make(map[*websocket.Conn]bool)
	}

	roomChat[room][conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			delete(roomChat[room], conn)

			return c.JSON(500, NewError(err))
		}

		message := string(p)

		for client := range roomChat[room] {
			err := client.WriteMessage(messageType, []byte(message))
			if err != nil {
				delete(roomChat[room], client)
			}
		}
	}
}