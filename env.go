package todoapi

import (
	"os"
)

// 必要な環境変数を格納するための構造体
type Env struct {
	Bind      string
	MasterURL string
	SlaveURL  string
}

func CreateEnv() (*Env, error) {

	env := Env{}

	bind := os.Getenv("TODO_BIND") // APIをListenするポート設定
	if bind == "" {
		bind = ":8080"
	}
	env.Bind = bind

	masterURL := os.Getenv("TODO_MASTER_URL") // MySQL Masterへの接続情報
	if masterURL == "" {
		masterURL = "https://master.com"
		//test// return nil, errors.New("TODO_MASTER_URL is not specified")
	}
	env.MasterURL = masterURL

	slaveURL := os.Getenv("TODO_SLAVE_URL") // MySQL Slaveへの接続情報
	if slaveURL == "" {
		slaveURL = "https://master.com"
		//test// return nil, errors.New("TODO_SLAVE_URL is not specified")
	}
	env.SlaveURL = slaveURL

	return &env, nil
}
