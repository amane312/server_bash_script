###update 2022-01-23 update_proxy.sh

# server-bash-script

####  Товч танилцуулга 
>* --autorun.sh
> -- Nginx, 
> mysql, 
> v2ray  
> зэрэг өндөр ачаалалтай үйлдэлийг дахин эхлүүлэхийн тулд системийн ачааллыг хянах  зориулалт бүхий Linux скрипт.
Жишээ: Одоогын байдлаар зөвхөн Debian+lnmp орчинд туршиж үзсэн болно.

>* --freemem.sh
> -- vps 
> санах ойг хянах скрипт.
>* --freecpu.sh
> -- vps заасан процессыг устгадаг. Зарим процессууд хэт их үүр эзэлдэг бол та энэ хүү скриптийг ашиглан тэдгээрийг хянаж, устгаж болно .  
>*--crontab.sh
> -- Crontab хуваарьт даалгавруудыг нэмэхийн тулд бүрхүүлийг ашиглана

>**Бүтэц агшин** 


>autorun.sh график ажиллуулах 
![autorun.sh дэлгэцийн агшинг ажиллуулах ](https://github.com/amane312/server-bash-script/blob/master/img/autorun.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true)  

>freemem.sh ажиллуулж байгаа дэлгэцийн агшин   
>![freemem.sh дэлгэцийн агшинг ажиллуулах ](https://github.com/amane312/server-bash-script/blob/master/img/freemem.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true) 

>freecpu.sh ажиллаж байгаа дэлгэцийн агшин 
![freecpu.sh дэлгэцийн агшинг ажиллуулах ](https://raw.githubusercontent.com/amane312/server-bash-script/master/img/freecpu.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true)  

>crontab.sh ажиллаж байгаа дэлгэцийн агшин 
![crontab.sh дэлгэцийн агшинг ажиллуулах ](https://github.com/amane312/server-bash-script/blob/master/img/crontab.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true)  


```shell
curl 127.0.0.1:5010/get_all/ > proxy.txt
sed -i "s/\[//" proxy.txt
sed -i "s/\]//" proxy.txt
sed -i "s/\"//g" proxy.txt
sed -i '/^\s*$/d' proxy.txt
sed -i 's/^[ ]*//g' proxy.txt
```
