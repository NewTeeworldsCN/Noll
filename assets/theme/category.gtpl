<!DOCTYPE html>
<html lang="zh-CN">

<head>
  {{ template "HeadTemplate" .Viewer }}
  <title>{{ .Data.Name }} Category —— {{ .Viewer.ShowName }}论坛 </title>
</head>

<body>
  {{ template "HeaderTemplate" . }}
  <h1> {{ .Data.Name }} </h1>
  {{ .Data.Description }}
  {{ template "DiscussionGroupTemplate" .Data.Discussions }}
  {{ template "footerTemplate" .Viewer }}
</body>

</html>