package proxy

import "testing"

func TestSetProxyForWin(t *testing.T) {
	SetProxyForWin("127.0.0.1:10001")
}

func TestCleanProxy(t *testing.T) {
	CleanProxy()
}
