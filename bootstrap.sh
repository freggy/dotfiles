#!/bin/bash -e

function copy_files {
	# find all paths to files in $dir	
	IFS=$'\n' files=($(find $dir -mindepth 1 -type f))

	for file in "${files[@]}"
	do
		dirs=$(dirname $file)

		# copy files to home directory
		if [ "$1" = true ]; 
		then
			home=${HOME:1} # home path without the leading /
			dirs="${dirs/__homedir/$home}"
		fi
		echo "creating directories if not exist: /$dirs"
		echo "copying $file to /$dirs/$(basename $file)"
	done
}


# installs base programs like brew to install other dependencies.
function install_base_progs {
	echo "install base progs"
}

function install_progs_brew {
	echo "install brew progs"
}

function install_files {
	for dir in */
	do
		if [ $dir == "__homedir/" ];
		then
			copy_files true
			continue
		fi

		copy_files
	done
}

# if --update do
#   git pull
#   install_files

install_files
install_base_progs
install_progs_brew

