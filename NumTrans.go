/*
NumTrans.go
trans num to chinese type.
*/
/*下数者，十十变之；中数者，万万变之；上数者，数穷则变。
万：代表的是10的四次方。
亿：代表的是10的八次方。 
兆：代表的是10的十二次方。 
京：代表的是10的十六次方。 
垓：代表的是10的二十次方。 
杼：代表的是10的二十四次方。 
穰：代表的是10的二十八次方。 
沟：代表的是10的三十二次方。 
涧：代表的是10的三十六次方。 
正：代表的是10的四十次方。 
载：代表的是10的四十四次方。 
极：代表的是10的四十八次方。 
恒河沙：代表的是10的五十二次方。 
阿僧示氏：代表的是10的五十六次方。 
那由它：代表的是10的六十次方。 
不可思议：代表的是10的六十四次方。 
无量：代表的是10的六十八次方。 
大数：代表的是10的七十二次方。*/

package main

import "fmt"
import "unicode/utf8"


var digChar  = [...]byte{'零','壹','贰','叁','肆','伍','陆','柒','捌','玖'};
var numChar  = [...]byte{'元','拾','百','千','万','亿','兆','京','垓','杼','穰','沟','涧','正','载','极'};
var unitChar = [...]byte{'元','角','分'};

func Dig2Char(digChar byte)
{

}


func transDigStr(digStr string) string
{
    c := []rune(digStr);
    for i , v := range c{
        if v != '.' {
            continue;
        }
        else
        {
            break;
        }
    }
    //Get num char
    var maxnum = 0;
    maxnum = (i/4)+4; 

    //get unit char
    var j = 0;
    if i != len(c) {
        j = len(c) = i ;
    }

    //construct string
    for num := 0 ; num < i ; num++ {

    }



    return;
}


func main() {

    s := transDigStr("23429325.39");
    fmt.Printf("%s",s);
    return;
}
