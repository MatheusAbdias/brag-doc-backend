#!/bin/bash

base_dir="$(pwd)"    
source "$base_dir/.env"

while [[ $# -gt 0 ]]; do
    key="$1"

    case $key in
        -up)
            up_flag="-verbose up"
            shift
            shift
            ;;
        -down)
            down_flag="-verbose down"
            shift
            shift
            ;;
        *)
            echo "Unknown option $1"
            exit 1
            ;;
    esac
done

migrations_dir="$base_dir/internal/db/migrations"




cmd="migrate -path $migrations_dir -database $DATABASE_URL"
if [ -n "$up_flag" ]; then
    cmd="$cmd $up_flag"
fi
if [ -n "$down_flag" ]; then
    cmd="$cmd $down_flag"
fi
echo "Running command: $cmd"
eval $cmd
