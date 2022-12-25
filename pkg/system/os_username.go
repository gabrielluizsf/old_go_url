package system

import (
  "os/user"
  "github.com/GabrielLuizSF/go_url/warning"
)

func myOSUsername(errmessage string)string{

  whoami , err := user.Current();
    warning.PRINT_DEFAULT_ERRORS(err,errmessage);
  username := whoami.Username;
  return username;
}
