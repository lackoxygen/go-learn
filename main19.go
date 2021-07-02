package main

import "fmt"

type MysqlInterface interface {
	setHost() MysqlInterface
	setPort() MysqlInterface
	setUser() MysqlInterface
	setPass() MysqlInterface
}

type Mysql struct {
	host string
	port int
	user string
	pass string
}

func (m *Mysql) setHost(host string) *Mysql {
	m.pass = host
	return m
}

func (m *Mysql) setPort(port int) *Mysql {
	m.port = port
	return m
}

func (m *Mysql) setUser(user string) *Mysql {
	m.user = user
	return m
}

func (m *Mysql) setPass(pass string) *Mysql {
	m.pass = pass
	return m
}

func NewMysql(host string, port int, user string, pass string) *Mysql {
	m := &Mysql{}
	m.setHost(host).setPort(port).setUser(user).setPass(pass)
	return m
}

func main() {
	m := NewMysql("127.0.0.1", 3306, "root", "123456")

	fmt.Println(*m)
}
