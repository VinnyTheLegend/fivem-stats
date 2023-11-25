This hosts a web server using the Gin web package for Go. It connects to a FiveM QbCore database and serves dynamic web pages of the character and vehicle data.

ENV Variables:
DBCONNECT={ip/url for mysql database}
DBNAME={name of the database}
DBUSER={username for mysql database}
DBPASS={password for mysql database}
PORT={port for web server}