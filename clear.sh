#!/bin/bash
function getdir()
{
   for dir in  `ls $1`
        do
            dir_or_file=$1"/"$dir
            if  [ -d $dir_or_file ] 
            then
                getdir $dir_or_file
           elif  [[ "$dir_or_file" == *.exe ]]
            then 
                echo "rm file ---  ""$dir_or_file"
                rm -rf $dir_or_file
            fi
        done 
}
getdir $PWD

