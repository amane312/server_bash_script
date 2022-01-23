###update 2022-01-23 update_proxy.sh

# server-bash-script
some userful server bash script
####  Товч танилцуулга 
>* --autorun.sh-- Nginx, mysql, v2ray зэрэг өндөр ачаалалтай зарим үйлчилгээгээ дахин эхлүүлэхийн тулд системийн ачааллыг хянах Linux скрипт.
Жич: Одоогоор зөвхөн Debian+lnmp орчинд туршиж үзсэн.Алдаа гарсан тохиолдолд алдаа гаргаж, засварлах цагийг олоорой. 
>* --freemem.sh-- vps санах ойг гаргах скрипт, санах ойг тогтмол гаргах (би буффыг стандарт болгон авдаг, та үүнийг бодит үлдэгдэл болгон өөрчлөх боломжтой, скрипт бичигдсэн, та өөрөө тайлбар хийж болно). 
>* --freecpu.sh-- vps заасан процессыг устгадаг. Зарим процессууд хэт их CPU эзэлдэг бол та энэ скриптийг ашиглан тэдгээрийг хянаж, устгаж болно (энэ нь хүчирхийллийн шийдэл бөгөөд үүнийг зөвлөдөггүй! Эх сурвалжийг олох нь дээр. асуудал, түүнийг шийдвэрлэх).  
>*--crontab.sh-- Crontab хуваарьт даалгавруудыг нэмэхийн тулд бүрхүүлийг ашиглана

>**Бүтэц агшин** 


>autorun.sh график ажиллуулах 
![autorun.sh дэлгэцийн агшинг ажиллуулах ](https://github.com/Mr-xn/server-bash-script/blob/master/img/autorun.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true)  

>freemem.sh ажиллуулж байгаа дэлгэцийн агшин   
>![freemem.sh дэлгэцийн агшинг ажиллуулах ](https://github.com/Mr-xn/server-bash-script/blob/master/img/freemem.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true) 

>freecpu.sh ажиллаж байгаа дэлгэцийн агшин 
![freecpu.sh дэлгэцийн агшинг ажиллуулах ](https://raw.githubusercontent.com/Mr-xn/server-bash-script/master/img/freecpu.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true)  

>crontab.sh ажиллаж байгаа дэлгэцийн агшин 
![crontab.sh дэлгэцийн агшинг ажиллуулах ](https://github.com/Mr-xn/server-bash-script/blob/master/img/crontab.sh%E8%BF%90%E8%A1%8C%E6%88%AA%E5%9B%BE.png?raw=true)  

#### Энэ скрипт нь Xiaobai-д тохиромжтой, хэрэв та ойлгохгүй байвал [blog](https://mrxn.net) руу очно уу. Мессеж үлдээнэ үү. 

#### Ашиглалтын явцад дараах алдааны мэдэгдэл гарч ирвэл ``dos2unix`` суулгаад Windows-с байршуулсан скриптийг Unix формат руу хөрвүүлнэ үү!  

`` xxxx.sh: line 2: $'\r': command not found ``   

#### Скриптийг ашиглахаасаа өмнө ``dos2unix`` суулгаж хөрвүүлэхийг зөвлөж байна! [Нарийвчилсан заавар](https://mrxn.net/jswz/570.html)<<—— 

#### ``Xshell`` дээрх ``lrzsz`` файлын хурдан дамжуулах туслах [дэлгэрэнгүй заавар](https://mrxn.net/Linux/542.html)<<——
### Энэ бүрхүүлийг [htpwdScan](https://github.com/lijiejie/htpwdScan) [proxy_pool](https://github.com/jhao104/proxy_pool)-д тусгайлан тохируулсан. 
> Дараах агуулгыг update_proxy.sh файл болгон хадгалж, htpwdScan лавлах доор байрлуулна уу, sh update_proxy.sh-г ажиллуулж болно, мөн та үүнийг хуваарьт ажил руу шидэж автоматаар гүйцэтгэх боломжтой: сэрүүлэгтэй цаг: */6 * * * * sh / your/path /to/htpwdScan/update_proxy.sh 

```shell
curl 127.0.0.1:5010/get_all/ > proxy.txt
sed -i "s/\[//" proxy.txt
sed -i "s/\]//" proxy.txt
sed -i "s/\"//g" proxy.txt
sed -i '/^\s*$/d' proxy.txt
sed -i 's/^[ ]*//g' proxy.txt
```

- Жич: Хэрэв та юу ч ойлгохгүй байвал эхлээд Google-ээс хай, эргүүлж чадахгүй бол миний [Google хайлт] (https://g.mrxn.net/) [зөвхөн хайлт, суралцахтай холбоотой) ашиглаж болно. ], дараа нь бүтэн хайсны дараа асуулт асуугаарай, үгүй ​​Намайг ирсэн даруйдаа би асуув: x: htpwdScan шиг зүйл юу хийдэг вэ? proxy_pool нь нэг төрлийн хуурай га юм. Хүн бүрт цаг хугацаа хязгаарлагдмал, баярлалаа! 
