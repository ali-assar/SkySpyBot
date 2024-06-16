#!/bin/sh
# Start the Redis server in the background
redis-server &

# Wait for Redis server to start
sleep 5

# Start your application
/app/skyspy
