# Wallpaper.net

> 该项目已停止维护,服务转移至 https://file.mo7.cc/daily_img

golang 实现获取 每日一图，每 8 小时自动更新一次 ， 跟 bing 每日一图保持一致

## 获取 cn.bing.com 官网当天的图片地址

<https://bz.mo7.cc/api/public/url>

使用方法：

get 请求, http 与 https 均兼容，允许跨域

返回的数据格式 `json`，是图片的 url 数组，可直接取用.

**第一张为当天 cn.bing.com 背景**

```json
{
  "Code": 1,
  "Data": [
    {
      "StartTime": 1653926400000,
      "EndTime": 1654012800000,
      "Path": "//bing.com/th?id=OHR.ParrotDay_EN-US9824481217_1920x1080.jpg",
      "Url": "https://bing.com/th?id=OHR.ParrotDay_EN-US9824481217_1920x1080.jpg",
      "Copyright": "Scarlet macaw in Costa Rica (© Harry Collins/Getty Images)",
      "CopyrightLink": "https://www.bing.com/search?q=Scarlet+macaw\u0026form=hpcapt\u0026filters=HpDate%3a%2220220531_0700%22",
      "Title": "This clever bird passes with flying colors"
    }
  ],
  "Msg": "Succeed"
}
```

## 直接返回图片

请求地址

```url

https://bz.mo7.cc/api/public/bz

```

使用方法

```html
<img src="http//bz.mo7.cc/api/public/bz" alt="" />
```

> 兼容 http 与 https

效果如下图

![bing.com](http://bz.mo7.cc/api/public/bz 'bz.mo7.cc')

请求参数

| 属性 | 类型 | 必填 | 默认 |
| :--: | :--: | :--: | :--: |
| idx  | int  |  否  | null |

> 0 则为随机返回某一张图片，1-8 则返回对应的壁纸, 例如 1 返回当前 cn.bing.com 的壁纸

使用方法

```html
<img src="//bz.mo7.cc/api/public/bz" alt="" />
<img src="//bz.mo7.cc/api/public/bz?idx=1" alt="" />
<img src="//bz.mo7.cc/api/public/bz?idx=1" alt="" />
```

效果如下图

![bing.com](http://bz.mo7.cc/api/public/bz?idx=1 'bz.mo7.cc')

> 博客地址：http://mo7.cc
