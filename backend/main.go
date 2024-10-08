package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/watermelo/realtime-chat-go-react/pkg/websocket"
)

func serveWs(pool *websocket.Pool, c *gin.Context) {
    fmt.Println("WebSocket Endpoint Hit")

    // Upgrade this connection to a WebSocket
    conn, err := websocket.Upgrade(c.Writer, c.Request)
    if err != nil {
        fmt.Fprintf(c.Writer, "%+v\n", err)
        return
    }

    // Create a new WebSocket client
    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client

    // Start listening for incoming chat messages
    client.Read()
}

func setupRoutes(router *gin.Engine) {
    // Configure websocket route
    pool := websocket.NewPool()
    go pool.Start()

    // websocket handler
    router.GET("/ws", func(c *gin.Context) {
        serveWs(pool, c)
    })
}

func main() {
    fmt.Println("Distributed Chat App v0.01")

    // Set the router as the default one shipped with Gin
    gin.SetMode(gin.ReleaseMode)

    // Creates a router without any middleware by default
    router := gin.Default()
    
    // Setup routes
    setupRoutes(router)
    
    // Start and run the server
    router.Run(":8080")
}
