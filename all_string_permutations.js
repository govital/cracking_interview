function test(shortStr, longStr) {
    let counter = 0;
    const shortStrArr = shortStr.split('');
    longStr = longStr.replace(/ /g,'');

    for (let i=0; i < longStr.length; i++) {
        if(longStr.length-i < shortStr.length){
            return counter;
        }
        let tempArr = shortStrArr.map((x) => x);
        for(let k=0; k < shortStrArr.length; k++){
            if (!tempArr.includes(longStr[k])){
                break;
            }else{
                tempArr.splice(tempArr.indexOf(longStr[k]), 1);
            }
            if(tempArr.length == 0) {
                counter++;
                i = shortStrArr.length;
            }
        }
    }
    return counter;
}


const string1 = "abbc";
const string2 = "abbc";

console.log("number of permutations is: ", test(string1, string2));
