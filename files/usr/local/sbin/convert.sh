#!/bin/sh

TIME_TO_EXIT=30

timeout -t $TIME_TO_EXIT /usr/local/sbin/unoconv -vvv -f pdf $FILENAME || echo 'Time out '$TIME_TO_EXIT'sec left !!!' >> /tmp/timeout.log
chown $UID:$GID "${FILENAME%.*}.pdf"
