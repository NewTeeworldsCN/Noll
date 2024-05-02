<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>{{ .Viewer.ShowName }}论坛 </title>
</head>

<body>
  {{ if .Data }}
  {{ template "HeaderTemplate" . }}
  <div class="clearfix">
    <div>
      <h1>分类</h1>
      {{ template "CategoryGroupTemplate" .Categories }}
    </div>
    <div>
      <h1>标签</h1>
      {{ template "LabelGroupTemplate" .Labels }}
    </div>
  </div>
  {{ end }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>