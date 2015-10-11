package twitter

import (
	"github.com/joeshaw/envdecode"
	"io"
	"log"
	"net"
	"time"
)

var (
	conn       net.Conn
	reader     io.ReadCloser
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func setupTwitterAuth() {
	var autinfo struct {
		ConsumerKey    string `env:"PTG_TWITTER_API_KEY,requited"`
		ConsumerSecret string `env:"PTG_TWITTER_API_SECRET,requited"`
		AccessToken    string `env:"PTG_TWITTER_ACCESS_TOKEN,requited"`
		AccessSecret   string `env:"PTG_TWITTER_ACCESS_TOKEN_SECRET,requited"`
	}

	if err := envencode.Decode(&authinfo); err != nil {
		log.Fatalln(err)
	}

	creds = &oauth.Credentials{
		Token:  authinfo.AccessToken,
		Secret: authinfo.AccessSecret,
	}

	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  authinfo.ConsumerKey,
			Secret: authinfo.ConsumerSecret,
		},
	}

}

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil { //close the connection and set it to nil if it's not closed properly before
		conn.Close()
		conn = nil
	}

	netc, err := net.DialTimeout(netw, addr, 3*time.Second)

	if err != nil { //if error then return nil connection and error
		return nil, err
	}

	conn = netc
	return netc, nil
}

func closeConn() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		reader.Close()
	}

}
