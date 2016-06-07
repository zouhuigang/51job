<!Doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>编号{{.User.User.Id51}}|PikaPika</title>
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">


    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="//cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="//cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->

    <style type="text/css">
      body { padding-top: 20px; }
      .starter-template {
          padding: 40px 15px;
          text-align: center;
        }
        th{
            text-align: center
        }
    </style>
</head>
<body class="starter-template">
<nav class="navbar">
<ul class="nav nav-tabs">
    <li role="presentation"><a href="/">首页</a></li>
    <li role="presentation"><a href="http://127.0.0.1:8099" target="_blank">简历爬取</a></li>
    <li role="presentation"><a href="/intro" target="_blank">使用说明</a></li>
</ul>
</nav>

<div id="body" class="container-fluid">
 <div class="col-md-12">
    <table class="table table-striped  table-hover">
    <thead>
      <tr class="info">
        <th>编号</th>
        <th>工作年限</th>
        <th>年龄</th>
        <th>姓名</th>
        <th>地址</th>
        <th>专业</th>
        <th>文凭</th>
      </tr>
    </thead>
    <tbody>
        <tr>
        <td>{{.User.User.Id51}}</td>
        <td>{{.User.User.Jobyear}}</td>
        <td>{{.User.User.Age}}</td>
        <td>{{.User.User.Sex}}</td>
        <td>{{.User.User.Address}}</td>
        <td>{{.User.User.Major}}</td>
        <td>{{.User.User.Study}}</td>
        </tr>
    </tbody>
    </table>
  </div>
  <div class="col-md-12">
    <table class="table table-striped  table-hover">
    <thead>
      <tr class="info">
        <th>序号</th>
        <th>搜索关键字</th>
        <th>匹配方式</th>
        <th>搜索地点</th>
        <th>本地保存地址</th>
        <th>简历更新时间</th>
        <th>简历抓取时间</th>
      </tr>
    </thead>
    <tbody>
    {{range $index, $u := .UserHistroy}}
        <tr>
        <td>{{$index}}</td>
        <td>{{$u.Keyword.Keyword}}</td>
        <td>{{$u.Keyword.Kind}}</td>
        <td>{{$u.Keyword.Address}}</td>
		<td><a href="/download?id={{$u.FileAddress}}" target="_blank">{{$u.FileAddress}}</a></td>
        <td>{{dateformat $u.Date51 "2006-01-02"}}</td>
        <td>{{dateformat $u.Created "2006-01-02 15:04:05"}}</td>
        </tr>
    {{end}}
    </tbody>
    </table>
  </div>
    <div class="col-md-8 col-md-offset-2">
  	<nav>
	  <ul class="pager">
        {{if eq "" .Brotherp}}
        {{else}}
	    <li class="previous"><a href="/job?k={{.Pk}}&id={{.Brotherp}}"><span aria-hidden="true">&larr;</span> 上一个</a></li>
        {{end}}
        
        {{if eq "" .Brothern}}
        {{else}}
        <li class="next"><a href="/job?k={{.Pk}}&id={{.Brothern}}">下一个<span aria-hidden="true">&rarr;</span></a></li>
        {{end}}
	  </ul>
	</nav>
	</div>
    <div class="col-md-8 col-md-offset-2">
	{{str2html .User.Content}}
    </div>
    <div class="col-md-8 col-md-offset-2">
  	<nav>
      <ul class="pager">
        {{if eq "" .Brotherp}}
        {{else}}
        <li class="previous"><a href="/job?k={{.Pk}}&id={{.Brotherp}}"><span aria-hidden="true">&larr;</span> 上一个</a></li>
        {{end}}
        
        {{if eq "" .Brothern}}
        {{else}}
        <li class="next"><a href="/job?k={{.Pk}}&id={{.Brothern}}">下一个<span aria-hidden="true">&rarr;</span></a></li>
        {{end}}
      </ul>
	</nav>
	</div>
 </div>


 <div id="footer">
 	
 </div>
 <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
    <script>
   var s=$("a img").parent()
   s.attr("href","/imgp?id="+{{.User.User.Id51}})
   s.attr("target","_blank")
    $("a img").attr("src","/imgp?id="+{{.User.User.Id51}})

    </script>
</body>

</html>