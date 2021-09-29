package main

import (
    "github.com/spf13/cobra"
    "github.com/sqmt/rich-cli/cmd"
)

//func init() {
//    cmd.InitCobra(func() {
//        fmt.Println("initialize")
//    })
//
//    cmd.Apply(func(cmd *cobra.Command) {
//        cmd.Use = "test"
//        cmd.Version = "1.2"
//    })
//}
func main() {
    cobra.CheckErr(cmd.Execute())
}
