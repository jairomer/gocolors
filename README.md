# gocolors

## What is this?

A program that injects coloring characters for linux shells.

For now it will only highlight success and error messages for GTest.

## Why?

I want color in my life.

The purpose of this utility is to pass the output of another program to it using unix pipes.

```
echo "This is not OK, buddy" | gocolors
echo "You have an ERROR there" | gocolors

```

## How?

To run this you will need a working go environment:
```
# To generate the binary.
make build

# To generate and move the binary to /bin
sudo make install
```
