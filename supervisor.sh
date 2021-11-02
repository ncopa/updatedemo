#!/bin/sh

byebye() {
	echo "Shutting down"
	if [ -n "$childpid" ]; then
		echo "Cleaning up pid $childpid"
		kill $childpid
	fi
	exit
}

prog="$1"

trap "echo 'We got SIGCHLD'" SIGCHLD

trap byebye SIGINT SIGTERM

while true; do
	echo "Spawning $prog"
	$prog &
	childpid=$!
	wait $childpid
	echo "Child exited. respawning in a sec"
	sleep 1
done

