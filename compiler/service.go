package compiler

import (
	"fmt"
)

/*
  Provide service related stuff
*/

func ServiceValidate(i interface{}) error {
	switch i.(type) {
	case map[interface{}]interface{}:
		return nil
	default:
		return fmt.Errorf("service: this runner expects a set of 'string:string' kvs. Got %T", i)
	}

	missingKeys := []string{}
	for _, k := range []string{"mode", "service"} {
		_, ok := i.(map[interface{}]interface{})[k]
		if !ok {
			missingKeys = append(missingKeys, k)
		}
	}

	if len(missingKeys) > 0 {
		return fmt.Errorf("service: this has missing 'args' value(s): %q", missingKeys)
	}

	return nil
}

func ServicePreamble() string {
	return `
_srv() {
  local operation="${1}"
  local service="${2}"

  if [ $(which rc-service) ] ; then
    case "${operation}" in
      "enable")
        rc-update add "${service}" default
        ;;
      *)
        rc-service "${service}" "${operation}"
    esac

  else
    err "service: no support for this distro" "${E_BAD_DISTRO}"
  fi
}
`
}

func ServiceCommand(i interface{}) string {
	args := i.(map[interface{}]interface{})
	return fmt.Sprintf("_srv %q %q || err 'service' ${E_BAD_CMD}\n", args["mode"], args["service"])
}
