#!/bin/bash
curl https://adam-jerusalem.nd.edu/assets/superhero/all.json | jq --argjson id "$HERO_ID" '.[] | select(.id == $id) | .connections.relatives' | sed 's/^"\(.*\)"$/\1/'