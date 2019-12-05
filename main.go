package main

import (
        "encoding/json"
        "fmt"
        guuid "github.com/google/uuid"
        "io/ioutil"
        "log"
        "net/http"
        "strconv"
        "time"
)

const format = `[Interface]
PrivateKey= %s
Address = %s
DNS = 8.8.8.8

[Peer]
PublicKey = %s
Endpoint = %s
AllowedIPs = 0.0.0.0/0
`

type RespS struct {
        ServerPort       int
        ServerPublicKey  string
        ClientPrivateKey string

        PeerIP string
}

func main() {
        tmpclient := http.Client{
          Timeout: 10 * time.Second,
        }
        req, err := http.NewRequest("GET", "http://198.181.47.35:8801/v1/peers/config", nil)
        uuid := guuid.New()
        req.Header.Add("UUID", uuid.String())
        req.Header.Add("Authentication", "zhimakaimen")
        resp, err := tmpclient.Do(req)
        if err != nil {
          log.Fatal(err)
        }
        defer resp.Body.Close()

        responseData, err := ioutil.ReadAll(resp.Body)
        var respS RespS
        json.Unmarshal(responseData, &respS)
        if err != nil {
          log.Fatal(err)
        }
        dataString := fmt.Sprintf(format, respS.ClientPrivateKey, respS.PeerIP, respS.ServerPublicKey, "198.181.47.35:"+strconv.Itoa(respS.ServerPort))
        ioutil.WriteFile("./test1.conf", []byte(dataString), 0666)

}

