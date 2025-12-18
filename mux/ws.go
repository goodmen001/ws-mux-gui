package main
import(
  "log"
  "net/http"
  "sync"
  "time"
  "github.com/gorilla/websocket"
)
type WSManager struct{
  url string
  header http.Header
  ws *websocket.Conn
  lock sync.Mutex
  cond *sync.Cond
}
func NewWSManager(url string,h http.Header)*WSManager{
  m:=&WSManager{url:url,header:h}
  m.cond=sync.NewCond(&m.lock)
  go m.loop(); return m
}
func(m *WSManager)loop(){
  for{
    ws,_,err:=websocket.DefaultDialer.Dial(m.url,m.header)
    if err!=nil{time.Sleep(3*time.Second);continue}
    m.lock.Lock(); m.ws=ws; m.cond.Broadcast(); m.lock.Unlock()
    for{ if _,_,err:=ws.ReadMessage(); err!=nil{break} }
    m.lock.Lock(); m.ws=nil; m.lock.Unlock()
  }
}
func(m *WSManager)Get()*websocket.Conn{
  m.lock.Lock(); defer m.lock.Unlock()
  for m.ws==nil{ m.cond.Wait() }
  return m.ws
}
