# tsgo
tsgo based on [mop](https://github.com/mop-tracker/mop)

## Installing

tsgo is implemented in Go and compiles down to a single executable file.

```BATCH
git clone github.com/saycv/tsgo
cd tsgo && make install
```

## Using

```BATCH

tsgo version
echo Use yahoo API
tsgo yahoo
echo Use QQ API
tsgo qq
echo Use Sina API
tsgo sina

```

For demonstration purposes tsgo comes preconfigured with a number of
stock tickers. You can easily change the default list by using the
following keyboard commands:

    +       Add stocks to the list.
    -       Remove stocks from the list.
    o       Change column sort order.
    g       Group stocks by advancing/declining issues.
    f       Set a filtering expression.
    F       Unset a filtering expression.
    ?       Display help screen.
    esc     Quit tsgo.
