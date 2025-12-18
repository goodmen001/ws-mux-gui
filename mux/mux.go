package mux
import(
  "fmt"
  "net"
  "sync"
)
type Mux struct{
  mgr *WSManager
  next uint32
  mu sync.Mutex
}
func NewMux(m *WSManager)*Mux{ return &Mux{mgr:m} }
func(m *Mux)Dial(addr string)(net.Conn,error){
  m.mu.Lock(); id:=m.next; m.next++; m.mu.Unlock()
  ws:=m.mgr.Get()
  ws.WriteMessage(1,[]byte(fmt.Sprintf("C %d %s",id,addr)))
  return NewMuxConn(ws,id),nil
}
