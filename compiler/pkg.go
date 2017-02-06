package compiler

import (
	"fmt"
)

/*
  Install packages on specific operating systems
*/

func PKGValidate(i interface{}) error {
	// pkg expects a package name only
	switch i.(type) {
	case string:
		return nil
	default:
		return fmt.Errorf("pkg: this runner expects a package name as it's sole argument")
	}
	return fmt.Errorf("pkg: an unexpected error occurred")
}

func PKGPreamble() string {
	return `
_pkg() {
  local package="${1}"

  if [ $(which apk) ] ; then
    [ "${PACKAGE_MANAGER_UPDATE}" -eq 0 ] && ( apk update && export PACKAGE_MANAGER_UPDATE=1 )
    apk add $package
  else
    err "service: no support for this distro" "${E_BAD_DISTRO}"
  fi
}
`
}

func PKGCommand(i interface{}) string {
	return fmt.Sprintf("_pkg %q || err 'pkg' ${E_BAD_CMD}\n", i.(string))
}
