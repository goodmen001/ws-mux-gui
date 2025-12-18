package main
import(
  "crypto/hmac"
  "crypto/sha256"
  "encoding/base64"
  "fmt"
  "time"
)
func genAuth(token, secret string) map[string]string{
  ts:=fmt.Sprintf("%d",time.Now().Unix())
  mac:=hmac.New(sha256.New,[]byte(secret))
  mac.Write([]byte(token+"|"+ts))
  sig:=base64.StdEncoding.EncodeToString(mac.Sum(nil))
  return map[string]string{"x-token":token,"x-ts":ts,"x-sig":sig}
}
