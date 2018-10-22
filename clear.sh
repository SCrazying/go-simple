echo $PWD

#echo `ls $PWD`

if [ "/e/code/src/go-simple/xml.go" = "*.go" ]
then
echo "same"
else
echo "no same"
fi
function getdir(){
    for dir in  `ls $1`
        do
            dir_or_file=$PWD"/"$dir
            if  [ -d $dir_or_file ] 
            then 
                
                getdir $dir_or_file
            elif  [ "$dir_or_file"="*.exe" ]
            then 
                #echo $dir_or_file
                echo "$dir_or_file"
                #rm -rf $dir_or_file
            else
                echo "da"
            fi
        done 
}

getdir $root_dir
    


