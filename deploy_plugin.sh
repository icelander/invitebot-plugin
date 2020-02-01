#!/bin/bash

# TODO: Build tar file name from latest tag
archive_name="dist/invitebot-0.1.0.tar.gz"

# mmctl makes this *easy*
mmctl plugin disable invitebot
mmctl plugin delete invitebot
mmctl plugin add $archive_name
mmctl plugin enable invitebot
