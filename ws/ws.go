package ws

import (
	"context"
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"gopkg.in/olahol/melody.v1"
)

const userSessionKey = "user_session"

// UserSession session
type UserSession struct {
	ID string
}

// Server server
type Server struct {
	*melody.Melody
	*sync.Mutex
	*echo.Echo
	UserSession map[string]*melody.Session
}

var instance *Server

// New init
func New() *Server {
	instance = &Server{
		Melody:      melody.New(),
		Mutex:       new(sync.Mutex),
		Echo:        echo.New(),
		UserSession: make(map[string]*melody.Session),
	}

	instance.File("/", "index.html")

	instance.GET("/ws", instance.handleRequest)

	instance.HandleMessage(instance.handleMessage)
	instance.HandleConnect(instance.handleConnect)
	instance.HandleDisconnect(instance.handleDisconnect)
	return instance

}

// SendMessage implement interface
func (s *Server) SendMessage(ctx context.Context, in *Message) (*Message, error) {
	var userID = in.GetUserId()
	if userID == "" {
		return nil, fmt.Errorf("User ID is required")
	}

	fmt.Printf("Handle message from user_id = %s\n", userID)

	return in, nil
}

func (s *Server) handleRequest(c echo.Context) error {
	var err = s.HandleRequest(c.Response(), c.Request())
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) handleConnect(session *melody.Session) {
	var token = session.Request.URL.Query().Get("token")
	claims, err := verifyToken("asdf", token)
	if err != nil {
		session.Close()
		return
	}

	session.Set(userSessionKey, &UserSession{
		ID: claims.ID,
	})
	s.UserSession[claims.ID] = session
}

func (s *Server) handleDisconnect(session *melody.Session) {
	if value, found := session.Get("user_id"); found {

		if userSession, ok := value.(*UserSession); ok {
			fmt.Printf("User %s was disconnected\n", userSession.ID)
			delete(s.UserSession, userSession.ID)

		}

	}

}

func (s *Server) handleMessage(session *melody.Session, msg []byte) {
	var userSession *UserSession
	if value, found := session.Get("user_id"); found {
		var ok = false
		if userSession, ok = value.(*UserSession); ok {

		}
	}
	if userSession == nil {
		return
	}

	fmt.Printf("Handle message from user_id = %s messsage = %s\n", userSession.ID, string(msg))
	s.Broadcast(msg)
}
