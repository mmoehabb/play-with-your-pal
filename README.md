Simple Go application to play local games together remotely; by sharing your the keyboard with your pals!

# Usage

If you have [Go](https://go.dev/) installed on your machine, you may build a binary executable with `go build .` command. But make sure to install [Templ](https://templ.guide/quick-start/installation) and compile templ files first:

```bash
$ go install github.com/a-h/templ/cmd/templ@latest
$ templ generate
$ go build .
```

Otherwise, you can directly download builds from [releases page](https://github.com/mmoehabb/play-with-your-pal/releases).

Once you have the executable file, you can straightforwardly run the application via your command line interface and the your application server will start listening on port `8080` by default. You may the port and/or other parameters with the following flags:

> Remember: You need to port-forward connection, to the specified port, from your router to your local machine.

```bash
$ ./pwyp -h
Usage of ./pwyp:
  -noscreen
        use this flag to disable sharing your screen.
  -password string
        the password of your session. (default "empty")
  -port int
        the port on which the server is listening. (default 8080)
  -quality int
        the quality of the video stream. (default 75)
```

> It's recommended to use -noscreen flag and use the app to share only keyboard commands; as streaming implementation might not be quite suitable.
> If you used screen sharing and it seemed to be buggy on phones or old poor hardwares, you may want to reduce the quality with -quality flag: try `-quality 25`.
