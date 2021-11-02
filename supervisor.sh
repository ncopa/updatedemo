#!/bin/sh

prog="$1"
trap "echo 'We got SIGCHLD'" SIGCHLD

while true; do
	echo "Spawning $prog"
	$prog &
	childpid=$!
	wait $childpid
	echo "Child exited. respawning in a sec"
	sleep 1
done

