package main

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/pelletier/go-toml"
	"github.com/thoj/go-ircevent"
	"strings"
)

type Plugin struct {
	ping        bool
	isup        bool
	stock       bool
	source      bool
	startups    bool
	moustachify bool
	isnsfw      bool
}

type Config struct {
	schar    string
	nick     string
	server   string
	autojoin string
}

var plugin Plugin
var config Config

func sendmsg(con *irc.Connection, msg string) {
	log.Infof("Output: %s", msg)
	con.Privmsg(config.autojoin, msg)
}

func iscommand(msg string) bool {
	if strings.HasPrefix(strings.ToLower(msg), config.schar) {
		return true
	}

	return false
}

func checkcommand(command string, msg string) bool {
	return strings.HasPrefix(command, msg)
}

func cleantoken(cmd string, msg string) string {
	a := strings.TrimPrefix(cmd, msg)
	return strings.TrimSpace(a)
}

func handlers(e *irc.Event, con *irc.Connection, schar string) {
	msg := e.Message()
	if iscommand(msg) {
		cmd := strings.TrimPrefix(msg, schar)
		fmt.Println("command is: " + cmd)

		if plugin.ping && checkcommand(cmd, "ping") {
			log.Info("Received ping")
			sendmsg(con, "pong")
		}

		if plugin.isup && checkcommand(cmd, "isup") {
			host := cleantoken(cmd, "isup")
			up, err := isup(host)
			if err != nil {
				sendmsg(con, "That's not a valid host!")
			} else {
				if up {
					sendmsg(con, "Seems to be up")
				} else {
					sendmsg(con, "It's down for me!")
				}
			}
		}

		if plugin.stock && checkcommand(cmd, "stock") {
			symbol := cleantoken(cmd, "stock")
			values := stock(symbol)

			sendmsg(con, values)
		}

		if plugin.source && checkcommand(cmd, "source") {
			sendmsg(con, "http://github.com/dysoco/dybot")
		}

		if plugin.startups && checkcommand(cmd, "startups") {
			log.Info("Startups has been called")
			sendmsg(con, startups())
		}
	}

	if plugin.moustachify && ispicture(msg) && isvalidhost(msg) {
		sendmsg(con, "Moustachified! "+moustachify(extracturl(msg)))
	}

	if plugin.isnsfw && ispicture(msg) && isvalidhost(msg) {
		nsfw, err := isnsfw(extracturl(msg))

		if err == nil {
			if nsfw {
				sendmsg(con, "Seems to be NSFW!")
			} else {
				sendmsg(con, "Looks safe for work")
			}
		}
	}

}

func loadcfg(path string) {
	file, _ := toml.LoadFile(path)

	config.schar = file.Get("schar").(string)
	config.nick = file.Get("nick").(string)
	config.server = file.Get("server").(string)
	config.autojoin = file.Get("autojoin").(string)

	plugin.ping = file.Get("plugin.ping").(bool)
	plugin.isup = file.Get("plugin.isup").(bool)
	plugin.stock = file.Get("plugin.stock").(bool)
	plugin.source = file.Get("plugin.source").(bool)
	plugin.startups = file.Get("plugin.startups").(bool)
	plugin.moustachify = file.Get("plugin.moustachify").(bool)
	plugin.isnsfw = file.Get("plugin.isnsfw").(bool)
}

func main() {
	loadcfg("config.toml")

	con := irc.IRC(config.nick, config.nick)
	err := con.Connect(config.server)
	if err != nil {
		fmt.Println("ERROR CONNECTING")
	}

	con.AddCallback("001", func(e *irc.Event) {
		con.Join(config.autojoin)
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		go handlers(e, con, config.schar)
	})

	con.Loop()
}
