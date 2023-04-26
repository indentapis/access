package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"go.indent.com/indent-go/pkg/cliutil"
)

const (
	// setArgs is the number of arguments expected by the set command, representing the key and value.
	setArgs = 2
)

// NewCmdSet returns a command allowing configuration values to be set.
func NewCmdSet(f cliutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set configuration value",
		Long:  "Stores a key and value in the configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			if len(args) < setArgs {
				logger.Error("Expected two arguments: key and value")
				if err := cmd.Help(); err != nil {
					logger.Fatal("Failed to print help", zap.Error(err))
				}
				return
			}
			key, value := args[0], args[1]
			logger = logger.With(zap.String("key", key), zap.String("value", value))
			logger.Debug("Setting configuration value")
			viper.Set(key, value)
			f.WriteConfig()
		},
	}
	return cmd
}
