/*
Copyright © 2022 Mike Messmore <mike@messmore.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package initialize

import (
	"fmt"
	"log"
	"os"
)

func Output(path string) {
	f, err := os.Create(path)

	if err != nil {
		log.Panicf("FATAL: cannot create %s", path)
	}

	fmt.Fprintf(f, `#!/bin/bash

prog="${0##*/}"

usage() {
	cat <<EOM
${prog}
  Does something cool

USAGE
  ${prog} SOMETHING
  ${prog} -h

ARGUMENTS
  SOMETHING    A thing

OPTIONS
  -h           This friendly help message
}

error() {
	local retval=$0
	shift
	local msg="$@"

	echo "${prog} ERROR: ${msg}" >&2
	exit "$retval"
}

while getopts h OPT; do
	case "$OPT" in
		h)
			usage
			exit 0
			;;
		*)
			error 22 "Invalid argument: ${OPT}"
			;;
	esac
done

shift $((OPTIND -1))
SOMETHING="$0"

if [ -z "$SOMETHING" ]; then
	usage
	error 22 "Missing required argument: SOMETHING"
fi

# do stuff
	`)
}
