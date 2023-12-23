#!/bin/sh

create_volume_subfolder() {
    # only change ownership if needed, if using an nfs mount this could be expensive
    if [ "$USER:$USER" != "$(stat /app -c '%U:%G')" ]
    then
        # Modify the owner of /app dir, make $USER(paopao) user have permission to create sub-dir in /app.
        chown -R "$USER:$USER" /app
    fi

    # Create VOLUME subfolder
    for f in /app/custom /app/meili_data; do
        if ! test -d $f; then
            gosu "$USER" mkdir -p $f
        fi
    done
}

setids() {
    export USER=paopao
    PUID=${PUID:-1000}
    PGID=${PGID:-1000}
    groupmod -o -g "$PGID" $USER
    usermod -o -u "$PUID" $USER
}

setids
create_volume_subfolder

# Exec CMD or S6 by default if nothing present
if [ $# -gt 0 ];then
    exec "$@"
else
    exec /bin/s6-svscan /app/gogs/docker/s6/
fi
