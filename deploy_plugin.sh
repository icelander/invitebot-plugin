#!/bin/bash

jq_cmd=jq
[[ $(type -P "$jq_cmd") ]] || { 
	echo "'$jq_cmd' command line JSON processor not found";
	echo "Please install on linux with 'sudo apt-get install jq'"
	echo "Please install on mac with 'brew install jq'"
	exit 1; 
}

mmctl_cmd=mmctl
[[ $(type -P "$mmctl_cmd") ]] || { 
	echo "'$mmctl_cmd' command line JSON processor not found";
	echo "Please install on linux with 'sudo apt-get install mmctl'"
	echo "Please install on mac with 'brew install mmctl'"
	exit 1; 
}


id=`${jq_cmd} -r '.id' plugin.json`
name=`${jq_cmd} -r '.name' plugin.json`
version=`${jq_cmd} -r '.version' plugin.json`
archive_name="dist/${id}-${version}.tar.gz"

${mmctl_cmd} auth current

read -p "To deploy ${name} to the above server, press enter:"

# mmctl makes this *easy*
${mmctl_cmd} plugin disable $id
${mmctl_cmd} plugin delete $id
${mmctl_cmd} plugin add $archive_name
${mmctl_cmd} plugin enable $id
