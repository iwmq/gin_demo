{{define "base"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/css/bootstrap.css">
    <link rel="stylesheet" href="/assets/css/style.css">
    <script src="/assets/js/jquery-3.5.1.js"></script>
    <script src="/assets/js/bootstrap.js"></script>
    <script src="/assets/js/vue.js"></script>
    <title>Gin Demo</title>
</head>
<body>
    {{ block "content" . }} {{ end }}
</body>
</html>
{{end}}