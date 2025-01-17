#!/bin/sh

# set wg config path as env variable
export WG_CONFIG_PATH=/app/data/wireguard

# copy initial config if not exists
mkdir -p "$WG_CONFIG_PATH"
cp -n /initial_config/* "$WG_CONFIG_PATH"

# setup wireguard device
WG_DEVICE_CONFIG_FILE="$WG_CONFIG_PATH/$WG_DEFAULT_DEVICE.conf"
wg-quick down "$WG_DEVICE_CONFIG_FILE"
wg-quick up "$WG_DEVICE_CONFIG_FILE"

# main program
/app/wgportal
