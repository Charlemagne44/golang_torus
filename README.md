# Golang Torus

Golang torus is an adaptation of Andy Sloane's [C-Donut](https://www.a1k0n.net/2011/07/20/donut-math.html) as a CLI written in Go

## Installation

Golang torus is released as a binary for all platforms under the releases tab. Download the binary for their respective platform, and add put it in a directory included in your PATH env variable

### Linux Example

```bash
wget https://github.com/Charlemagne44/golang_torus/releases/download/1.0/donut_linux_amd64```

# optional rename to just "donut" -> make sure destination dir is in your PATH
sudo mv donut_linux_amd64 /usr/bin/donut
sudo chmod +x /usr/bin/donut

OR

sudo mv donut_linux_amd64 ~/bin/donut
sudo chmod +x ~/bin/donut

```

If you are unsure of how to add a directory to your path, or if the example directories are already in your path, please follow the guides below according to your platform

[Windows](https://www.autodesk.com/support/technical/article/caas/sfdcarticles/sfdcarticles/Adding-folder-path-to-Windows-PATH-environment-variable.html)

[MacOS](https://www.cyberciti.biz/faq/appleosx-bash-unix-change-set-path-environment-variable/)

[Linux](https://phoenixnap.com/kb/linux-add-to-path)

## Usage

```
./donut -h - print help menu
./donut -w - the width / height of the donut in terminal characters
```

If you selected a width (-w) larger than your current terminal screen can fit, there will be overflows of characters and the animation will not appear smooth. To fix this, simply zoom out in your terminal. This shortcut in most terminal emulators is typically (ctrl shift -) or (ctrl -)

![](https://media.giphy.com/media/ycobgCQ3yMn4eJ949Q/giphy.gif)

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)