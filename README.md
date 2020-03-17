<h1 align="center">Welcome to SimpleGolangChatServer 👋</h1>
<p>
</p>

> 	This Program is a very simple chat Server written in go, that works following a room concept, for instance the server create/manage rooms and users connect to it. Each users can connect to a room or join one if there is an available spot. The server checks if there is an unoccupied spot in active rooms, if not, it will create a new one

## Usage

```sh
USAGE
        go run main.go [options]
OPTIONS
        -S, --server <[ipaddr], port>
                start server using the ip and port pass in argv
        -C, --client <ipaddr, port>
                start client on the ip and port pass in argv
        -h, --help
                Display the program usage
```

## Run server

```sh
  go run main.go --server <[ipaddr], port>
```

## Run Client

```sh
  go run main.go --client <ipaddr, port>
```

## Author

👤 **hadi-ilies**

* Website: https://hadibereksi.fr/
* Github: [@hadi-ilies](https://github.com/hadi-ilies)
* LinkedIn: [@https:\/\/www.linkedin.com\/in\/hadibereksi\]

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
