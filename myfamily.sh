#! /bin/bash/

export HERO_ID
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json  | jq ".[] | select(.id==$HERO_ID) | .connections.relatives" | tr -d '"'