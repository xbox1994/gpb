package config

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

import _ "github.com/spf13/viper/remote"

func Init()  {
	configUrl := os.Getenv("CONFIG_URL")
	if configUrl == "" {
		return
	}
	//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	//file://FILE_PWD/config/package.json
	//etcd://wafdsafdsaf:2379/config.json
	//consul://eafea.fdasom.com/eawfea
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(configUrl)
	if err != nil {
		panic(err)
	}
	//直接访问 scheme。
	if u.Scheme != "file" && u.Scheme != "etcd" && u.Scheme != "consul" {
		panic("config scheme must be file, etcd or consul")
	}
	if u.Scheme == "file" {
		viper.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(u.Path)), "."))
		viper.SetConfigName(strings.TrimSuffix(path.Base(u.Path), path.Ext(path.Base(u.Path))))
		if u.Host == "FILE_PWD" {
			viper.AddConfigPath("." + path.Dir(u.Path))
		}else{
			viper.AddConfigPath(path.Dir(u.Path))
		}
		err = viper.ReadInConfig() // Find and read the config file
		if err != nil {
			panic(fmt.Errorf("unable to read etcd config: %v", err))
		}
		return
	}
	if u.Scheme == "etcd" {
		viper.AddRemoteProvider("etcd", "http://" + u.Host, u.Path)
		viper.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(u.Path)),".")) // Need to explicitly set this to json
		err = viper.ReadRemoteConfig()

		viper.WatchRemoteConfig()
		cfg := client.Config{
			Endpoints:               []string{"http://" + u.Host},
			Transport:               client.DefaultTransport,
			// set timeout per request to fail fast when the target endpoint is unavailable
			HeaderTimeoutPerRequest: time.Second,
		}

		c, err := client.New(cfg)
		if err != nil {
			panic(fmt.Errorf("unable to read etcd config: %v", err))
		}

		kApi := client.NewKeysAPI(c)
		// open a goroutine to watch remote changes forever
		go func(){
			for {
				// currently, only tested with etcd support
				watcher := kApi.Watcher(u.Path, &client.WatcherOptions{Recursive: false})
				res, err := watcher.Next(context.Background())
				if err != nil {
					panic(fmt.Errorf("unable to watch etcd config: %v", err))
				}
				fmt.Print(res)
				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				//runtime_viper.Unmarshal(&runtime_conf)
			}
		}()

		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		return
	}
	if u.Scheme == "consul" {
		viper.AddRemoteProvider("consul", "http://" + u.Host, u.Path)
		viper.SetConfigType(strings.TrimPrefix(path.Ext(path.Base(u.Path)),".")) // Need to explicitly set this to json
		err = viper.ReadRemoteConfig()

		viper.WatchRemoteConfig()

		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		return
	}
	return
}
