#!/bin/bash

while [[ $# -gt 0 ]]; do
    key="$1"

    case $key in
        -d|--directory)
            directory="$2"
            shift
            shift
            ;;
        -n|--name)
            migration_name="$2"
            shift
            shift
            ;;
        *)
            echo "Unknown option $1"
            exit 1
            ;;
    esac
done

cmd="migrate create -ext sql -dir $(pwd)/$directory -seq $migration_name"

eval $cmd
