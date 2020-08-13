<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    ユーザ名:<input type="text" name="username">
    パスワード:<input type="password" name="password">
    年齢:<input type="text" name="age">
    <select name="fruit">
        <option value="apple">apple</option>
        <option value="pear">pear</option>
        <option value="banane">banane</option>
    </select>

    <input type="checkbox" name="interest" value="football">サッカー
    <input type="checkbox" name="interest" value="basketball">バスケットボール
    <input type="checkbox" name="interest" value="tennis">テニス
    ユーザ名:<input type="text" name="username">
    パスワード:<input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="ログイン">
</form>
</body>
</html>