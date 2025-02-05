#! /bin/bash to-git-or-not-to-git.sh
result=$(jq -r ".[] | select(.id == $superhero_id) | .name, .power, .gender" "$json_file")
if [[ -n "$result" ]]; then
  printf "$result" | tr '\n' ' ' # Print the result on a single line
  printf # Add a newline after the output
else
  printf "Superhero with ID $superhero_id not found." >&2
fi