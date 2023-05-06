package config

import (
	"QBot/configs"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	EnvLocal = "local"
	EnvDev   = "development"
	EnvProd  = "production"
)

var (
	Env   string
	DB    configs.DB
	Pixiv configs.Pixiv
)

func Init() {
	file, _ := os.Open(".env")
	b, _ := io.ReadAll(file)
	_ = file.Close()
	Env = string(b)
	fmt.Printf("Config->Init->Env:%+v\n", Env)

	Load("db", &DB)
	Load("pixiv", &Pixiv)
}

func Load(name string, obj interface{}) {
	path := filepath.Join("./configs", name+".json")

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	b, _ := io.ReadAll(file)
	_ = file.Close()
	err = json.Unmarshal(b, &obj)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config->Load->%v\n", name)
}

func IsLocal() bool {
	return Env == EnvLocal
}
