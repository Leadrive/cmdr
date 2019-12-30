// Copyright © 2019 Hedzr Yeh.

package shell

import "github.com/hedzr/cmdr"

func WithShellModule() cmdr.ExecOption {
	return func(w *cmdr.ExecWorker) {
		// daemonImpl = daemonImplX

		w.AddOnBeforeXrefBuilding(func(root *cmdr.RootCommand, args []string) {

			// if modifier != nil {
			// 	root.SubCommands = append(root.SubCommands, modifier(DaemonServerCommand))
			// } else {
			// 	root.SubCommands = append(root.SubCommands, DaemonServerCommand)
			// }
			// 
			// // prefix = strings.Join(append(cmdr.RxxtPrefix, "server"), ".")
			// // prefix = "server"
			// 
			// attachPreAction(root, preAction)
			// attachPostAction(root, postAction)

			rootOpt := cmdr.NewCmdFrom(&root.Command)
			rootOpt.NewSubCommand().
				Titles("sh", "shell").
				Action(doAction)
		})
	}
}

func doAction(cmd *cmdr.Command, args []string) (err error) {
	return
}

// 
func AttachToCmdr(parent *cmdr.Command) {

}
