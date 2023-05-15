package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"

	"docker-echo-template/api/interfaces/database"
	"docker-echo-template/api/usecase"
)

type PositionController struct {
	Interactor usecase.PositionInteractor
}

type Position struct {
	UsersId int `json:"users_id"`
	PositionX float64 `json:"position_x"`
	PositionY float64 `json:"position_y"`
	PositionZ float64 `json:"position_z"`
	RotationX float64 `json:"rotation_x"`
	RotationY float64 `json:"rotation_y"`
	RotationZ float64 `json:"rotation_z"`
}

var (
	roomPosition = make(map[string]map[*websocket.Conn]bool)
	roomPositionUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func NewPositionController(sqlHandler database.SqlHandler) *PositionController {
	return &PositionController{
		Interactor: usecase.PositionInteractor{
			PositionRepository: &database.PositionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PositionController) RoomPosition(c echo.Context) (err error) {
	room := c.Param("roomId")
	conn, err := roomPositionUpgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.JSON(500, NewError(err))
	}

	if roomPosition[room] == nil {
		roomPosition[room] = make(map[*websocket.Conn]bool)
	}

	roomPosition[room][conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			delete(roomPosition[room], conn)

			return c.JSON(500, NewError(err))
		}

		var position Position
		err = json.Unmarshal(p, &position)
		if err != nil {
			delete(roomPosition[room], conn)

			return c.JSON(500, NewError(err))
		}

		for client := range roomPosition[room] {
			err := client.WriteMessage(messageType, []byte(p))
			if err != nil {
				delete(roomPosition[room], client)
			}
		}
	}
}