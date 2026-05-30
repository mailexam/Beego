package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	beego "github.com/beego/beego/v2/server/web"
)

type sendRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Text    string `json:"text"`
}

type MailController struct {
	beego.Controller
}

func (c *MailController) PostTest() {
	var req sendRequest
	if len(c.Ctx.Input.RequestBody) > 0 {
		_ = json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	}

	body := req.Body
	if body == "" {
		body = req.Text
	}

	if err := sendTest(req.To, req.Subject, body); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"status": "ok"}
	c.ServeJSON()
}

func listenHost() string {
	if host := os.Getenv("HTTP_HOST"); host != "" {
		return host
	}
	return "127.0.0.1"
}

func listenPort() int {
	if value := os.Getenv("HTTP_PORT"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return 8080
}

func main() {
	_ = godotenv.Load()

	beego.BConfig.Listen.HTTPAddr = listenHost()
	beego.BConfig.Listen.HTTPPort = listenPort()

	beego.Router("/mail/test", &MailController{}, "post:PostTest")

	beego.Run()
}
