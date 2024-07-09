package websocket

import (
	"context"
	"fmt"
	"hidroponic/cmd/hidroponic/http/helpers/response"
	"hidroponic/internal/errors"
	"hidroponic/internal/models/sensor"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	sync.Mutex
	clients      map[*Client]bool
	runBoardcast bool
}

func New() *WebSocketService {
	return &WebSocketService{}
}

func (h *WebSocketService) register(client *Client) {
	h.Lock()
	defer h.Unlock()
	h.clients[client] = true
	if !h.runBoardcast {
		h.runBoardcast = true
	}
}

func (h *WebSocketService) unregister(client *Client) {
	h.Lock()
	defer h.Unlock()
	delete(h.clients, client)
	if len(h.clients) == 0 {
		h.runBoardcast = false
	}
}

func (h *WebSocketService) Broadcast(clientDataType sensor.DataType, data any) {
	for client := range h.clients {
		if clientDataType == client.clientDataType {
			client.Write(data)
		}
	}
}

func (cs *WebSocketService) Setup() (err error) {
	cs.clients = map[*Client]bool{}

	return
}

func (cs *WebSocketService) ServeHTTP(readHandler ReadHandlerType) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		clientDataType := r.URL.Query().Get("type")
		if clientDataType == "" {
			err := errors.ErrorInvalidRequestBody.New("params type can not be empty")
			response.WriterResponseError(w, err)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		defer conn.Close()
		conn.SetReadLimit(maxMessageSize)

		client := &Client{conn: conn, readHandler: readHandler, isShouldPing: true, clientDataType: sensor.DataType(clientDataType)}
		cs.register(client)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go client.ping(ctx)
		client.read()

		cs.unregister(client)
	}
}

const (
	writeWait      = 4 * time.Second
	readWait       = 5 * 60 * time.Second
	pingPeriod     = (readWait * 9 / 10)
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ReadHandlerType func(messageData []byte) *string

type Client struct {
	sync.Mutex
	conn           *websocket.Conn
	readHandler    ReadHandlerType
	isShouldPing   bool
	clientDataType sensor.DataType
}

func (c *Client) read() {
	for {
		c.conn.SetReadDeadline(time.Now().Add(readWait))

		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error(fmt.Sprintf("error: %v", err))
			}
			break
		}

		resp := c.readHandler(message)

		if resp != nil {
			if err := c.Write(*resp); err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					slog.Error(fmt.Sprintf("error: %v", err))
				}
				break
			}
			if c.isShouldPing {
				c.isShouldPing = false
			}
		}
	}
}

func (c *Client) ping(ctx context.Context) {
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(readWait)); return nil })

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if c.isShouldPing {
				c.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					slog.Error(fmt.Sprintln(err.Error()))
				}
			}
			c.isShouldPing = true
		}
	}
}

func (c *Client) Write(data any) error {
	c.Lock()
	defer c.Unlock()
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	err := c.conn.WriteJSON(data)
	if err != nil {
		return err
	}

	if c.isShouldPing {
		c.isShouldPing = false
	}

	return nil
}
