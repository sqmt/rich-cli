package cmd

import (
    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:              "rich-cli",
    TraverseChildren: true,
    SilenceErrors:    true,
    SilenceUsage:     true,
    RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Help()
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(args ...string) error {
    if len(args) > 0 {
        rootCmd.SetArgs(args)
    }
    return rootCmd.Execute()
}

// InitCobra cobra初始化配置处理，需要在调用 Execute 之前使用
func InitCobra(f ...func()) {
    cobra.OnInitialize(f...)
}

// Apply 对rootCmd进行赋值
func Apply(f func(cmd *cobra.Command)) {
    if f != nil {
        f(rootCmd)
    }
}

// AddCommand adds one or more commands to this parent command.
func AddCommand(commands ...*cobra.Command) {
    rootCmd.AddCommand(commands...)
}

// RemoveCommand removes one or more commands from a parent command.
func RemoveCommand(commands ...*cobra.Command) {
    rootCmd.RemoveCommand(commands...)
}
