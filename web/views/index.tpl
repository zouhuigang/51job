<!Doctype html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>PikaPika简历|Sunteng</title>
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
    <li role="presentation" class="active"><a href="/">首页</a></li>
    <li role="presentation"><a href="http://127.0.0.1:8099" target="_blank">简历爬取</a></li>
    <li role="presentation"><a href="/intro" target="_blank">使用说明</a></li>
</ul>
</nav>

<div id="body" class="container-fluid">
  <div> 
    <h1>简历 <small>PikaPika|{{.Count}}条</small>
<br />
    {{if .K}}
    <blockquote>
       <p>目前状态：{{.K.Keyword}}-----{{.K.Kind}}-----{{.K.Address}}------创建{{dateformat .K.Created "2006-01-02"}}------更新{{dateformat .K.Updated "2006-01-02"}}-------检索{{.K.Time51}}次
       </p>
       </blockquote>
    {{end}}
    </h1>
  </div>
  <div class="row">
      <div class="col-md-7"> 
    <select class="form-control" id="jinhan">
    <option value=""></option>
    {{range $index, $u := .Keyword}}
    <option value="{{$u.Id}}">{{$u.Keyword}}-----{{$u.Kind}}-----{{$u.Address}}------创建{{dateformat $u.Created "2006-01-02"}}------更新{{dateformat $u.Updated "2006-01-02"}}-------
    检索{{$u.Time51}}次</option>
    {{end}}
    </select>
    </div>
    <div class="col-md-2"> 
      <select class="form-control" id="sex">
      <option value=""></option>
      <option value="1">妹纸</option>
      <option value="0">屌丝</option>
      </select>
    </div>
    <div class="col-md-3">
       <button class="btn btn-primary btn-default" style="width:100%" onclick="search()">搜索</button>
       <script type="text/javascript">
           function search(){
            var id=$("#jinhan").val()
            var sex=$("#sex").val()
            window.location="/?k="+id+"&s="+sex
           }
       </script>
    </div>  
  </div>
  <div class="table-responsive row">
 <div class="col-md-12">
     {{if .paginator.HasPages}}
     <ul class="pagination">
         {{if .paginator.HasPrev}}
             <li><a href="{{.paginator.PageLinkFirst}}">{{ i18n .Lang "paginator.first_page"}}</a></li>
             <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
         {{else}}
             <li class="disabled"><a>{{ i18n .Lang "paginator.first_page"}}</a></li>
             <li class="disabled"><a>&laquo;</a></li>
         {{end}}
         {{range $index, $page := .paginator.Pages}}
             <li{{if $.paginator.IsActive .}} class="active"{{end}}>
                 <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
             </li>
         {{end}}
         {{if .paginator.HasNext}}
             <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
             <li><a href="{{.paginator.PageLinkLast}}">{{ i18n .Lang "paginator.last_page"}}</a></li>
         {{else}}
             <li class="disabled"><a>&raquo;</a></li>
             <li class="disabled"><a>{{ i18n .Lang "paginator.last_page"}}</a></li>
         {{end}}
     </ul>
     {{end}}
</div>
  <div class="col-md-12">
    <table class="table table-striped  table-hover">
    <thead>
      <tr class="info">
        <th>序号</th>
        <th>编号</th>
        <th>工作年限</th>
        <th>年龄</th>
        <th>性别</th>
        <th>地址</th>
        <th>专业</th>
        <th>文凭</th>
        <th>简历更新时间</th>
        <th>第一次抓取时间</th>
        <th>最后一次抓取时间</th>
      </tr>
    </thead>
    <tbody>
    {{range $index, $u := .User}}
        <tr>
        <td>{{$index}}</td>
        <td>
        <a href="/job?k={{$.Pk}}&id={{$u.Id51}}" target="__blank" alt="{{$u.Id51}}">
        <img src="/img?id={{$u.Id51}}" style="width:150px;height:150px" alt="{{$u.Id51}}"/></a>

        </td>
        <td>{{$u.Jobyear}}</td>
        <td>{{$u.Age}}</td>
        <td>{{$u.Sex}}</td>
        <td>{{$u.Address}}</td>
        <td>{{$u.Major}}</td>
        <td>{{$u.Study}}</td>
        <td>{{dateformat $u.Date51 "2006-01-02"}}</td>
        <td>{{dateformat $u.Created "2006-01-02 15:04:05"}}</td>
        <td>{{dateformat $u.Updated "2006-01-02 15:04:05"}}</td>
        </tr>
    {{end}}
    </tbody>
    </table>
  </div>
  <div class="col-md-12">
     {{if .paginator.HasPages}}
     <ul class="pagination">
         {{if .paginator.HasPrev}}
             <li><a href="{{.paginator.PageLinkFirst}}">{{ i18n .Lang "paginator.first_page"}}</a></li>
             <li><a href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
         {{else}}
             <li class="disabled"><a>{{ i18n .Lang "paginator.first_page"}}</a></li>
             <li class="disabled"><a>&laquo;</a></li>
         {{end}}
         {{range $index, $page := .paginator.Pages}}
             <li{{if $.paginator.IsActive .}} class="active"{{end}}>
                 <a href="{{$.paginator.PageLink $page}}">{{$page}}</a>
             </li>
         {{end}}
         {{if .paginator.HasNext}}
             <li><a href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
             <li><a href="{{.paginator.PageLinkLast}}">{{ i18n .Lang "paginator.last_page"}}</a></li>
         {{else}}
             <li class="disabled"><a>&raquo;</a></li>
             <li class="disabled"><a>{{ i18n .Lang "paginator.last_page"}}</a></li>
         {{end}}
     </ul>
     {{end}}
</div>
</div>
<div id="footer"></div>
 <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="/static/js/bootstrap.min.js"></script>
</body>

</html>