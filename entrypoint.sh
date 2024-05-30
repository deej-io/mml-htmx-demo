#!/usr/bin/env bash

ls -la mml-server


./main &
(cd mml-server && npm start) &

wait -n

exit $?
