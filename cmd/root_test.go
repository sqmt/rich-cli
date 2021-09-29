package cmd

import (
    "github.com/spf13/cobra"
    "testing"
)

func reset() {
    rootCmd.Use = "rich-cli"
    rootCmd.Version = ""
    rootCmd.RemoveCommand(rootCmd.Commands()...)
}

func TestAddCommand(t *testing.T) {
    defer reset()
    type args struct {
        commands []*cobra.Command
    }
    tests := []struct {
        name string
        args args
        hasSub bool
    }{
        {name: "none", args: args{}, hasSub: false},
        {name: "add", args: args{commands: []*cobra.Command{&cobra.Command{Use: "test"}}}, hasSub: true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            AddCommand(tt.args.commands...)
            if rootCmd.HasSubCommands() != tt.hasSub {
                t.Errorf("AddCommand failed")
            }
        })
    }
}

func TestApply(t *testing.T) {
    defer reset()
    type args struct {
        f func(cmd *cobra.Command)
    }
    tests := []struct {
        name     string
        args     args
        wantBool func(cmd *cobra.Command) bool
    }{
        {name: "change use", args: args{f: func(cmd *cobra.Command) {
            cmd.Use = "test"
        }}, wantBool: func(cmd *cobra.Command) bool {
            return cmd.Use == "test"
        }},
        {name: "add version", args: args{f: func(cmd *cobra.Command) {
            cmd.Version = "0001"
        }}, wantBool: func(cmd *cobra.Command) bool {
            return cmd.Version == "0001"
        }},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            Apply(tt.args.f)
            if !tt.wantBool(rootCmd) {
                t.Errorf("Apply() failed")
            }
        })
    }
}

func TestExecute(t *testing.T) {
    defer reset()
    type args struct {
        args []string
        apply func(cmd *cobra.Command)
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {name: "show version1", args: args{args: []string{"-v"}}, wantErr: true},
        {name: "show version2", args: args{args: []string{"-v"}, apply: func(cmd *cobra.Command) {
            cmd.Version = "0.0.1"
        }}, wantErr: false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.args.apply != nil {
                Apply(tt.args.apply)
            }
            if err := Execute(tt.args.args...); (err != nil) != tt.wantErr {
                t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestInitCobra(t *testing.T) {
    defer reset()
    var v int
    type args struct {
        f []func()
        v interface{}
    }
    tests := []struct {
        name string
        args args
    }{
        {name:"init value", args: args{
            f: []func(){func() {
                v = 1
            }},
            v: 1,
        }},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            InitCobra(tt.args.f...)
            if err := Execute(); err != nil {
                t.Errorf("Execute faild, %v", err)
            }
            if v != tt.args.v {
                t.Errorf("InitCobra faild, go %v, want %v", v, tt.args.v)
            }
        })
    }
}

func TestRemoveCommand(t *testing.T) {
    defer reset()
    type args struct {
        commands []*cobra.Command
    }
    tests := []struct {
        name string
        args args
        hasSub bool
    }{
        {name: "has", args: args{commands: []*cobra.Command{&cobra.Command{Use: "test"}}}, hasSub: false},
        {name: "none", args: args{}, hasSub: false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            RemoveCommand(tt.args.commands...)
            if rootCmd.HasSubCommands() != tt.hasSub {
                t.Errorf("RemoveCommand failed")
            }
        })
    }
}
