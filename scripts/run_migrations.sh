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
entities=($(find "$migrations_dir" -maxdepth 1 -mindepth 1 -type d))

if [ ${#entities[@]} -eq 0 ]; then
    echo "No entities found, make sure you have migrations in $migrations_dir \n 
    and you have set the DATABASE_URL in .env file"
    exit 1
fi


for entity in "${entities[@]}"; do
    echo "Running migrations for $entity"
    cmd="migrate -path $entity -database $DATABASE_URL"
    if [ -n "$up_flag" ]; then
        cmd="$cmd $up_flag"
    fi
    if [ -n "$down_flag" ]; then
        cmd="$cmd $down_flag"
    fi
    echo "Running command: $cmd"
    eval $cmd
done