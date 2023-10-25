#! /bin/bash/

find . -iname  "*.sh" | cut -d "." -f2 | sed 's|.*/||'