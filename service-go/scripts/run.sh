#!/usr/bin/env bash

# Just a script to run the demo
cmd=$1
service=$2


control_service() {
    pushd kubernetes >/dev/null

    echo $cmd-$service

    go run ../scripts/kubes.go -cmd $cmd -config -t config/$service-server-config.template.yaml -v values/$service-server-values.yaml

    popd >/dev/null
}

control() {
	if [ -z $service ]; then

        service="UnitableServer"
	    control_service $cmd

		return
	fi

    case $service in
    UnitableServer)
    control_service $cmd
       ;;
    esac
}

case $cmd in
	start)
	control create $service
	;;
	stop)
	control delete $service
	;;
	restart)
	control delete $service
	control create $service
	;;
	*)
	echo "$0 <start|stop|restart> [service]"
	exit
	;;
esac