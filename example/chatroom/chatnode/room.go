package chatnode

import (
	"github.com/helloh2o/lucky/core/iduck"
	"github.com/helloh2o/lucky/core/inet"
)

var testChatRoom iduck.INode

func GetRoom() iduck.INode {
	if testChatRoom == nil {
		testChatRoom = inet.NewBroadcastNode()
		testChatRoom.Serve()
	}
	return testChatRoom
}
