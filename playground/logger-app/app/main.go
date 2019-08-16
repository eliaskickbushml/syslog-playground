package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/syslog"
	"net/http"
	"os"
)

var (
	SYSLOG_SERVER string
	SYSLOG_PORT string
)

func main(){
	SYSLOG_SERVER = os.Getenv("SYSLOG_SERVER")
	SYSLOG_PORT = os.Getenv("SYSLOG_PORT")
	r := gin.Default()
	r.GET("/sendlog", SendLog)
	r.GET("/ping", Ping)
	r.Run(":5141")
}

func SendLog(c *gin.Context){
	msg := c.Param("msg")
	fmt.Printf("%s\n", fmt.Sprintf("%s:%s", SYSLOG_SERVER, SYSLOG_PORT))
	lg, err := syslog.Dial("tcp",fmt.Sprintf("%s:%s", SYSLOG_SERVER, SYSLOG_PORT),syslog.LOG_DEBUG | syslog.LOG_LOCAL0, "Bee")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	err = lg.Info(msg)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}

func Ping(c *gin.Context){
	c.String(http.StatusOK, "PONG")
}
