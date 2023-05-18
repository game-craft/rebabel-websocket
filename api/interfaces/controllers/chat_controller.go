package controllers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"

	"docker-echo-template/api/domain"
	"docker-echo-template/api/interfaces/database"
	"docker-echo-template/api/usecase"
)

type ChatController struct {
	Interactor usecase.ChatInteractor
}

var (
	worldChat = make(map[string]map[*websocket.Conn]bool)
	worldChatUpgrader = websocket.Upgrader{
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

func (controller *ChatController) WorldChat(c echo.Context) (err error) {
	world := c.Param("worldsId")
	worldId, _ := strconv.Atoi(world)
	conn, err := worldChatUpgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	if worldChat[world] == nil {
		worldChat[world] = make(map[*websocket.Conn]bool)
	}

	worldChat[world][conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			delete(worldChat[world], conn)

			return c.JSON(500, NewError(err))
		}

		message := string(p)
		go controller.addChatRecord(c, worldId, message)

		for client := range worldChat[world] {
			err := client.WriteMessage(messageType, []byte(message))
			if err != nil {
				delete(worldChat[world], client)
			}
		}
	}
}

func (controller *ChatController) addChatRecord(c echo.Context, worldId int, message string) {
	u := domain.Chat{}
	c.Bind(&u)
	u.WorldsId = worldId
	u.ChatsContent = message
	controller.Interactor.Add(u)
}