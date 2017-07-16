<html>
<head>
    <title></title>
</head>
<body>
<form action="/show" method="post">
    {{ range . }}
    {{ . }}
    <br>
    {{ end }}
    <br>
    事情:<input type="text" name="message">
    <input type="submit" value="添加">
</form>
</body>
</html>