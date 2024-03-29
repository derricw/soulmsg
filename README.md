# soulmsg

Generate random Dark Souls messages. Use for inspiration, fortunes, or message of the day!

```
$ soulmsg
try miracles
```

Try a live version [here](https://derricw.github.io/soulmsg/)

or with curl:

```
curl -L fir3.link
```

Try `-h` for options.

### build

```bash
go build
```

### install

```bash
go install github.com/derricw/soulmsg@latest
```

### examples

Make 10 messages, allow conjuctions.
```bash
soulmsg -c -n 10
```

Make messages forever.
```bash
soulmsg -c -n -1
```

Grab the first message containing something I want.
```bash
soulmsg -c -n -1 | grep woman | head -n 1
```

Cowsay
```bash
$ soulmsg | cowsay
 ________________________________
< be wary of close-ranged battle >
 --------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```
