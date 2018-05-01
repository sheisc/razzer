#!/bin/bash
##############################################################################
# Copyright (c) 2014 Mark Charlebois
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to 
# deal in the Software without restriction, including without limitation the 
# rights to use, copy, modify, merge, publish, distribute, sublicense, and/or 
# sell copies of the Software, and to permit persons to whom the Software is 
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in 
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR 
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING 
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
# IN THE SOFTWARE.
##############################################################################

PROG=`basename $0`
USAGE="Usage: ${PROG} infile outfile [noexternal]"

# Sanity checking
[ ! -f $1 ] && echo ${USAGE} && exit
[ -z $2 ] && echo ${USAGE} && exit
[ ! -z $3 ] && [ ! "$3" = "noexternal" ] && echo ${USAGE} && exit

# Find find all unnamed sync nodes and prune them
# That removes the extra Node0x???????? generated by LLVM
cat $1 | sed -e "s/[\t]//" | grep "^Node" | cut -d' ' -f1 | sort | uniq > x2
cat $1 | sed -e "s/.*-> //" | grep "^Node" | cut -d';' -f1 | sort |uniq > x1
NODE=`comm -23 x1 x2`

# Remove external node links if noexternal passed
if [ "x$3" = "xnoexternal" ]; then
EXTERNAL_NODE=`cat $1 | grep "{external node}" | sed -e "s/[\t]//" | grep "^Node" | cut -d' ' -f1`
cat $1 | grep -v ${NODE} | grep -v ${EXTERNAL_NODE} > $2
else
cat $1 | grep -v ${NODE} > $2
fi

