<html>
    <head>
    <title>Login Page</title>
    </head>
    <body>
        <form action="/register" method= "post">
            Username: <input type="text" name="username">
            Password: <input type="text" name="password">
            <input type="submit" value="Register Account">
        </form>
        <p> {{.usercount}} users! </p>
    </body>
</html>