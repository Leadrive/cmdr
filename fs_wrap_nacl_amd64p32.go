// +build nacl,amd64p32

/*
 * Copyright © 2019 Hedzr Yeh.
 */

package cmdr

import (
	"sync"
)

func fsWatcherRoutine(s *Options, configDir string, initWG *sync.WaitGroup) {
	initWG.Done() // done initializing the watch in this go routine, so the parent routine can move on...
}

// func fsWatchRunner(s *Options, configDir string, watcher *fsnotify.Watcher, eventsWG *sync.WaitGroup) {
// 	eventsWG.Done()
// }