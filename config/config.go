/*
如何添加配置

1. 修改 config.yaml 添加新的配置，的名称、默认值、以及必要的配置说明
2. 修改下面的 Config struct 添加新配置
3. 修改下面 mapConfig func 将最终的配置值拷贝到 Config struct
*/
package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	_ "embed"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var defaultConfig []byte

type Config struct {
	raw *viper.Viper

	ListenAddr string `yaml:"listen_addr" json:"listen_addr"`
	SystemKey  string `yaml:"system_key" json:"system_key"`

	CosConfig CosConfig `yaml:"cos" json:"cos"`
	ObsConfig ObsConfig `yaml:"obs" json:"obs"`
}

type CosConfig struct {
	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	IsDebug   bool   `json:"is_debug"`
}
type ObsConfig struct {
	Ak       string `json:"ak"`
	Sk       string `json:"sk"`
	Endpoint string `json:"endpoint"`
	Bucket   string `json:"bucket"`
}

func NewConfig(envPrefix string) Config {
	raw := viper.New()
	raw.SetConfigType("yaml")
	raw.SetEnvPrefix(envPrefix)
	raw.AutomaticEnv()
	raw.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := raw.MergeConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		log.Fatal().Err(err).Msg("Cannot load default config")
	}

	c := Config{
		raw: raw,
	}

	if err := c.mapConfig(); err != nil {
		log.Fatal().Err(err).Msg("Cannot map default config")
	}
	return c
}

func (c *Config) mapConfig() error {
	c.ListenAddr = c.raw.GetString("listen_addr")
	c.SystemKey = c.raw.GetString("system_key")

	c.ObsConfig.Ak = c.raw.GetString("obs.ak")
	c.ObsConfig.Sk = c.raw.GetString("obs.sk")
	c.ObsConfig.Endpoint = c.raw.GetString("obs.endpoint")
	c.ObsConfig.Bucket = c.raw.GetString("obs.bucket")

	c.CosConfig.SecretId = c.raw.GetString("cos.secret_id")
	c.CosConfig.SecretKey = c.raw.GetString("cos.secret_key")
	c.CosConfig.Bucket = c.raw.GetString("cos.bucket")
	c.CosConfig.Region = c.raw.GetString("cos.region")
	c.CosConfig.IsDebug = c.raw.GetBool("cos.is_debug")

	return nil
}

func (c *Config) mergeConfig(r io.Reader) error {
	if err := c.raw.MergeConfig(r); err != nil {
		return err
	}
	c.mapConfig()
	return nil
}

func (c *Config) ReadConfigFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("faile to open file filename=%s, err=%s", filename, err)
	}
	if err := c.mergeConfig(f); err != nil {
		return fmt.Errorf("failed to read config filename=%s, err=%s", filename, err)
	}

	return nil
}

func DefaultTestConfig() Config {
	raw := viper.New()
	raw.SetConfigType("yaml")
	raw.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := raw.MergeConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		log.Fatal().Err(err).Msg("Cannot load default config")
	}

	c := Config{
		raw: raw,
	}

	if err := c.mapConfig(); err != nil {
		log.Fatal().Err(err).Msg("Cannot map default config")
	}
	return c
}
