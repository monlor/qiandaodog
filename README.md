#基于Go的签到程序
    用户直接执行./main即可。
    程序默认读取当前主程序文件夹内的 cookie.txt 、customize.txt 内容进行签到


#与旧版不同，用户存储的内容必须为严格json格式的字符串。
    cookie.txt里将存放签到数据的json字符串,如
    [
        {"Name": "baidu","Cookie": ""},
        {"Name": "v2ex","Cookie": ""}
    ]
    Name字段与Cookie字段的数组json
    Name字段重复的值即为多用户签到。

    customize.txt里存放的为用户自定义签到格式json字符串，如
    [
        {"Name":"Demo","Url":"https://github.com/Carseason/qiandaodog","Cookie":"demo","Types":"Get","Ua":""},
        {"Name":"Demo","Url":"https://github.com/Carseason/qiandaodog","Cookie":"demo","Types":"Post","Ua":"","Value":[{"Name":"PostFormName","Value":"PostFormValue"},{"Name":"PostFormName","Value":"PostFormValue"}]}
    
    ]
    Name字段为用户标识符，Url字段为要签到的Url,Cookie为用户的Cookie,Types字段为执行的请求,目前只支持 Get/POst ,Ua为用户执行签到的User-Agent，留空即为Pc浏览器，
    Vlaue字段为PostForm表单所传递的数据，类型为数组对象。如{"Name":"表单字段","Value":"表单值"}


#目前默认签到模板为:
    baidu           - 百度贴吧&百度文库
    v2ex            - V2EX
    hostloc         - hostloc.com
    acfun           - A站
    bilibili        - B站
    163music        - 网易云音乐PC
    miui            - 小米论坛
    52pojie         - 吾爱破解
    kafan           - 卡饭论坛
    smzdm           - 什么值得买
    zimuzu          - 人人影视字幕组
    gztown          - 港知堂社区PT站
    meizu           - 魅族论坛
    hdpfans         - 高清范
    chh             - chh论坛
    koolshare       - koolshare论坛
    right           - 恩山论坛
    huawei          - 花粉俱乐部
    ~~jd            - 移动端京东金融钢镚&京东领流量（因Cookie有效期为一天，故已删除）~~