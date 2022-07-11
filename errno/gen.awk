# process errno header (ie /usr/include/asm-generic/errno-base.h)

/^#define.*[0-9]/ {
    name=$2
    val=$3
    msg=$5
    for(i=6; i<NF; i++) {msg = msg" "$i;}

    printf "\t// %s: %d\n", name, val
    printf "\tErrnosStr[\"%s\"] = Errno{\n", name
    printf "\t\tName: \"%s\",\n", name
    printf "\t\tValue: %d,\n", val
    printf "\t\tMessage: \"%s\",\n", msg
    print "\t}\n"
    printf "\tErrnosInt[%d] = Errno{\n", val
    printf "\t\tName: \"%s\",\n", name
    printf "\t\tValue: %d,\n", val
    printf "\t\tMessage: \"%s\",\n", msg
    print "\t}\n\n"
}
