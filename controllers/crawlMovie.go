package controllers

import (
	"crawl_movie/models"
	"github.com/astaxie/beego"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {

	sMovieHtml := `<!DOCTYPE html>
<html lang="zh-cmn-Hans" class="ua-windows ua-webkit">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="renderer" content="webkit">
    <meta name="referrer" content="always">
    <meta name="google-site-verification" content="ok0wCgT20tBBgo9_zat2iAcimtN4Ftf5ccsh092Xeyw" />
    <title>
        扫毒2天地对决 (豆瓣)
</title>
    
    <meta name="baidu-site-verification" content="cZdR4xxR7RxmM4zE" />
    <meta http-equiv="Pragma" content="no-cache">
    <meta http-equiv="Expires" content="Sun, 6 Mar 2005 01:00:00 GMT">
    
    <link rel="apple-touch-icon" href="https://img3.doubanio.com/f/movie/d59b2715fdea4968a450ee5f6c95c7d7a2030065/pics/movie/apple-touch-icon.png">
    <link href="https://img3.doubanio.com/f/shire/52c9997d6d42db58eab418e976a14d5f3eff981e/css/douban.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/shire/ae3f5a3e3085968370b1fc63afcecb22d3284848/css/separation/_all.css" rel="stylesheet" type="text/css">
    <link href="https://img3.doubanio.com/f/movie/8864d3756094f5272d3c93e30ee2e324665855b0/css/movie/base/init.css" rel="stylesheet">
    <script type="text/javascript">var _head_start = new Date();</script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/movie/0495cb173e298c28593766009c7b0a953246c5b5/js/movie/lib/jquery.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/1316664523258f7b8b536e4ce45afc9cb37b8963/js/douban.js"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/0efdc63b77f895eaf85281fb0e44d435c6239a3f/js/separation/_all.js"></script>
    
    <meta name="keywords" content="扫毒2天地对决,掃毒2：天地對決,扫毒2天地对决影评,剧情介绍,电影图片,预告片,影讯,在线购票,论坛">
    <meta name="description" content="扫毒2天地对决电影简介和剧情介绍,扫毒2天地对决影评、图片、预告片、影讯、论坛、在线购票">
    <meta name="mobile-agent" content="format=html5; url=http://m.douban.com/movie/subject/30171425/"/>
    <link rel="alternate" href="android-app://com.douban.movie/doubanmovie/subject/30171425/" />
    <link rel="stylesheet" href="https://img3.doubanio.com/dae/cdnlib/libs/LikeButton/1.0.5/style.min.css">
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript">
      Do.add('dialog', {path: 'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js', type: 'js'});
      Do.add('dialog-css', {path: 'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css', type: 'css'});
      Do.add('handlebarsjs', {path: 'https://img3.doubanio.com/f/movie/3d4f8e4a8918718256450eb6e57ec8e1f7a2e14b/js/movie/lib/handlebars.current.js', type: 'js'});
    </script>
    
  <script type='text/javascript'>
  var _vwo_code = (function() {
    var account_id = 249272,
      settings_tolerance = 0,
      library_tolerance = 2500,
      use_existing_jquery = false,
      // DO NOT EDIT BELOW THIS LINE
      f=false,d=document;return{use_existing_jquery:function(){return use_existing_jquery;},library_tolerance:function(){return library_tolerance;},finish:function(){if(!f){f=true;var a=d.getElementById('_vis_opt_path_hides');if(a)a.parentNode.removeChild(a);}},finished:function(){return f;},load:function(a){var b=d.createElement('script');b.src=a;b.type='text/javascript';b.innerText;b.onerror=function(){_vwo_code.finish();};d.getElementsByTagName('head')[0].appendChild(b);},init:function(){settings_timer=setTimeout('_vwo_code.finish()',settings_tolerance);var a=d.createElement('style'),b='body{opacity:0 !important;filter:alpha(opacity=0) !important;background:none !important;}',h=d.getElementsByTagName('head')[0];a.setAttribute('id','_vis_opt_path_hides');a.setAttribute('type','text/css');if(a.styleSheet)a.styleSheet.cssText=b;else a.appendChild(d.createTextNode(b));h.appendChild(a);this.load('//dev.visualwebsiteoptimizer.com/j.php?a='+account_id+'&u='+encodeURIComponent(d.URL)+'&r='+Math.random());return settings_timer;}};}());

  +function () {
    var bindEvent = function (el, type, handler) {
        var $ = window.jQuery || window.Zepto || window.$
       if ($ && $.fn && $.fn.on) {
           $(el).on(type, handler)
       } else if($ && $.fn && $.fn.bind) {
           $(el).bind(type, handler)
       } else if (el.addEventListener){
         el.addEventListener(type, handler, false);
       } else if (el.attachEvent){
         el.attachEvent("on" + type, handler);
       } else {
         el["on" + type] = handler;
       }
     }

    var _origin_load = _vwo_code.load
    _vwo_code.load = function () {
      var args = [].slice.call(arguments)
      bindEvent(window, 'load', function () {
        _origin_load.apply(_vwo_code, args)
      })
    }
  }()

  _vwo_settings_timer = _vwo_code.init();
  </script>


    


<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "name": "扫毒2天地对决 掃毒2：天地對決",
  "url": "/subject/30171425/",
  "image": "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561172733.webp",
  "director": 
  [
    {
      "@type": "Person",
      "url": "/celebrity/1274313/",
      "name": "邱礼涛 Herman Yau"
    }
  ]
,
  "author": 
  [
    {
      "@type": "Person",
      "url": "/celebrity/1289664/",
      "name": "李敏 Erica Lee"
    }
  ]
,
  "actor": 
  [
    {
      "@type": "Person",
      "url": "/celebrity/1054424/",
      "name": "刘德华 Andy Lau"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1027577/",
      "name": "古天乐 Louis Koo"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1115415/",
      "name": "苗侨伟 Kiu Wai Miu"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1204410/",
      "name": "林嘉欣 Karena Lam"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1313563/",
      "name": "周秀娜 Chrissie Chaw"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1019372/",
      "name": "应采儿 Cherrie Ying"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1275972/",
      "name": "陈家乐 Carlos Chan"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1313357/",
      "name": "卫诗雅 Michelle Wai"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1019278/",
      "name": "郑则仕 Kent Cheng"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1274568/",
      "name": "张国强 Kwok Keung Cheung"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1413084/",
      "name": "陈佩诗 Peishi Chen"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1347734/",
      "name": "狼森 Philippe Joly"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1193019/",
      "name": "欧阳靖 Jin Auyeung"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1275254/",
      "name": "恭硕良 Jun Kung"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1050329/",
      "name": "林家栋 Ka Tung Lam"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1399689/",
      "name": "李赏 Faith Lee"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1371353/",
      "name": "张竣杰 Chun Kit Chang"
    }
    ,
    {
      "@type": "Person",
      "url": "/celebrity/1028522/",
      "name": "李灿森 Sam Lee"
    }
  ]
,
  "datePublished": "2019-07-05",
  "genre": ["\u5267\u60c5", "\u52a8\u4f5c", "\u72af\u7f6a", "\u60ac\u7591"],
  "duration": "PT1H39M",
  "description": "毒品市场维持四分天下的格局已久，但自从地藏与墨西哥大毒枭跨境合作扩展势力，再加上一连串黑吃黑的动作，毒界变得风声鹤唳。另一方面，因儿时亲眼目睹父亲被毒品所毁而嫉毒如仇的慈善家兼金融巨子余顺天，正悬赏一...",
  "@type": "Movie",
  "aggregateRating": {
    "@type": "AggregateRating",
    "ratingCount": "39966",
    "bestRating": "10",
    "worstRating": "2",
    "ratingValue": "6.3"
  }
}
</script>


    <style type="text/css">
  
</style>
    <style type="text/css">img { max-width: 100%; }</style>
    <script type="text/javascript"></script>
    <link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/59f798f5b75e51f0.css">

    <link rel="shortcut icon" href="https://img3.doubanio.com/favicon.ico" type="image/x-icon">
</head>

<body>
  
    <script type="text/javascript">var _body_start = new Date();</script>

    
    



    <link href="//img3.doubanio.com/dae/accounts/resources/ec819da/shire/bundle.css" rel="stylesheet" type="text/css">



<div id="db-global-nav" class="global-nav">
  <div class="bd">
    
<div class="top-nav-info">
  <a href="https://accounts.douban.com/passport/login?source=movie" class="nav-login" rel="nofollow">登录/注册</a>
</div>


    <div class="top-nav-doubanapp">
  <a href="https://www.douban.com/doubanapp/app?channel=top-nav" class="lnk-doubanapp">下载豆瓣客户端</a>
  <div id="doubanapp-tip">
    <a href="https://www.douban.com/doubanapp/app?channel=qipao" class="tip-link">豆瓣 <span class="version">6.0</span> 全新发布</a>
    <a href="javascript: void 0;" class="tip-close">×</a>
  </div>
  <div id="top-nav-appintro" class="more-items">
    <p class="appintro-title">豆瓣</p>
    <p class="qrcode">扫码直接下载</p>
    <div class="download">
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=iOS">iPhone</a>
      <span>·</span>
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=Android" class="download-android">Android</a>
    </div>
  </div>
</div>

    


<div class="global-nav-items">
  <ul>
    <li class="">
      <a href="https://www.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-main&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣</a>
    </li>
    <li class="">
      <a href="https://book.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-book&quot;,&quot;uid&quot;:&quot;0&quot;}">读书</a>
    </li>
    <li class="on">
      <a href="https://movie.douban.com"  data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-movie&quot;,&quot;uid&quot;:&quot;0&quot;}">电影</a>
    </li>
    <li class="">
      <a href="https://music.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-music&quot;,&quot;uid&quot;:&quot;0&quot;}">音乐</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/location" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-location&quot;,&quot;uid&quot;:&quot;0&quot;}">同城</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/group" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-group&quot;,&quot;uid&quot;:&quot;0&quot;}">小组</a>
    </li>
    <li class="">
      <a href="https://read.douban.com&#47;?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-read&quot;,&quot;uid&quot;:&quot;0&quot;}">阅读</a>
    </li>
    <li class="">
      <a href="https://douban.fm&#47;?from_=shire_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-fm&quot;,&quot;uid&quot;:&quot;0&quot;}">FM</a>
    </li>
    <li class="">
      <a href="https://time.douban.com&#47;?dt_time_source=douban-web_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-time&quot;,&quot;uid&quot;:&quot;0&quot;}">时间</a>
    </li>
    <li class="">
      <a href="https://market.douban.com&#47;?utm_campaign=douban_top_nav&amp;utm_source=douban&amp;utm_medium=pc_web" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-market&quot;,&quot;uid&quot;:&quot;0&quot;}">豆品</a>
    </li>
    <li>
      <a href="#more" class="bn-more"><span>更多</span></a>
      <div class="more-items">
        <table cellpadding="0" cellspacing="0">
          <tbody>
            <tr>
              <td>
                <a href="https://ypy.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-ypy&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣摄影</a>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </li>
  </ul>
</div>

  </div>
</div>
<script>
  ;window._GLOBAL_NAV = {
    DOUBAN_URL: "https://www.douban.com",
    N_NEW_NOTIS: 0,
    N_NEW_DOUMAIL: 0
  };
</script>



    <script src="//img3.doubanio.com/dae/accounts/resources/ec819da/shire/bundle.js" defer="defer"></script>




    



    <link href="//img3.doubanio.com/dae/accounts/resources/ec819da/movie/bundle.css" rel="stylesheet" type="text/css">




<div id="db-nav-movie" class="nav">
  <div class="nav-wrap">
  <div class="nav-primary">
    <div class="nav-logo">
      <a href="https:&#47;&#47;movie.douban.com">豆瓣电影</a>
    </div>
    <div class="nav-search">
      <form action="https:&#47;&#47;movie.douban.com/subject_search" method="get">
        <fieldset>
          <legend>搜索：</legend>
          <label for="inp-query">
          </label>
          <div class="inp"><input id="inp-query" name="search_text" size="22" maxlength="60" placeholder="搜索电影、电视剧、综艺、影人" value=""></div>
          <div class="inp-btn"><input type="submit" value="搜索"></div>
          <input type="hidden" name="cat" value="1002" />
        </fieldset>
      </form>
    </div>
  </div>
  </div>
  <div class="nav-secondary">
    

<div class="nav-items">
  <ul>
    <li    ><a href="https://movie.douban.com/cinema/nowplaying/"
     >影讯&购票</a>
    </li>
    <li    ><a href="https://movie.douban.com/explore"
     >选电影</a>
    </li>
    <li    ><a href="https://movie.douban.com/tv/"
     >电视剧</a>
    </li>
    <li    ><a href="https://movie.douban.com/chart"
     >排行榜</a>
    </li>
    <li    ><a href="https://movie.douban.com/tag/"
     >分类</a>
    </li>
    <li    ><a href="https://movie.douban.com/review/best/"
     >影评</a>
    </li>
    <li    ><a href="https://movie.douban.com/annual/2018?source=navigation"
     >2018年度榜单</a>
    </li>
    <li    ><a href="https://www.douban.com/standbyme/2018?source=navigation"
     >2018书影音报告</a>
    </li>
  </ul>
</div>

    <a href="https://movie.douban.com/annual/2018?source=movie_navigation" class="movieannual2018"></a>
  </div>
</div>

<script id="suggResult" type="text/x-jquery-tmpl">
  <li data-link="{{= url}}">
            <a href="{{= url}}" onclick="moreurl(this, {from:'movie_search_sugg', query:'{{= keyword }}', subject_id:'{{= id}}', i: '{{= index}}', type: '{{= type}}'})">
            <img src="{{= img}}" width="40" />
            <p>
                <em>{{= title}}</em>
                {{if year}}
                    <span>{{= year}}</span>
                {{/if}}
                {{if sub_title}}
                    <br /><span>{{= sub_title}}</span>
                {{/if}}
                {{if address}}
                    <br /><span>{{= address}}</span>
                {{/if}}
                {{if episode}}
                    {{if episode=="unknow"}}
                        <br /><span>集数未知</span>
                    {{else}}
                        <br /><span>共{{= episode}}集</span>
                    {{/if}}
                {{/if}}
            </p>
        </a>
        </li>
  </script>




    <script src="//img3.doubanio.com/dae/accounts/resources/ec819da/movie/bundle.js" defer="defer"></script>





    
    <div id="wrapper">
        

        
    <div id="content">
        

    <div id="dale_movie_subject_top_icon"></div>
    <h1>
        <span property="v:itemreviewed">扫毒2天地对决 掃毒2：天地對決</span>
            <span class="year">(2019)</span>
    </h1>

        <div class="grid-16-8 clearfix">
            

            
            <div class="article">
                
    

    





        <div class="indent clearfix">
            <div class="subjectwrap clearfix">
                <div class="subject clearfix">
                    



<div id="mainpic" class="">
    <a class="nbgnbg" href="https://movie.douban.com/subject/30171425/photos?type=R" title="点击看更多海报">
        <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561172733.webp" title="点击看更多海报" alt="掃毒2：天地對決" rel="v:image" />
   </a>
</div>

                    


<div id="info">
        <span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/1274313/" rel="v:directedBy">邱礼涛</a></span></span><br/>
        <span ><span class='pl'>编剧</span>: <span class='attrs'><a href="/celebrity/1289664/">李敏</a></span></span><br/>
        <span class="actor"><span class='pl'>主演</span>: <span class='attrs'><a href="/celebrity/1054424/" rel="v:starring">刘德华</a> / <a href="/celebrity/1027577/" rel="v:starring">古天乐</a> / <a href="/celebrity/1115415/" rel="v:starring">苗侨伟</a> / <a href="/celebrity/1204410/" rel="v:starring">林嘉欣</a> / <a href="/celebrity/1313563/" rel="v:starring">周秀娜</a> / <a href="/celebrity/1019372/" rel="v:starring">应采儿</a> / <a href="/celebrity/1275972/" rel="v:starring">陈家乐</a> / <a href="/celebrity/1313357/" rel="v:starring">卫诗雅</a> / <a href="/celebrity/1019278/" rel="v:starring">郑则仕</a> / <a href="/celebrity/1274568/" rel="v:starring">张国强</a> / <a href="/celebrity/1413084/" rel="v:starring">陈佩诗</a> / <a href="/celebrity/1347734/" rel="v:starring">狼森</a> / <a href="/celebrity/1193019/" rel="v:starring">欧阳靖</a> / <a href="/celebrity/1275254/" rel="v:starring">恭硕良</a> / <a href="/celebrity/1050329/" rel="v:starring">林家栋</a> / <a href="/celebrity/1399689/" rel="v:starring">李赏</a> / <a href="/celebrity/1371353/" rel="v:starring">张竣杰</a> / <a href="/celebrity/1028522/" rel="v:starring">李灿森</a></span></span><br/>
        <span class="pl">类型:</span> <span property="v:genre">剧情</span> / <span property="v:genre">动作</span> / <span property="v:genre">悬疑</span> / <span property="v:genre">犯罪</span><br/>
        
        <span class="pl">制片国家/地区:</span> 中国大陆 / 香港<br/>
        <span class="pl">语言:</span> 汉语普通话 / 粤语<br/>
        <span class="pl">上映日期:</span> <span property="v:initialReleaseDate" content="2019-07-05(中国大陆)">2019-07-05(中国大陆)</span> / <span property="v:initialReleaseDate" content="2019-07-25(香港)">2019-07-25(香港)</span><br/>
        <span class="pl">片长:</span> <span property="v:runtime" content="99">99分钟</span><br/>
        <span class="pl">又名:</span> 扫毒2：天地对决 / The White Storm 2: Drug Lords<br/>
        <span class="pl">IMDb链接:</span> <a href="http://www.imdb.com/title/tt8239806" target="_blank" rel="nofollow">tt8239806</a><br>

</div>




                </div>
                    


<div id="interest_sectl">
    <div class="rating_wrap clearbox" rel="v:rating">
        <div class="clearfix">
          <div class="rating_logo ll">豆瓣评分</div>
          <div class="output-btn-wrap rr" style="display:none">
            <img src="https://img3.doubanio.com/f/movie/692e86756648f29457847c5cc5e161d6f6b8aaac/pics/movie/reference.png" />
            <a class="download-output-image" href="#">引用</a>
          </div>
          
          
        </div>
        



<div class="rating_self clearfix" typeof="v:Rating">
    <strong class="ll rating_num" property="v:average">6.3</strong>
    <span property="v:best" content="10.0"></span>
    <div class="rating_right ">
        <div class="ll bigstar bigstar30"></div>
        <div class="rating_sum">
                <a href="collections" class="rating_people">
                    <span property="v:votes">40424</span>人评价
                </a>
        </div>
    </div>
</div>
<div class="ratings-on-weight">
    
        <div class="item">
        
        <span class="stars5 starstop" title="力荐">
            5星
        </span>
        <div class="power" style="width:6px"></div>
        <span class="rating_per">5.6%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars4 starstop" title="推荐">
            4星
        </span>
        <div class="power" style="width:29px"></div>
        <span class="rating_per">24.1%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars3 starstop" title="还行">
            3星
        </span>
        <div class="power" style="width:64px"></div>
        <span class="rating_per">51.5%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars2 starstop" title="较差">
            2星
        </span>
        <div class="power" style="width:20px"></div>
        <span class="rating_per">16.1%</span>
        <br />
        </div>
        <div class="item">
        
        <span class="stars1 starstop" title="很差">
            1星
        </span>
        <div class="power" style="width:3px"></div>
        <span class="rating_per">2.7%</span>
        <br />
        </div>
</div>

    </div>
        <div class="rating_betterthan">
            好于 <a href="/typerank?type_name=犯罪&type=3&interval_id=30:20&action=">25% 犯罪片</a><br/>
            好于 <a href="/typerank?type_name=动作&type=5&interval_id=40:30&action=">36% 动作片</a><br/>
        </div>
</div>


                
            </div>
                




<div id="interest_sect_level" class="clearfix">
        
            <a href="https://www.douban.com/reason=collectwish&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-30171425-wish">
                <span>想看</span>
            </a>
            <a href="https://www.douban.com/reason=collectcollect&amp;ck=" rel="nofollow" class="j a_show_login colbutt ll" name="pbtn-30171425-collect">
                <span>看过</span>
            </a>
        <div class="ll j a_stars">
            
    
    评价:
    <span id="rating"> <span id="stars" data-solid="https://img3.doubanio.com/f/shire/5a2327c04c0c231bced131ddf3f4467eb80c1c86/pics/rating_icons/star_onmouseover.png" data-hollow="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" data-solid-2x="https://img3.doubanio.com/f/shire/7258904022439076d57303c3b06ad195bf1dc41a/pics/rating_icons/star_onmouseover@2x.png" data-hollow-2x="https://img3.doubanio.com/f/shire/95cc2fa733221bb8edd28ad56a7145a5ad33383e/pics/rating_icons/star_hollow_hover@2x.png">

            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30171425-1">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star1" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30171425-2">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star2" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30171425-3">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star3" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30171425-4">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star4" width="16" height="16"/></a>
            <a href="https://www.douban.com/register?reason=rate" class="j a_show_login" name="pbtn-30171425-5">
        <img src="https://img3.doubanio.com/f/shire/2520c01967207a1735171056ec588c8c1257e5f8/pics/rating_icons/star_hollow_hover.png" id="star5" width="16" height="16"/></a>
    </span><span id="rateword" class="pl"></span>
    <input id="n_rating" type="hidden" value=""  />
    </span>

        </div>
    

</div>


            


















<div class="gtleft">
    <ul class="ul_subject_menu bicelink color_gray pt6 clearfix">
        
    
        
                <li> 
    <img src="https://img3.doubanio.com/f/shire/cc03d0fcf32b7ce3af7b160a0b85e5e66b47cc42/pics/short-comment.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_cmnt_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写短评</a>
 </li>
                    <li> 
    
    <img src="https://img3.doubanio.com/f/shire/5bbf02b7b5ec12b23e214a580b6f9e481108488c/pics/add-review.gif" />&nbsp;
        <a onclick="moreurl(this, {from:'mv_sbj_wr_rv_login'})" class="j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">写影评</a>
 </li>
                    <li> 
    <img src="https://img3.doubanio.com/f/shire/61cc48ba7c40e0272d46bb93fe0dc514f3b71ec5/pics/add-doulist.gif" />&nbsp;
    <a class="" href="/subject/30171425/questions/ask?from=subject_top">提问题</a>
 </li>
                <li> 
    



 </li>
                <li> 
   

   
    
    <span class="rec" id="电影-30171425">
    <a href= "#"
        data-type="电影"
        data-url="https://movie.douban.com/subject/30171425/"
        data-desc="电影《扫毒2天地对决 掃毒2：天地對決》 (来自豆瓣) "
        data-title="电影《扫毒2天地对决 掃毒2：天地對決》 (来自豆瓣) "
        data-pic="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561172733.jpeg"
        class="bn-sharing ">
        分享到
    </a> &nbsp;&nbsp;
    </span>

    <script>
        if (!window.DoubanShareMenuList) {
            window.DoubanShareMenuList = [];
        }
        var __cache_url = __cache_url || {};

        (function(u){
            if(__cache_url[u]) return;
            __cache_url[u] = true;
            window.DoubanShareIcons = 'https://img3.doubanio.com/f/shire/d15ffd71f3f10a7210448fec5a68eaec66e7f7d0/pics/ic_shares.png';

            var initShareButton = function() {
                $.ajax({url:u,dataType:'script',cache:true});
            };

            if (typeof Do == 'function' && 'ready' in Do) {
                Do(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js',
                    initShareButton
                );
            } else if(typeof Douban == 'object' && 'loader' in Douban) {
                Douban.loader.batch(
                    'https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css',
                    'https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js',
                    'https://img3.doubanio.com/f/movie/c4ab132ff4d3d64a83854c875ea79b8b541faf12/js/movie/lib/qrcode.min.js'
                ).done(initShareButton);
            }

        })('https://img3.doubanio.com/f/movie/32be6727ed3ad8f6c4a417d8a086355c3e7d1d27/js/movie/lib/sharebutton.js');
    </script>


  </li>
            

    </ul>

    <script type="text/javascript">
        $(function(){
            $(".ul_subject_menu li.rec .bn-sharing").bind("click", function(){
                $.get("/blank?sbj_page_click=bn_sharing");
            });
            $(".ul_subject_menu .create_from_menu").bind("click", function(e){
                e.preventDefault();
                var $el = $(this);
                var glRoot = document.getElementById('gallery-topics-selection');
                if (window.has_gallery_topics && glRoot) {
                    // 判断是否有话题
                    glRoot.style.display = 'block';
                } else {
                    location.href = $el.attr('href');
                }
            });
        });
    </script>
</div>




                





<div class="rec-sec">
<span class="rec">
    <script id="movie-share" type="text/x-html-snippet">
        
    <form class="movie-share" action="/j/share" method="POST">
        <div class="clearfix form-bd">
            <div class="input-area">
                <textarea name="text" class="share-text" cols="72" data-mention-api="https://api.douban.com/shuo/in/complete?alt=xd&amp;callback=?"></textarea>
                <input type="hidden" name="target-id" value="30171425">
                <input type="hidden" name="target-type" value="0">
                <input type="hidden" name="title" value="扫毒2天地对决 掃毒2：天地對決‎ (2019)">
                <input type="hidden" name="desc" value="导演 邱礼涛 主演 刘德华 / 古天乐 / 中国大陆 / 香港 / 6.3分(40424评价)">
                <input type="hidden" name="redir" value=""/>
                <div class="mentioned-highlighter"></div>
            </div>

            <div class="info-area">
                    <img class="media" src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561172733.webp" />
                <strong>扫毒2天地对决 掃毒2：天地對決‎ (2019)</strong>
                <p>导演 邱礼涛 主演 刘德华 / 古天乐 / 中国大陆 / 香港 / 6.3分(40424评价)</p>
                <p class="error server-error">&nbsp;</p>
            </div>
        </div>
        <div class="form-ft">
            <div class="form-ft-inner">
                



                <span class="avail-num-indicator">140</span>
                <span class="bn-flat">
                    <input type="submit" value="推荐" />
                </span>
            </div>
        </div>
    </form>
    
    <div id="suggest-mention-tmpl" style="display:none;">
        <ul>
            {{#users}}
            <li id="{{uid}}">
              <img src="{{avatar}}">{{{username}}}&nbsp;<span>({{{uid}}})</span>
            </li>
            {{/users}}
        </ul>
    </div>


    </script>

        
        <a href="/accounts/register?reason=recommend"  class="j a_show_login lnk-sharing" share-id="30171425" data-mode="plain" data-name="扫毒2天地对决 掃毒2：天地對決‎ (2019)" data-type="movie" data-desc="导演 邱礼涛 主演 刘德华 / 古天乐 / 中国大陆 / 香港 / 6.3分(40424评价)" data-href="https://movie.douban.com/subject/30171425/" data-image="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561172733.webp" data-properties="{}" data-redir="" data-text="" data-apikey="" data-curl="" data-count="10" data-object_kind="1002" data-object_id="30171425" data-target_type="rec" data-target_action="0" data-action_props="{&#34;subject_url&#34;:&#34;https:\/\/movie.douban.com\/subject\/30171425\/&#34;,&#34;subject_title&#34;:&#34;扫毒2天地对决 掃毒2：天地對決‎ (2019)&#34;}">推荐</a>
</span>


</div>






            <script type="text/javascript">
                $(function() {
                    $('.collect_btn', '#interest_sect_level').each(function() {
                        Douban.init_collect_btn(this);
                    });
                    $('html').delegate(".indent .rec-sec .lnk-sharing", "click", function() {
                        moreurl(this, {
                            from : 'mv_sbj_db_share'
                        });
                    });
                });
            </script>
        </div>
            


    <div id="collect_form_30171425"></div>


        



<div class="related-info" style="margin-bottom:-10px;">
    <a name="intro"></a>
    
        
            
            
    <h2>
        <i class="">扫毒2天地对决的剧情简介</i>
              · · · · · ·
    </h2>

            <div class="indent" id="link-report">
                    
                        <span property="v:summary" class="">
                                　　毒品市场维持四分天下的格局已久，但自从地藏与墨西哥大毒枭跨境合作扩展势力，再加上一连串黑吃黑的动作，毒界变得风声鹤唳。另一方面，因儿时亲眼目睹父亲被毒品所毁而嫉毒如仇的慈善家兼金融巨子余顺天，正悬赏一亿歼灭香港最大毒贩，此举在社会上引起轩然大波。警员林正风本致力搜证拘捕地藏，毒贩却因巨额悬赏导致人身安全受威胁，他在执行保护毒贩的任务时深感无奈，但又被余顺天的出价所诱惑，陷入黑白正邪的心理挣扎。原来，余顺天和地藏有不可告人的同门关系，一场天地对决一触即发。
                        </span>
                        

            </div>
</div>


    








<div id="celebrities" class="celebrities related-celebrities">

  
    <h2>
        <i class="">扫毒2天地对决的演职员</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="/subject/30171425/celebrities">全部 21</a>
            )
            </span>
    </h2>


  <ul class="celebrities-list from-subject __oneline">
        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1274313/" title="邱礼涛 Herman Yau" class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p28346.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1274313/" title="邱礼涛 Herman Yau" class="name">邱礼涛</a></span>

      <span class="role" title="导演">导演</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1054424/" title="刘德华 Andy Lau" class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1378956633.91.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1054424/" title="刘德华 Andy Lau" class="name">刘德华</a></span>

      <span class="role" title="饰 余顺天">饰 余顺天</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1027577/" title="古天乐 Louis Koo" class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1421047436.79.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1027577/" title="古天乐 Louis Koo" class="name">古天乐</a></span>

      <span class="role" title="饰 地藏">饰 地藏</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1115415/" title="苗侨伟 Kiu Wai Miu" class="">
      <div class="avatar" style="background-image: url(https://img1.doubanio.com/view/celebrity/s_ratio_celebrity/public/p2247.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1115415/" title="苗侨伟 Kiu Wai Miu" class="name">苗侨伟</a></span>

      <span class="role" title="饰 林正风">饰 林正风</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1204410/" title="林嘉欣 Karena Lam" class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p2040.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1204410/" title="林嘉欣 Karena Lam" class="name">林嘉欣</a></span>

      <span class="role" title="饰 邹文凤">饰 邹文凤</span>

    </div>
  </li>


        
    

    
  
  <li class="celebrity">
    

  <a href="https://movie.douban.com/celebrity/1313563/" title="周秀娜 Chrissie Chaw" class="">
      <div class="avatar" style="background-image: url(https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p1361191369.34.webp)">
    </div>
  </a>

    <div class="info">
      <span class="name"><a href="https://movie.douban.com/celebrity/1313563/" title="周秀娜 Chrissie Chaw" class="name">周秀娜</a></span>

      <span class="role" title="饰 陈静美">饰 陈静美</span>

    </div>
  </li>


  </ul>
</div>


    


<link rel="stylesheet" href="https://img3.doubanio.com/f/verify/16c7e943aee3b1dc6d65f600fcc0f6d62db7dfb4/entry_creator/dist/author_subject/style.css">
<div id="author_subject" class="author-wrapper">
    <div class="loading"></div>
</div>
<script type="text/javascript">
    var answerObj = {
      ISALL: 'False',
      TYPE: 'movie',
      SUBJECT_ID: '30171425',
      USER_ID: 'None'
    }
</script>
<script type="text/javascript" src="https://img3.doubanio.com/f/movie/61252f2f9b35f08b37f69d17dfe48310dd295347/js/movie/lib/react/15.4/bundle.js"></script>
<script type="text/javascript" src="https://img3.doubanio.com/f/verify/ac140ef86262b845d2be7b859e352d8196f3f6d4/entry_creator/dist/author_subject/index.js"></script>


    









    
    <div id="related-pic" class="related-pic">
        
    
    
    <h2>
        <i class="">扫毒2天地对决的视频和图片</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30171425/trailer#trailer">预告片25</a>&nbsp;|&nbsp;<a href="/video/create?subject_id=30171425">添加视频评论</a>&nbsp;|&nbsp;<a href="https://movie.douban.com/subject/30171425/all_photos">图片208</a>&nbsp;·&nbsp;<a href="https://movie.douban.com/subject/30171425/mupload">添加</a>
            )
            </span>
    </h2>


        <ul class="related-pic-bd  ">
                <li class="label-trailer">
                    <a class="related-pic-video" href="https://movie.douban.com/trailer/249286/#content" title="预告片" style="background-image:url(https://img3.doubanio.com/img/trailer/medium/2561192416.jpg?1561978675)">
                    </a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2558844586/"><img src="https://img3.doubanio.com/view/photo/sqxs/public/p2558844586.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2560826998/"><img src="https://img1.doubanio.com/view/photo/sqxs/public/p2560826998.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2560914867/"><img src="https://img1.doubanio.com/view/photo/sqxs/public/p2560914867.webp" alt="图片" /></a>
                </li>
                <li>
                    <a href="https://movie.douban.com/photos/photo/2560827015/"><img src="https://img3.doubanio.com/view/photo/sqxs/public/p2560827015.webp" alt="图片" /></a>
                </li>
        </ul>
    </div>




    
    



<style type="text/css">
.award li { display: inline; margin-right: 5px }
.awards { margin-bottom: 20px }
.awards h2 { background: none; color: #000; font-size: 14px; padding-bottom: 5px; margin-bottom: 8px; border-bottom: 1px dashed #dddddd }
.awards .year { color: #666666; margin-left: -5px }
.mod { margin-bottom: 25px }
.mod .hd { margin-bottom: 10px }
.mod .hd h2 {margin:24px 0 3px 0}
</style>



    








    <div id="recommendations" class="">
        
        
    <h2>
        <i class="">喜欢这部电影的人也喜欢</i>
              · · · · · ·
    </h2>

        
    
    <div class="recommendations-bd">
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26995475/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2504801728.webp" alt="风再起时" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26995475/?from=subject-page" class="" >风再起时</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/30423193/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2559370302.webp" alt="使徒行者2：谍影行动" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/30423193/?from=subject-page" class="" >使徒行者2：谍影行动</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26351864/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2500366160.webp" alt="风林火山" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26351864/?from=subject-page" class="" >风林火山</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/30373723/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2550533959.webp" alt="限期破案" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/30373723/?from=subject-page" class="" >限期破案</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/27125418/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2497765052.webp" alt="杀破狼3" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/27125418/?from=subject-page" class="" >杀破狼3</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/24284175/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2181035841.webp" alt="九龙城寨" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/24284175/?from=subject-page" class="" >九龙城寨</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/26816090/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561779934.webp" alt="沉默的证人" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/26816090/?from=subject-page" class="" >沉默的证人</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/27163278/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2561542272.webp" alt="速度与激情：特别行动" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/27163278/?from=subject-page" class="" >速度与激情：特别行动</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/32579501/?from=subject-page" >
                    <img src="https://img3.doubanio.com/view/photo/s_ratio_poster/public/p2551253920.webp" alt="反贪风暴5" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/32579501/?from=subject-page" class="" >反贪风暴5</a>
            </dd>
        </dl>
        <dl class="">
            <dt>
                <a href="https://movie.douban.com/subject/30174085/?from=subject-page" >
                    <img src="https://img1.doubanio.com/view/photo/s_ratio_poster/public/p2551096308.webp" alt="怒火" class="" />
                </a>
            </dt>
            <dd>
                <a href="https://movie.douban.com/subject/30174085/?from=subject-page" class="" >怒火</a>
            </dd>
        </dl>
    </div>

    </div>



        


<script type="text/x-handlebar-tmpl" id="comment-tmpl">
    <div class="dummy-fold">
        {{#each comments}}
        <div class="comment-item" data-cid="id">
            <div class="comment">
                <h3>
                    <span class="comment-vote">
                            <span class="votes">{{votes}}</span>
                        <input value="{{id}}" type="hidden"/>
                        <a href="javascript:;" class="j {{#if ../if_logined}}a_vote_comment{{else}}a_show_login{{/if}}">有用</a>
                    </span>
                    <span class="comment-info">
                        <a href="{{user.path}}" class="">{{user.name}}</a>
                        {{#if rating}}
                        <span class="allstar{{rating}}0 rating" title="{{rating_word}}"></span>
                        {{/if}}
                        <span>
                            {{time}}
                        </span>
                        <p> {{content_tmpl content}} </p>
                    </span>
            </div>
        </div>
        {{/each}}
    </div>
</script>












    

    <div id="comments-section">
        <div class="mod-hd">
            
        <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow">
            <span>我要写短评</span>
        </a>

            
            
    <h2>
        <i class="">扫毒2天地对决的短评</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30171425/comments?status=P">全部 16356 条</a>
            )
            </span>
    </h2>

        </div>
        <div class="mod-bd">
                
    <div class="tab-hd">
        <a id="hot-comments-tab" href="comments" data-id="hot" class="on">热门</a>&nbsp;/&nbsp;
            <a id="new-comments-tab" href="comments?sort=time" data-id="new">最新</a>&nbsp;/&nbsp;
        <a id="following-comments-tab" href="follows_comments" data-id="following"  class="j a_show_login">好友</a>
    </div>

    <div class="tab-bd">
        <div id="hot-comments" class="tab">
            
    
        
        <div class="comment-item" data-cid="1847477039">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">731</span>
                <input value="1847477039" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/lingrui1995/" class="">凌睿</a>
                    <span>看过</span>
                    <span class="allstar30 rating" title="还行"></span>
                <span class="comment-time " title="2019-07-05 13:34:55">
                    2019-07-05
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">两个杨过和一个杨康的故事。
一个明面上是灭毒委员会会长、太平绅士，实则是为了私欲欺世盗名；
一个号称是正经商人、慈善大使，实则吃着人血馒头，赚的全是不干净的钱。
一个是伪君子，一个是真小人。
“扫毒”是表象，本质是黑吃黑。
电影的出发点是好的，充分揭露了某些上流人物人面兽心、道貌岸然的形象。
他们满嘴的道义，看不起毒贩，其实他们和毒贩没什么两样。
同时也探讨了毒贩到底有没有人权的问题。
抓毒贩是否...</span>
                <span class="hide-item full">两个杨过和一个杨康的故事。
一个明面上是灭毒委员会会长、太平绅士，实则是为了私欲欺世盗名；
一个号称是正经商人、慈善大使，实则吃着人血馒头，赚的全是不干净的钱。
一个是伪君子，一个是真小人。
“扫毒”是表象，本质是黑吃黑。
电影的出发点是好的，充分揭露了某些上流人物人面兽心、道貌岸然的形象。
他们满嘴的道义，看不起毒贩，其实他们和毒贩没什么两样。
同时也探讨了毒贩到底有没有人权的问题。
抓毒贩是否可以凌驾于法律之上？
以正义为名的犯罪，是正义，还是犯罪？
可惜影片没有深入探讨，本可以是一部关于权谋、关于人性的政治惊悚片，却被拍成了毫无深度的爆米花电影。
后半部分完全是为了打架而打架，一个是富商，一个是毒枭，一个比一个有钱，用得着他们亲自动手？
少一点动作戏，多一点拼计谋、拼智慧会更好。</span>
                <span class="expand">(<a href="javascript:;">展开</a>)</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1846941491">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">375</span>
                <input value="1846941491" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/165400923/" class="">Vince Smith</a>
                    <span>看过</span>
                    <span class="allstar30 rating" title="还行"></span>
                <span class="comment-time " title="2019-07-05 00:02:02">
                    2019-07-05
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">本片又名《港影老熟人们比谁更快领便当的游戏》。</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1847117855">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">334</span>
                <input value="1847117855" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/45991053/" class="">复制灵魂</a>
                    <span>看过</span>
                    <span class="allstar20 rating" title="较差"></span>
                <span class="comment-time " title="2019-07-05 05:11:40">
                    2019-07-05
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">这剧情尴尬到死，没想到古天乐和刘德华都这岁数了还在拍这种没营养的爆米花电影，难道这就是你们口中所谓的对港片的担当？</span>
                
                <a class="source-icon" href="https://www.douban.com/doubanapp/" target="_blank"><img src="https://img3.doubanio.com/f/shire/f62b2d2de3fc4a56d176b01cc3bbd47d2681fb38/pics/comment/android.png" title="发自Android" alt="Android" rel="v:image"/></a>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1846943523">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">235</span>
                <input value="1846943523" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/dkjune/" class="">ZeonGin Sou</a>
                    <span>看过</span>
                    <span class="allstar30 rating" title="还行"></span>
                <span class="comment-time " title="2019-07-05 00:03:21">
                    2019-07-05
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">刘德华化身蝙蝠侠，砸几个亿机关枪扫毒。纯港片，比混血追龙2强上N倍，连客串阵容都星光熠熠。遗憾的是，角色过于平面，情义苍白，人物关系虽有角力但联系疲软，还不如主题曲歌词有力，似有若无的反转还不如硬干到底。邱礼涛操刀，就不用有太多想象了，四平八稳完成剧本交差，节奏控制莫名其妙。纵深有限。PS：高潮一战的灵感应该是来自港铁经常出故障吧。</span>
        </p>
    </div>

        </div>
        
        <div class="comment-item" data-cid="1847241952">
            
    
    <div class="comment">
        <h3>
            <span class="comment-vote">
                <span class="votes">84</span>
                <input value="1847241952" type="hidden"/>
                <a href="javascript:;" class="j a_show_login" onclick="">有用</a>
            </span>
            <span class="comment-info">
                <a href="https://www.douban.com/people/3287454/" class="">ariella</a>
                    <span>看过</span>
                    <span class="allstar20 rating" title="较差"></span>
                <span class="comment-time " title="2019-07-05 10:01:25">
                    2019-07-05
                </span>
            </span>
        </h3>
        <p class="">
            
                <span class="short">剧本敷衍得厉害。平庸的烂片依旧是烂片。</span>
        </p>
    </div>

        </div>



                
                &gt; <a href="comments?sort=new_score&status=P" >更多短评16356条</a>
        </div>
        <div id="new-comments" class="tab">
            <div id="normal">
            </div>
            <div class="fold-hd hide">
                <a class="qa" href="/help/opinion#t2-q0" target="_blank">为什么被折叠？</a>
                <a class="btn-unfold" href="#">有一些短评被折叠了</a>
                <div class="qa-tip">
                    评论被折叠，是因为发布这条评论的帐号行为异常。评论仍可以被展开阅读，对发布人的账号不造成其他影响。如果认为有问题，可以<a href="https://help.douban.com/help/ask?category=movie">联系</a>豆瓣电影。
                </div>
            </div>
            <div class="fold-bd">
            </div>
            <span id="total-num"></span>
        </div>
        <div id="following-comments" class="tab">
            
    


        <div class="comment-item">
            你关注的人还没写过短评
        </div>

        </div>
    </div>


            
            
        </div>
    </div>



        

<link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/5b893040c800d5cd.css">

<section class="topics mod">
    <header>
        <h2>
            扫毒2天地对决的话题 · · · · · ·
            <span class="pl">( <span class="gallery_topics">全部 <span id="topic-count"></span> 条</span> )</span>
        </h2>
    </header>

    




<section class="subject-topics">
    <div class="topic-guide" id="topic-guide">
        <img class="ic_question" src="//img3.doubanio.com/f/ithildin/b1a3edea3d04805f899e9d77c0bfc0d158df10d5/pics/export/icon_question.png">
        <div class="tip_content">
            <div class="tip_title">什么是话题</div>
            <div class="tip_desc">
                <div>无论是一部作品、一个人，还是一件事，都往往可以衍生出许多不同的话题。将这些话题细分出来，分别进行讨论，会有更多收获。</div>
            </div>
        </div>
        <img class="ic_guide" src="//img3.doubanio.com/f/ithildin/529f46d86bc08f55cd0b1843d0492242ebbd22de/pics/export/icon_guide_arrow.png">
        <img class="ic_close" id="topic-guide-close" src="//img3.doubanio.com/f/ithildin/2eb4ad488cb0854644b23f20b6fa312404429589/pics/export/close@3x.png">
    </div>

    <div id="topic-items"></div>

    <script>
        window.subject_id = 30171425;
        window.join_label_text = '写影评参与';

        window.topic_display_count = 4;
        window.topic_item_display_count = 1;
        window.no_content_fun_call_name = "no_topic";

        window.guideNode = document.getElementById('topic-guide');
        window.guideNodeClose = document.getElementById('topic-guide-close');
    </script>
    
        <link rel="stylesheet" href="https://img3.doubanio.com/f/ithildin/f731c9ea474da58c516290b3a6b1dd1237c07c5e/css/export/subject_topics.css">
        <script src="https://img3.doubanio.com/f/ithildin/d3590fc6ac47b33c804037a1aa7eec49075428c8/js/export/moment-with-locales-only-zh.js"></script>
        <script src="https://img3.doubanio.com/f/ithildin/c600fdbe69e3ffa5a3919c81ae8c8b4140e99a3e/js/export/subject_topics.js"></script>

</section>

    <script>
        function no_topic(){
            $('#content .topics').remove();
        }
    </script>
</section>

<section class="reviews mod movie-content">
    <header>
        <a href="new_review" rel="nofollow" class="create-review comment_btn "
            data-isverify="False"
            data-verify-url="https://www.douban.com/accounts/phone/verify?redir=http://movie.douban.com/subject/30171425/new_review">
            <span>我要写影评</span>
        </a>
        <h2>
            扫毒2天地对决的影评 · · · · · ·
            <span class="pl">( <a href="reviews">全部 473 条</a> )</span>
        </h2>
    </header>

    

<style>
#gallery-topics-selection {
  position: fixed;
  width: 595px;
  padding: 40px 40px 33px 40px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 16px 0 rgba(0, 0, 0, 0.2);
  top: 50%;
  left: 50%;
  -webkit-transform: translate(-50%, -50%);
  transform: translate(-50%, -50%);
  z-index: 9999;
}
#gallery-topics-selection h1 {
  font-size: 18px;
  color: #007722;
  margin-bottom: 36px;
  padding: 0;
  line-height: 28px;
  font-weight: normal;
}
#gallery-topics-selection .gl_topics {
  border-bottom: 1px solid #dfdfdf;
  max-height: 298px;
  overflow-y: scroll;
}
#gallery-topics-selection .topic {
  margin-bottom: 24px;
}
#gallery-topics-selection .topic_name {
  font-size: 15px;
  color: #333;
  margin: 0;
  line-height: inherit;
}
#gallery-topics-selection .topic_meta {
  font-size: 13px;
  color: #999;
}
#gallery-topics-selection .topics_skip {
  display: block;
  cursor: pointer;
  font-size: 16px;
  color: #3377AA;
  text-align: center;
  margin-top: 33px;
}
#gallery-topics-selection .topics_skip:hover {
  background: transparent;
}
#gallery-topics-selection .close_selection {
  position: absolute;
  width: 30px;
  height: 20px;
  top: 46px;
  right: 40px;
  background: #fff;
  color: #999;
  text-align: right;
}
#gallery-topics-selection .close_selection:hover{
  background: #fff;
  color: #999;
}
</style>




        <div class="review_filter">
            <a href="javascript:;;" class="cur" data-sort="">热门</a href="javascript:;;"> /
            <a href="javascript:;;" data-sort="time">最新</a href="javascript:;;"> /
            <a href="javascript:;;" data-sort="follow">好友</a href="javascript:;;">
            
        </div>


        



<div class="review-list  ">
        
    

        
    
    <div data-cid="10286946">
        <div class="main review-item" id="10286946">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/155515686/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u155515686-2.jpg">
        </a>

        <a href="https://www.douban.com/people/155515686/" class="name">三只龙猫</a>

            <span class="allstar20 main-title-rating" title="较差"></span>

        <span content="2019-07-05" class="main-meta">2019-07-05 00:24:43</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10286946/">空顶影帝外衣的烂片</a></h2>

                <div id="review_10286946_short" class="review-short" data-rid="10286946">
                    <div class="short-content">

                        鉴于一些人声讨我在影院拍了一张照片，在这里深感抱歉，已经删除，之前没有在意过，不过以后不会了。还有那些吹水的，评分是最好的说明，开局7.1，24％的五星，20%的四星，一个下午就掉到6.8再看看现在的比例。足够说明了。你们都努力了我知道，不用再在我这蹦踏了。以下是影评...

                        &nbsp;(<a href="javascript:;" id="toggle-10286946-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10286946_full" class="hidden">
                    <div id="review_10286946_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10286946" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10286946">
                                257
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10286946" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10286946">
                                122
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10286946/#comments" class="reply ">588回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10287069">
        <div class="main review-item" id="10287069">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/geshu/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u1016423-5.jpg">
        </a>

        <a href="https://www.douban.com/people/geshu/" class="name">哥舒夜带刀</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2019-07-05" class="main-meta">2019-07-05 02:01:14</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10287069/">两位杨过决斗，老爹杨康给跪了</a></h2>

                <div id="review_10287069_short" class="review-short" data-rid="10287069">
                    <div class="short-content">

                        《扫毒2》应该是今年目前为止最好的港片，比《反贪风暴4》和《追龙2》都高出一截。 作为香港最为大众所知的cult片导演，邱礼涛三十多年来一直奋战在第一线，还拥有丰富的摄影、编剧、和剪辑经验，于各种类型片中游刃有余地切换，涉猎之广鲜有匹敌。那部让无数摇滚迷津津乐道的...

                        &nbsp;(<a href="javascript:;" id="toggle-10287069-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10287069_full" class="hidden">
                    <div id="review_10287069_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10287069" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10287069">
                                172
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10287069" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10287069">
                                33
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10287069/#comments" class="reply ">76回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10295969">
        <div class="main review-item" id="10295969">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/134527565/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u134527565-16.jpg">
        </a>

        <a href="https://www.douban.com/people/134527565/" class="name">优优飞扬</a>

            <span class="allstar30 main-title-rating" title="还行"></span>

        <span content="2019-07-09" class="main-meta">2019-07-09 07:13:44</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10295969/">血缘关系害死多年兄弟，《扫毒2》评分不该这么低</a></h2>

                <div id="review_10295969_short" class="review-short" data-rid="10295969">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        《扫毒2：天地对决》公映，票房不俗。 首日狂揽1.46亿，稳站院线片C位。 影片投资两亿港币，折合人民币有1.76亿。 公映第四天，票房5亿＋，我可以负责任地说，《扫毒2》是稳赚不赔了。 面对票房佳绩，带资进组的刘德华也会很欣慰（刘德华也是本片监制）。 事实上，《扫毒2》和...

                        &nbsp;(<a href="javascript:;" id="toggle-10295969-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10295969_full" class="hidden">
                    <div id="review_10295969_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10295969" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10295969">
                                15
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10295969" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10295969">
                                1
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10295969/#comments" class="reply ">9回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10293827">
        <div class="main review-item" id="10293827">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/180789499/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/user_normal.jpg">
        </a>

        <a href="https://www.douban.com/people/180789499/" class="name">。</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2019-07-08" class="main-meta">2019-07-08 03:19:10</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10293827/">（涉及剧透）你觉得不行，我觉得还行。先吐槽某些人</a></h2>

                <div id="review_10293827_short" class="review-short" data-rid="10293827">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        我看了觉得这电影还行，至于别人说什么极差、垃圾，也貌似有人家的道理。每个人都有自己的看法嘛，不需要强制人家和你一样。 但是某些人，看那种不需要脑子的电影看多了吗，剧情铺垫都看不懂，还说某些铺垫没用，然后洋洋洒洒一大段，还有上百人点赞，真是瞒不过你们这些懂行的...

                        &nbsp;(<a href="javascript:;" id="toggle-10293827-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10293827_full" class="hidden">
                    <div id="review_10293827_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10293827" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10293827">
                                17
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10293827" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10293827">
                                2
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10293827/#comments" class="reply ">13回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10294240">
        <div class="main review-item" id="10294240">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/176155637/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u176155637-1.jpg">
        </a>

        <a href="https://www.douban.com/people/176155637/" class="name">毛豆小龙虾</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2019-07-08" class="main-meta">2019-07-08 11:08:34</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10294240/">关于纸片人的使用方法</a></h2>

                <div id="review_10294240_short" class="review-short" data-rid="10294240">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        先说主角。结尾看的很奇怪，首先是地藏和余顺天。 余顺天恨毒品这一点剧中有刻画，也就这点觉得这个人物还算立体，但也仅此而已了。 结尾看起来很妙，却有点莫名其妙。你说我地藏深恨你余顺天二十多年兄弟一句自己到底做没做过都不问就砍了自己手指，甚至砍完了一句对不起都不...

                        &nbsp;(<a href="javascript:;" id="toggle-10294240-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10294240_full" class="hidden">
                    <div id="review_10294240_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10294240" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10294240">
                                13
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10294240" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10294240">
                                1
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10294240/#comments" class="reply ">3回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10290362">
        <div class="main review-item" id="10290362">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/lingrui1995/" class="avator">
            <img width="24" height="24" src="https://img1.doubanio.com/icon/u63688511-18.jpg">
        </a>

        <a href="https://www.douban.com/people/lingrui1995/" class="name">凌睿</a>

            <span class="allstar30 main-title-rating" title="还行"></span>

        <span content="2019-07-06" class="main-meta">2019-07-06 17:41:30</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10290362/">可惜只是半部佳作</a></h2>

                <div id="review_10290362_short" class="review-short" data-rid="10290362">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        《扫毒2》乍一看是《扫毒1》的续集，其实导演由陈木胜换为了邱礼涛，演员也由古天乐+张家辉+刘青云的组合换为了古天乐+刘德华+苗侨伟的组合。 （巧合的是，刘德华和古天乐都演过杨过，而苗侨伟演过杨康） 《扫毒2》是全新的角色、全新的剧情，和《扫毒1》没有任何关系。 这种现...

                        &nbsp;(<a href="javascript:;" id="toggle-10290362-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10290362_full" class="hidden">
                    <div id="review_10290362_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10290362" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10290362">
                                68
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10290362" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10290362">
                                17
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10290362/#comments" class="reply ">45回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10294350">
        <div class="main review-item" id="10294350">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/1026205/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u1026205-1.jpg">
        </a>

        <a href="https://www.douban.com/people/1026205/" class="name">wendy</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2019-07-08" class="main-meta">2019-07-08 12:06:42</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10294350/">我喜欢这种“恶趣味”</a></h2>

                <div id="review_10294350_short" class="review-short" data-rid="10294350">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        看了《扫毒2》，我有理由相信这个“2”不是“续集”的意思。 这真的是一部恶趣味满满的电影。 第一，“仇女”。 电影里有三个男主角：刘德华、古天乐、苗侨伟。 年轻时候的刘德华是一个黑社会的小混混，女朋友不想过这种朝不保夕的生活，于是离开了他。十几年后，这个前女友得...

                        &nbsp;(<a href="javascript:;" id="toggle-10294350-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10294350_full" class="hidden">
                    <div id="review_10294350_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10294350" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10294350">
                                13
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10294350" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10294350">
                                1
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10294350/#comments" class="reply ">5回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10290215">
        <div class="main review-item" id="10290215">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/174054285/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u174054285-1.jpg">
        </a>

        <a href="https://www.douban.com/people/174054285/" class="name">追梦人 Starcraf</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2019-07-06" class="main-meta">2019-07-06 16:29:59</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10290215/">人与人充满了怀疑，我们的世界不容易。。。</a></h2>

                <div id="review_10290215_short" class="review-short" data-rid="10290215">
                    <div class="short-content">
                            <p class="spoiler-tip">这篇影评可能有剧透</p>

                        从剧头地藏殴打散粉弟，到句尾地藏问余顺天：是不是自己被冤枉，以及片尾曲的歌词，基本上可以说地藏就是被冤枉的，可惜，一个人的命要走好很多时候需要天时地利人和。首先，如果余顺天能力劝余南（他亲叔）提供足够依据证明地藏确实卖毒，地藏可能也会幸免于难，可惜因为女友...

                        &nbsp;(<a href="javascript:;" id="toggle-10290215-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10290215_full" class="hidden">
                    <div id="review_10290215_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10290215" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10290215">
                                5
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10290215" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10290215">
                                4
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10290215/#comments" class="reply ">4回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10288792">
        <div class="main review-item" id="10288792">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/147790959/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u147790959-1.jpg">
        </a>

        <a href="https://www.douban.com/people/147790959/" class="name">ReJunic</a>

            <span class="allstar30 main-title-rating" title="还行"></span>

        <span content="2019-07-05" class="main-meta">2019-07-05 22:15:48</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10288792/">一直都是这些人在扛，真的快顶不住了！</a></h2>

                <div id="review_10288792_short" class="review-short" data-rid="10288792">
                    <div class="short-content">

                        纯正的港片，带着希望能给我点惊喜的去看、看完略失望，比起第一部还是差的多！ 古仔演技稍微在线，华仔感觉还是中规中矩。为什么这种续集要换导演呢，搞不懂，感觉邱礼涛还是零几年那种拍黑帮片的套路，一个大佬郑则仕，后面陆续出现嘻哈侠欧阳靖、应采儿、还有那个谁（林忆莲...

                        &nbsp;(<a href="javascript:;" id="toggle-10288792-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10288792_full" class="hidden">
                    <div id="review_10288792_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10288792" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10288792">
                                18
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10288792" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10288792">
                                7
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10288792/#comments" class="reply ">12回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>

        
    
    <div data-cid="10293792">
        <div class="main review-item" id="10293792">

            
    
    <header class="main-hd">
        <a href="https://www.douban.com/people/134194726/" class="avator">
            <img width="24" height="24" src="https://img3.doubanio.com/icon/u134194726-12.jpg">
        </a>

        <a href="https://www.douban.com/people/134194726/" class="name">叶凌生</a>

            <span class="allstar40 main-title-rating" title="推荐"></span>

        <span content="2019-07-08" class="main-meta">2019-07-08 02:02:32</span>


    </header>


            <div class="main-bd">

                <h2><a href="https://movie.douban.com/review/10293792/">《扫毒2》中的1234|商业片是娱乐大众，并不是教化众生，宽容一点！</a></h2>

                <div id="review_10293792_short" class="review-short" data-rid="10293792">
                    <div class="short-content">

                        经历了《追龙2》、《反弹风暴2+2》的绝望后，2019年下半年开启之月，抱着对港片救赎的期待，还是买了《扫毒2》的影票。 观影结束，期待依然是期待…… 整部影片剧情依然延续着香港警匪缉毒片一贯手法，枪战、飞车、祸及家人、主人公进行复仇。当然片子中刘德华（余顺天）的复仇...

                        &nbsp;(<a href="javascript:;" id="toggle-10293792-copy" class="unfold" title="展开">展开</a>)
                    </div>
                </div>

                <div id="review_10293792_full" class="hidden">
                    <div id="review_10293792_full_content" class="full-content"></div>
                </div>

                <div class="action">
                    <a href="javascript:;" class="action-btn up" data-rid="10293792" title="有用">
                        <img src="https://img3.doubanio.com/f/zerkalo/536fd337139250b5fb3cf9e79cb65c6193f8b20b/pics/up.png" />
                        <span id="r-useful_count-10293792">
                                6
                        </span>
                    </a>
                    <a href="javascript:;" class="action-btn down" data-rid="10293792" title="没用">
                        <img src="https://img3.doubanio.com/f/zerkalo/68849027911140623cf338c9845893c4566db851/pics/down.png" />
                        <span id="r-useless_count-10293792">
                                1
                        </span>
                    </a>
                    <a href="https://movie.douban.com/review/10293792/#comments" class="reply ">0回应</a>

                    <a href="javascript:;;" class="fold hidden">收起</a>
                </div>
            </div>
        </div>
    </div>




    

    

    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/666ac852ec3a604e.js"></script>
    <!-- COLLECTED CSS -->
</div>








            <p class="pl">
                &gt;
                <a href="reviews">
                    更多影评473篇
                </a>
            </p>
</section>

<!-- COLLECTED JS -->

    <br/>

        <div class="section-discussion">
                
                <div class="mod-hd">
                        <a class="comment_btn j a_show_login" href="https://www.douban.com/register?reason=review" rel="nofollow"><span>添加新讨论</span></a>
                    
    <h2>
        讨论区
         &nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;&nbsp; &middot;
    </h2>

                </div>
                
  <table class="olt"><tr><td><td><td><td></tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30171425/discussion/616361049/" title="这种垃圾片过6分，豆瓣已经沦陷了">这种垃圾片过6分，豆瓣已经沦陷了</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/187084036/">多喝六个核弹</a></td>
          <td class="pl"><span>3 回应</span></td>
          <td class="pl"><span>2019-07-09</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30171425/discussion/616356414/" title="我很喜欢最后地铁站飙车戏">我很喜欢最后地铁站飙车戏</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/146722298/">Louis Koo</a></td>
          <td class="pl"><span>18 回应</span></td>
          <td class="pl"><span>2019-07-09</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30171425/discussion/616361485/" title="兄弟不怀疑的国语版是什么？">兄弟不怀疑的国语版是什么？</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/168693069/">小锅红米线</a></td>
          <td class="pl"><span>1 回应</span></td>
          <td class="pl"><span>2019-07-09</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30171425/discussion/616357695/" title="刘德华的演技被捧过分了吧">刘德华的演技被捧过分了吧</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/76545433/">杀死金鱼</a></td>
          <td class="pl"><span>69 回应</span></td>
          <td class="pl"><span>2019-07-09</span></td>
        </tr>
        
        <tr>
          <td class="pl"><a href="https://movie.douban.com/subject/30171425/discussion/616360665/" title="欧阳靖他两卖的蘑菇也算毒品吗">欧阳靖他两卖的蘑菇也算毒品吗</a></td>
          <td class="pl"><span>来自</span><a href="https://www.douban.com/people/WiliamCCC/">卫斯豪</a></td>
          <td class="pl"><span>5 回应</span></td>
          <td class="pl"><span>2019-07-09</span></td>
        </tr>
  </table>

                <p class="pl" align="right">
                    <a href="/subject/30171425/discussion/" rel="nofollow">
                        &gt; 去这部影片的讨论区（全部320条）
                    </a>
                </p>
        </div>

        
    
        
                





<div id="askmatrix">
    <div class="mod-hd">
        <h2>
            关于《扫毒2天地对决》的问题
            · · · · · ·
            <span class="pl">
                (<a href='https://movie.douban.com/subject/30171425/questions/?from=subject'>
                    全部4个
                </a>)
            </span>
        </h2>


        
    
    <a class='j a_show_login comment_btn'
        href='https://movie.douban.com/subject/30171425/questions/ask/?from=subject'>我来提问</a>

    </div>

    <div class="mod-bd">
        <ul class="">
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/30171425/questions/837642/?from=subject" class="">
                            经典问题重提 ——《失孤》刘德华是怎么把摩托车骑到火车站台上去的？
                    </a>
                </span>
                <span class="meta">
                    0人回答
                </span>
            </li>
            <li class="">
                <span class="tit">
                    <a href="https://movie.douban.com/subject/30171425/questions/837653/?from=subject" class="">
                            有没有人和我一样觉得电影中刘德华老婆超漂亮、超有气质？
                    </a>
                </span>
                <span class="meta">
                    0人回答
                </span>
            </li>
        </ul>

        <p>&gt;
            <a href='https://movie.douban.com/subject/30171425/questions/?from=subject'>
                全部4个问题
            </a>
        </p>

    </div>
</div>



            


    <script type="text/javascript">
        $(function(){if($.browser.msie && $.browser.version == 6.0){
            var $info = $('#info'),
                maxWidth = parseInt($info.css('max-width'));
            if($info.width() > maxWidth) {
                $info.width(maxWidth);
            }
        }});
    </script>


            </div>
            <div class="aside">
                


    







            






<div class="ticket">
        <a class="ticket-btn" href="">购票</a>
</div>



    <!-- douban ad begin -->
    <div id="dale_movie_subject_top_right"></div>
    <div id="dale_movie_subject_top_middle"></div>
    <!-- douban ad end -->

    



<style type="text/css">
    .m4 {margin-bottom:8px; padding-bottom:8px;}
    .movieOnline {background:#FFF6ED; padding:10px; margin-bottom:20px;}
    .movieOnline h2 {margin:0 0 5px;}
    .movieOnline .sitename {line-height:2em; width:160px;}
    .movieOnline td,.movieOnline td a:link,.movieOnline td a:visited{color:#666;}
    .movieOnline td a:hover {color:#fff;}
    .link-bt:link,
    .link-bt:visited,
    .link-bt:hover,
    .link-bt:active {margin:5px 0 0; padding:2px 8px; background:#a8c598; color:#fff; -moz-border-radius: 3px; -webkit-border-radius: 3px; border-radius: 3px; display:inline-block;}
</style>



    







    
    <div class="tags">
        
        
    <h2>
        <i class="">豆瓣成员常用的标签</i>
              · · · · · ·
    </h2>

        <div class="tags-body">
                <a href="/tag/香港" class="">香港</a>
                <a href="/tag/犯罪" class="">犯罪</a>
                <a href="/tag/警匪" class="">警匪</a>
                <a href="/tag/缉毒" class="">缉毒</a>
                <a href="/tag/动作" class="">动作</a>
                <a href="/tag/毒品" class="">毒品</a>
                <a href="/tag/华仔" class="">华仔</a>
                <a href="/tag/2019" class="">2019</a>
        </div>
    </div>


    <div id="dale_movie_review_right_guess_you_like"></div>
    <div id="dale_movie_subject_inner_middle"></div>
    <div id="dale_movie_subject_download_middle"></div>
        








<div id="subject-doulist">
    
    
    <h2>
        <i class="">以下豆列推荐</i>
              · · · · · ·
            <span class="pl">
            (
                <a href="https://movie.douban.com/subject/30171425/doulists">全部</a>
            )
            </span>
    </h2>


    
    <ul>
            <li>
                <a href="https://www.douban.com/doulist/1504454/" target="_blank">ღ♩♪生活有这些期待很有动力♫♬ღ</a>
                <span>(freedom♪)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1434921/" target="_blank">2019—2021值得关注的华语电影</a>
                <span>(closer)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/3519267/" target="_blank">2010年代最经典的100部电影</a>
                <span>(架构师)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/42121383/" target="_blank">2019年所有值得关注的电影</a>
                <span>(皮皮鸦)</span>
            </li>
            <li>
                <a href="https://www.douban.com/doulist/1714855/" target="_blank">影像香港</a>
                <span>(林伶仃)</span>
            </li>
    </ul>

</div>

            








<div id="subject-others-interests">
    
    
    <h2>
        <i class="">谁在看这部电影</i>
              · · · · · ·
    </h2>

    
    <ul class="">
            
            <li class="">
                <a href="https://www.douban.com/people/166003077/" class="others-interest-avatar">
                    <img src="https://img3.doubanio.com/icon/u166003077-2.jpg" class="pil" alt="文辰">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/166003077/" class="">文辰</a>
                    <div class="">
                        刚刚
                        看过
                        <span class="allstar40" title="推荐"></span>
                    </div>
                </div>
            </li>
            
            <li class="">
                <a href="https://www.douban.com/people/130168103/" class="others-interest-avatar">
                    <img src="https://img3.doubanio.com/icon/u130168103-1.jpg" class="pil" alt="非你丶不可">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/130168103/" class="">非你丶不可</a>
                    <div class="">
                        刚刚
                        看过
                        <span class="allstar30" title="还行"></span>
                    </div>
                </div>
            </li>
            
            <li class="">
                <a href="https://www.douban.com/people/echojules/" class="others-interest-avatar">
                    <img src="https://img3.doubanio.com/icon/u15026209-20.jpg" class="pil" alt="echo">
                </a>
                <div class="others-interest-info">
                    <a href="https://www.douban.com/people/echojules/" class="">echo</a>
                    <div class="">
                        刚刚
                        看过
                        <span class="allstar30" title="还行"></span>
                    </div>
                </div>
            </li>
    </ul>

    
    <div class="subject-others-interests-ft">
        
            <a href="https://movie.douban.com/subject/30171425/collections">42919人看过</a>
                &nbsp;/&nbsp;
            <a href="https://movie.douban.com/subject/30171425/wishes">23774人想看</a>
    </div>

</div>



    
    

<!-- douban ad begin -->
<div id="dale_movie_subject_middle_right"></div>
<script type="text/javascript">
    (function (global) {
        if(!document.getElementsByClassName) {
            document.getElementsByClassName = function(className) {
                return this.querySelectorAll("." + className);
            };
            Element.prototype.getElementsByClassName = document.getElementsByClassName;

        }
        var articles = global.document.getElementsByClassName('article'),
            asides = global.document.getElementsByClassName('aside');

        if (articles.length > 0 && asides.length > 0 && articles[0].offsetHeight >= asides[0].offsetHeight) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_middle_right');
        }
    })(this);
</script>
<!-- douban ad end -->



    <br/>

    
<p class="pl">订阅扫毒2天地对决的评论: <br/><span class="feed">
    <a href="https://movie.douban.com/feed/subject/30171425/reviews"> feed: rss 2.0</a></span></p>


            </div>
            <div class="extra">
                
    
<!-- douban ad begin -->
<div id="dale_movie_subject_bottom_super_banner"></div>
<script type="text/javascript">
    (function (global) {
        var body = global.document.body,
            html = global.document.documentElement;

        var height = Math.max(body.scrollHeight, body.offsetHeight, html.clientHeight, html.scrollHeight, html.offsetHeight);
        if (height >= 2000) {
            (global.DoubanAdSlots = global.DoubanAdSlots || []).push('dale_movie_subject_bottom_super_banner');
        }
    })(this);
</script>
<!-- douban ad end -->


            </div>
        </div>
    </div>

        
    <div id="footer">
            <div class="footer-extra"></div>
        
<span id="icp" class="fleft gray-link">
    &copy; 2005－2019 douban.com, all rights reserved 北京豆网科技有限公司
</span>

<a href="https://www.douban.com/hnypt/variformcyst.py" style="display: none;"></a>

<span class="fright">
    <a href="https://www.douban.com/about">关于豆瓣</a>
    · <a href="https://www.douban.com/jobs">在豆瓣工作</a>
    · <a href="https://www.douban.com/about?topic=contactus">联系我们</a>
    · <a href="https://www.douban.com/about/legal">法律声明</a>
    
    · <a href="https://help.douban.com/?app=movie" target="_blank">帮助中心</a>
    · <a href="https://www.douban.com/doubanapp/">移动应用</a>
    · <a href="https://www.douban.com/partner/">豆瓣广告</a>
</span>

    </div>

    </div>
    <script type="text/javascript" src="https://img3.doubanio.com/misc/mixed_static/a5b18df53993d62.js"></script>
        
        
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/shire/8377b9498330a2e6f056d863987cc7a37eb4d486/css/ui/dialog.css" />
    <link rel="stylesheet" type="text/css" href="https://img3.doubanio.com/f/movie/4aca95d66d37ec0712b3d19973b5d8feb75f2f05/css/movie/mod/reg_login_pop.css" />
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/77323ae72a612bba8b65f845491513ff3329b1bb/js/do.js" data-cfg-autoload="false"></script>
    <script type="text/javascript" src="https://img3.doubanio.com/f/shire/383a6e43f2108dc69e3ff2681bc4dc6c72a5ffb0/js/ui/dialog.js"></script>
    <script type="text/javascript">
        var HTTPS_DB='https://www.douban.com';
var account_pop={open:function(o,e){e?referrer="?referrer="+encodeURIComponent(e):referrer="?referrer="+window.location.href;var n="",i="",t=448;n="用户登录",i="https://accounts.douban.com/passport/login_popup?source=movie";var r=document.location.protocol+"//"+document.location.hostname,a=dui.Dialog({width:340,title:n,height:t,cls:"account_pop",isHideTitle:!0,modal:!0,content:"<iframe scrolling='no' frameborder='0' width='340' height='"+t+"' src='"+i+"' name='"+r+"'></iframe>"},!0),c=a.node;if(c.undelegate(),c.delegate(".dui-dialog-close","click",function(){var o=$("body");o.find("#login_msk").hide(),o.find(".account_pop").remove()}),$(window).width()<478){var d="";"reg"===o?d=HTTPS_DB+"/accounts/register"+referrer:"login"===o&&(d=HTTPS_DB+"/accounts/login"+referrer),window.location.href=d}else a.open();$(window).bind("message",function(o){"https://accounts.douban.com"===o.originalEvent.origin&&(c.find("iframe").css("height",o.originalEvent.data),c.height(o.originalEvent.data),a.update())})}};Douban&&Douban.init_show_login&&(Douban.init_show_login=function(o){var e=$(o);e.click(function(){var o=e.data("ref")||"";return account_pop.open("login",o),!1})}),Do(function(){$("body").delegate(".pop_register","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("reg",e),!1}),$("body").delegate(".pop_login","click",function(o){o.preventDefault();var e=$(this).data("ref")||"";return account_pop.open("login",e),!1})});
    </script>

    
    
    
    




    
<script type="text/javascript">
    (function (global) {
        var newNode = global.document.createElement('script'),
            existingNode = global.document.getElementsByTagName('script')[0],
            adSource = '//erebor.douban.com/',
            userId = '',
            browserId = 'IZkjkS_KGG4',
            criteria = '7:毒品|7:黑帮|7:林嘉欣|7:苗侨伟|7:李赏|7:郑则仕|7:动作|7:应采儿|7:剧情|7:华仔|7:李灿森|7:犯罪|7:陈家乐|7:恭硕良|7:2019|7:林家栋|7:警匪|7:陈佩诗|7:张国强|7:周秀娜|7:刘德华|7:缉毒|7:狼森|7:悬疑|7:欧阳靖|7:古天乐|7:张竣杰|7:卫诗雅|7:邱礼涛|7:香港|3:/subject/30171425/?from=showing',
            preview = '',
            debug = false,
            adSlots = ['dale_movie_subject_top_icon', 'dale_movie_subject_top_right', 'dale_movie_subject_top_middle', 'dale_movie_subject_inner_middle', 'dale_movie_subject_download_middle', 'dale_movie_review_right_guess_you_like'];

        global.DoubanAdRequest = {src: adSource, uid: userId, bid: browserId, crtr: criteria, prv: preview, debug: debug};
        global.DoubanAdSlots = (global.DoubanAdSlots || []).concat(adSlots);

        newNode.setAttribute('type', 'text/javascript');
        newNode.setAttribute('src', 'https://img3.doubanio.com/f/adjs/5ec0e90cac83524d8c3e7f7bf95260c9c1875bca/ad.release.js');
        newNode.setAttribute('async', true);
        existingNode.parentNode.insertBefore(newNode, existingNode);
    })(this);
</script>











    
  









<script type="text/javascript">
var _paq = _paq || [];
_paq.push(['trackPageView']);
_paq.push(['enableLinkTracking']);
(function() {
    var p=(('https:' == document.location.protocol) ? 'https' : 'http'), u=p+'://fundin.douban.com/';
    _paq.push(['setTrackerUrl', u+'piwik']);
    _paq.push(['setSiteId', '100001']);
    var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
    g.type='text/javascript';
    g.defer=true;
    g.async=true;
    g.src=p+'://img3.doubanio.com/dae/fundin/piwik.js';
    s.parentNode.insertBefore(g,s);
})();
</script>

<script type="text/javascript">
var setMethodWithNs = function(namespace) {
  var ns = namespace ? namespace + '.' : ''
    , fn = function(string) {
        if(!ns) {return string}
        return ns + string
      }
  return fn
}

var gaWithNamespace = function(fn, namespace) {
  var method = setMethodWithNs(namespace)
  fn.call(this, method)
}

var _gaq = _gaq || []
  , accounts = [
      { id: 'UA-7019765-1', namespace: 'douban' }
    , { id: 'UA-7019765-19', namespace: '' }
    ]
  , gaInit = function(account) {
      gaWithNamespace(function(method) {
        gaInitFn.call(this, method, account)
      }, account.namespace)
    }
  , gaInitFn = function(method, account) {
      _gaq.push([method('_setAccount'), account.id]);
      _gaq.push([method('_setSampleRate'), '5']);

      
  _gaq.push([method('_addOrganic'), 'google', 'q'])
  _gaq.push([method('_addOrganic'), 'baidu', 'wd'])
  _gaq.push([method('_addOrganic'), 'soso', 'w'])
  _gaq.push([method('_addOrganic'), 'youdao', 'q'])
  _gaq.push([method('_addOrganic'), 'so.360.cn', 'q'])
  _gaq.push([method('_addOrganic'), 'sogou', 'query'])
  if (account.namespace) {
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣'])
    _gaq.push([method('_addIgnoredOrganic'), 'douban'])
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣网'])
    _gaq.push([method('_addIgnoredOrganic'), 'www.douban.com'])
  }

      if (account.namespace === 'douban') {
        _gaq.push([method('_setDomainName'), '.douban.com'])
      }

        _gaq.push([method('_setCustomVar'), 1, 'responsive_view_mode', 'desktop', 3])

        _gaq.push([method('_setCustomVar'), 2, 'login_status', '0', 2]);

      _gaq.push([method('_trackPageview')])
    }

for(var i = 0, l = accounts.length; i < l; i++) {
  var account = accounts[i]
  gaInit(account)
}


;(function() {
    var ga = document.createElement('script');
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    ga.setAttribute('async', 'true');
    document.documentElement.firstChild.appendChild(ga);
})()
</script>








      
    

    <!-- brand39-docker-->

  <script>_SPLITTEST=''</script>
</body>

</html>


`
	c.Ctx.WriteString(models.GetMovieDirector(sMovieHtml) + " | ")
	c.Ctx.WriteString(models.GetMovieName(sMovieHtml) + " | ")
	c.Ctx.WriteString(models.GetMovieMainCharacters(sMovieHtml))

}
