package main

import (
	"fmt"
	"strconv"

	"github.com/ahmetcanozcan/reves"
	"github.com/ahmetcanozcan/reves/engines"
	"github.com/ahmetcanozcan/reves/sockets"
	"github.com/ahmetcanozcan/reves/sockets/messages"
)

func main() {

	reves.Config.PORT = "8080"

	engines.OnPlayerInvoked(func(p *engines.Player) {
		fmt.Println("a Player invoked")
		p.GetSocket().On("Foo", func(msg messages.Payload) {
			fmt.Println("Hello from player", p.GetID())
		})
	})

	sockets.WhenMatchmakingRoomIsFilled(func(r *sockets.Room) {
		fmt.Println("a MatchMaking room is filled...")
		engine := engines.NewEngine(r)
		engine.AddEntity(&Clock{time: 0})
		engine.Start()
	})

	reves.HandleNewSocket(func(socket *sockets.Socket) {
		socket.On("Hello", func(payload messages.Payload) {
			for k, v := range payload {
				fmt.Println(k, v)
			}
		})
	})

	reves.Start()

}

//Clock : Test entity
type Clock struct {
	time int64
}

//Start :
func (c *Clock) Start() {}

//Check :
func (c *Clock) Check() {}

//Update :
func (c *Clock) Update(deltaTime int64) { c.time += deltaTime }

//Share :
func (c *Clock) Share(players engines.PlayerList) {
	payload := messages.NewPayload()
	payload["Time"] = strconv.FormatInt(c.time, 10)
	for _, player := range players {
		player.GetSocket().Emit("Clock", payload)
	}
}
