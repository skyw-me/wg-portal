#!/bin/sh

# setup wireguard device
wg-quick down ${WG_DEFAULT_DEVICE}
wg-quick up ${WG_DEFAULT_DEVICE}

# main program
if [ -n "${WAIT_FOR}" ]; then
    /wait-for.sh "${WAIT_FOR}" -- /app/wgportal
else
    /app/wgportal
fi
