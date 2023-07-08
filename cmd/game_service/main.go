package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type GameServer struct {
	clients     map[string]net.Conn //用户存储客户端的连接
	matchmaking chan string         //匹配队列
	clientsLock sync.RWMutex
}

func NewGameServer() *GameServer {
	return &GameServer{
		clients:     make(map[string]net.Conn),
		matchmaking: make(chan string),
	}
}

// 处理客户端连接
func (gs *GameServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	clientID, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取客户端ID失败")
		return
	}

	clientID = strings.TrimSpace(clientID)
	fmt.Println("客户端连接成功ID:", clientID)
	// 验证客户端身份
	if !gs.Authenticate(clientID) {
		log.Printf("客户端身份验证失败: %s\n", clientID)
		return
	}

	gs.clientsLock.Lock()
	gs.clients[clientID] = conn
	gs.clientsLock.Unlock()
	gs.matchmaking <- clientID
}

func (gs *GameServer) Authenticate(clientID string) bool {
	// TODO: 根据实际需求进行客户端身份验证
	// 这里只是简单示例，始终返回true
	return true
}
