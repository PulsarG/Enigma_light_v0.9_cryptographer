/*
__      _________      __
\ \    / /_   _\ \    / /\
 \ \  / /  | |  \ \  / /  \
  \ \/ /   | |   \ \/ / /\ \
   \  /   _| |_   \  / ____ \
 _  \/   |_____|  _\/_/    \_\     _   _
| | | |          |  _ \           | | | |
| |_| |__   ___  | |_) | ___  __ _| |_| | ___  ___
| __| '_ \ / _ \ |  _ < / _ \/ _` | __| |/ _ \/ __|
| |_| | | |  __/ | |_) |  __/ (_| | |_| |  __/\__ \
 \__|_| |_|\___| |____/ \___|\__,_|\__|_|\___||___/


*/

package main

import (
	"fyne.io/fyne/v2/app"

	"enigma/crypt"
)

func main() {
	app := app.New()

	Cryptor := crypt.NewCryptor(app)
	Cryptor.Start()

	app.Run()
}
