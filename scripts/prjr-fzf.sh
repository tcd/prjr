pjcd() {
  local destination

  destination=$(prjr list | fzf -0 -1 | awk '{print $2}')
  cd "$destination"
}
