// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package settings

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mailchain/mailchain/cmd/mailchain/internal/settings/defaults"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus" // nolint: depguard
	"github.com/spf13/viper"         // nolint: depguard
)

// // MailchainHome set home directory for mailchain
// func MailchainHome() (string, error) {
// 	usr, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	p := usr.HomeDir + "/.mailchain"

// 	if _, err := os.Stat(p); os.IsNotExist(err) {
// 		_ = os.Mkdir(p, 0700)
// 	}
// 	return p
// }
func InitStore(v *viper.Viper, cfgFile, logLevel string, createFile bool) error {
	if cfgFile == "" {
		// working directory
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		cfgFile = filepath.Join(dir, defaults.ConfigFileName, defaults.ConfigFileName+"."+defaults.ConfigFileKind)
		// home directory
		// dir, err := homedir.Dir()
		// if err != nil {
		// 	return err
		// }
		// cfgFile = filepath.Join(dir, defaults.ConfigFileName+"."+defaults.ConfigFileKind)
	}
	lvl, err := log.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.Warningf("Invalid 'log-level' %q, default to [Warning]", logLevel)
		lvl = log.WarnLevel
	}
	log.SetLevel(lvl)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	v.SetEnvPrefix("mc")
	v.AutomaticEnv()

	v.SetConfigFile(cfgFile)

	err = v.ReadInConfig()
	_, ok := err.(viper.ConfigFileNotFoundError)
	if ok || err != nil && strings.Contains(err.Error(), "no such file or directory") {
		if createFile {
			return createEmptyFile(v, cfgFile)
		}
		return errors.WithMessage(err, "config creation disabled")
	}
	return err
}

func createEmptyFile(v *viper.Viper, fileName string) error {
	dir, err := filepath.Abs(filepath.Dir(fileName))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return v.WriteConfigAs(fileName)
}

// func WriteConfig(v *viper.Viper) func(cmd *cobra.Command, args []string) error {
// 	return func(cmd *cobra.Command, args []string) error {
// 		if err := v.WriteConfig(); err != nil {
// 			return errors.WithStack(err)
// 		}
// 		cmd.Printf(chalk.Green.Color("Config saved\n"))
// 		return nil
// 	}
// }
