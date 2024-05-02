{{define "HeaderTemplate"}}
<header>
<ul class="ul" style="margin-left: -10px;">
  <li class="li"><a href='{{ url "/" }}'>
      <h3 style="display: contents;">{{ .Viewer.ShowName }}论坛</h3>
    </a></li>
  <li class="li"><a href='{{ url "Categories" }}'>
      <h4 style="display: contents;">📑 分类</h4>
    </a></li>
  <li class="li"><a href='{{ url "Labels" }}'>
      <h4 style="display: contents;">🏷️ 标签</h4>
    </a></li>
</ul>
</header>
{{end}}