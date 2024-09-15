package chat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"

	"chatProject/routes/shared"
)

type route struct {
	state State
}

type State struct {
	rooms *ThreadSafeRoomList
}

func NewState() State {
	return State{
		rooms: NewRoomList(),
	}
}

func ConfigureRoutes(e *echo.Echo, state State) {
	routes := route{state}

	e.GET("/chat:id", routes.chat)
	e.GET("/chat/connection", routes.chatWebsocket)
	e.POST("/chat", routes.createChat)
}

func (route *route) chat(ctx echo.Context) error {
	ctx.Logger().Debug("salam!")
	return shared.Page(ctx, page(), scripts())
}

func (route *route) chatWebsocket(ctx echo.Context) error {
	handler := func(ws *websocket.Conn) {
		handle(ws, route.state)
	}

	websocket.Handler(handler).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func (route *route) createChat(ctx echo.Context) error {
	room := NewRoom("salam")
	route.state.rooms.Add(&room)

	return ctx.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/chat/%v", "salam"))
}

func handle(ws *websocket.Conn, state State) {
	defer ws.Close()
	receiver, sender := websocket.Message.Receive, websocket.Message.Send

	// clients sends in first payload which gets their username(TODO!)

	broker := state.rooms.Get("salam").Broker

	// subscribe to receive messages from other clients
	subCh := broker.Subscribe()
	defer broker.Unsubscribe(subCh)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer cancel()
		for msg := range subCh {
			select {
			case <-ctx.Done():
				return
			default:
			}
			if err := sender(ws, msg); err != nil {
				return
			}
		}
	}()

	go func() {
		msg := ""
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			if err := receiver(ws, &msg); err != nil {
				return
			}
			broker.Publish(msg)
		}
	}()

	<-ctx.Done()

}
