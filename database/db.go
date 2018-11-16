package database

import (
	"database/sql"
	"io"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v2"
)

// Configs 環境ごとの設定を持つ
type Configs map[string]*Config

// Open 指定された環境に接続する
func (cs Configs) Open(env string) (*sql.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, nil
	}
	return config.Open()
}

// Config 設定
type Config struct {
	Datasource string `yaml:"datasource"`
}

// DSN 設定されているDSNを返す
func (c *Config) DSN() string {
	return c.Datasource
}

// Open DBに対して接続する
// MySQL固定
func (c *Config) Open() (*sql.DB, error) {
	return sql.Open("mysql", c.DSN())
}

// NewConfigsFromFile Configから設定を読み取る
func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

// NewConfigs io.Readerから設定を読み取る
func NewConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}
