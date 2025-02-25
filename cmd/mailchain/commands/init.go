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

package commands

// import (
// 	"fmt"
// 	"os"

// 	"github.com/mailchain/mailchain"
// 	"github.com/mailchain/mailchain/cmd/mailchain/internal/config"
// 	"github.com/mailchain/mailchain/cmd/mailchain/internal/config/defaults"
// 	"github.com/mailchain/mailchain/cmd/mailchain/internal/setup"
// 	"github.com/mailchain/mailchain/internal/chains"
// 	"github.com/mailchain/mailchain/internal/chains/ethereum"
// 	"github.com/manifoldco/promptui"
// 	"github.com/pkg/errors"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper" // nolint: depguard
// 	"github.com/ttacon/chalk"
// )

// func initCmd(viper *viper.Viper) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "init",
// 		Short: "Initialize mailchain configuration",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			cancel, err := ensureConfigFileRemoved(cmd, viper)
// 			if err != nil {
// 				return err
// 			}
// 			if cancel {
// 				return nil
// 			}
// 			// Configure default keystore
// 			// TODO: remove the defaults here
// 			viper.Set("storage.state", defaults.MailboxStateType)
// 			viper.Set(fmt.Sprintf("stores.%s.path", defaults.MailboxStateType), defaults.MailboxStatePath)

// 			viper.Set("server.port", defaults.Port)

// 			viper.Set("server.cors.allowed-origins", defaults.CORSAllowedOrigins)
// 			viper.Set("server.cors.disabled", defaults.CORSDisabled)

// 			for _, network := range ethereum.Networks() {
// 				if _, err := setup.DefaultNetwork().Select(cmd, args, chains.Ethereum, network); err != nil {
// 					return err
// 				}
// 			}

// 			if _, err := setup.DefaultSentStorage().Select(mailchain.Mailchain); err != nil {
// 				return err
// 			}

// 			if _, err := setup.DefaultKeystore().Select(mailchain.RequiresValue, mailchain.RequiresValue); err != nil {
// 				return err
// 			}

// 			viper.SetConfigFile(defaults.ConfigFile)
// 			if err := viper.WriteConfig(); err != nil {
// 				return err
// 			}

// 			fmt.Println(chalk.Green, "Config created: ", chalk.White, viper.ConfigFileUsed())
// 			return nil
// 		},
// 	}
// }

// func ensureConfigFileRemoved(cmd *cobra.Command, v *viper.Viper) (cancel bool, err error) {
// 	cfgFile, _ := cmd.PersistentFlags().GetString("config")
// 	logLevel, _ := cmd.PersistentFlags().GetString("log-level")

// 	switch e := config.Init(v, cfgFile, logLevel).(type) {
// 	case viper.ConfigFileNotFoundError:
// 		// Do nothing
// 	case nil:
// 		fmt.Println(chalk.Red, "Config already exists: ", chalk.White, viper.ConfigFileUsed())
// 		fmt.Println("Remove this file first and re-run this command. To edit an existing file use `mailchain config`")
// 		fmt.Println("By continuing it will delete your existing mailchain configuration")
// 		prompt := promptui.Prompt{
// 			Label:     "Continue",
// 			Default:   "n",
// 			IsConfirm: true,
// 		}
// 		_, err := prompt.Run()
// 		if err == promptui.ErrAbort {
// 			return true, nil
// 		}
// 		if err != nil {
// 			return false, errors.WithMessage(err, "can not confirm")
// 		}
// 		if err := os.Remove(viper.ConfigFileUsed()); err != nil {
// 			return false, errors.WithMessage(err, "failed to remove existing config")
// 		}
// 		viper.Reset()

// 		err = config.Init(v, cfgFile, logLevel)
// 		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
// 			return false, errors.WithMessage(err, "failed to re-init config")
// 		}

// 	default:
// 		return false, e
// 	}
// 	if err := os.MkdirAll(defaults.ConfigPathFirst, os.ModePerm); err != nil {
// 		return false, err
// 	}

// 	return false, nil
// }
