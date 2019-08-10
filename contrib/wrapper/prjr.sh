#!/usr/bin/env sh

# Check if one string contains another.
stringContain() {
  [ -z "${2##*$1*}" ] && [ -z "$1" -o -n "$2" ];
}

# This wrapper is necessary because we can't change the shell's working
# directory from inside a program.
main() {
  if [ "$1" = "jump" ]; then
    local cmd_output=$(prjr "$@")
    if stringContain 'jumpto:' "$cmd_output"; then
      local dest=$(echo "${cmd_output}" | sed 's/.*: \(.*\)$/\1/' | sed -e 's/jumpto://g')
      cd "${dest}" || return
    else
      echo "$cmd_output" && return
    fi
  else
    prjr "$@" && return
  fi
}

main "$@"
