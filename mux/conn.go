package main
import(
  "bytes"
  "net"
  "sync"
  "time"
  "github.com/gorilla/websocket"
)
type MuxConn struct{
  id uint32
  ws *websocket.Conn
  buf bytes.Buffer
  mu sync.Mutex
}
func NewMuxConn(ws *websocket.Conn,id uint32)*MuxConn{
  return &MuxConn{id:id,ws:ws}
}
func(c *MuxConn)Read(p []byte)(int,error){
  for{
    c.mu.Lock()
    if c.buf.Len()>0{
      n,_:=c.buf.Read(p); c.mu.Unlock(); return n,nil
    }
    c.mu.Unlock(); time.Sleep(5*time.Millisecond)
  }
}
func(c *MuxConn)Write(p []byte)(int,error){
  err:=c.ws.WriteMessage(websocket.BinaryMessage,p)
  return len(p),err
}
func(c *MuxConn)Close()error{return nil}
func(c *MuxConn)LocalAddr()net.Addr{return nil}
func(c *MuxConn)RemoteAddr()net.Addr{return nil}
func(c *MuxConn)SetDeadline(t time.Time)error{return nil}
func(c *MuxConn)SetReadDeadline(t time.Time)error{return nil}
func(c *MuxConn)SetWriteDeadline(t time.Time)error{return nil}
