#!/bin/bash -xe

#Remember to set GOROOT accordingly with your installation

checkgopath () {

    GOPATH=$(printenv GOPATH)

    if [ -z $GOPATH ]
    then
        mkdir -p ${HOME}/go
        export GOPATH=${HOME}/go
    fi

}

# check if GOPATH variable is blank or not
checkgopath

# actual build
VERSION=`cat main.go| grep "const VERSION" |cut -f4 -d " " | tr -d '"'`

declare -a target_folders=("linux_amd64" "linux_386" "linux_arm" "linux_arm64" "linux_mips" "linux_mipsle" "darwin_amd64" "windows_386")

rm -rf distrib
mkdir distrib

package_index=`cat package_index.template | sed s/%%VERSION%%/${VERSION}/`

for folder in "${target_folders[@]}"
do
   rm -rf arduino-builder*
   rm -rf bin/
   mkdir -p bin
   IFS=_ read -a fields <<< $folder
   GOOS=${fields[0]} GOARCH=${fields[1]} go build
   FILENAME=arduino-builder-${VERSION}-${folder}.tar.bz2
   cp -r arduino-builder* bin/
   tar cjvf ${FILENAME} bin/
   T_OS=`echo ${folder} | awk '{print toupper($0)}'`
   SHASUM=`sha256sum ${FILENAME} | cut -f1 -d" "`
   SIZE=`stat --printf="%s" ${FILENAME}`
   package_index=`echo $package_index |
		sed s/%%FILENAME_${T_OS}%%/${FILENAME}/ |
		sed s/%%FILENAME_${T_OS}%%/${FILENAME}/ |
		sed s/%%SIZE_${T_OS}%%/${SIZE}/ |
		sed s/%%SHA_${T_OS}%%/${SHASUM}/`

   mv ${FILENAME} distrib/
done

set +x

# cleaning up
rm -rf bin/
rm arduino-builder*

# restoring original folder
git checkout .

echo ================== CUT ME HERE =====================

echo ${package_index} | python -m json.tool
