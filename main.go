package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// EventType Klevr event type
type EventType string

// Klevr Event constraints
const (
	// 에이전트 접속 이벤트
	AgentConnect EventType = "AGENT_CONNECT"
	// 에이전트 접속 해제 이벤트
	AgentDisconnect EventType = "AGENT_DISCONNECT"
	// primary 에이전트 선출 이벤트
	PrimaryElected EventType = "PRIMARY_ELECTED"
	// primary 에이전트 retire 이벤트
	PrimaryRetire EventType = "PRIMARY_RETIRE"
	// task 수행 결과 전달 이벤트
	TaskCallback EventType = "TASK_CALLBACK"
)

type KlevrEvent struct {
	EventType EventType `json:"eventType"`
	AgentKey  string    `json:"agentKey"`
	GroupID   uint64    `json:"groupId"`
	EventTime *JSONTime `json:"eventTime"`
	Result    string    `json:"result"`
	Log       string    `json:"log"`
}

func main() {
	route := echo.New()

	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	route.POST("/events/klevr", EventHandler)

	if err := route.Start(fmt.Sprintf(":%d", 8080)); err != nil {
		panic(err)
	}
}

func EventHandler(ctx echo.Context) error {
	event := &[]KlevrEvent{}
	if err := ctx.Bind(event); err != nil {
		panic(err)
	}

	fmt.Println(event)

	return ctx.JSON(http.StatusOK, nil)
}
