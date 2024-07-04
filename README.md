# ASCII-ART-OUTPUT
## Description
This program takes a string and converts it into ASCII art, writing the result into a `specified file`. The file name is provided using the `--output=<fileName.txt>` flag.

We also have a default output file called `output.txt` that stores the output if no flag has been specified. 


## Instalation
1. Download and set up [go](https://go.dev/doc/install) in your machine.

2. Clone the program in your machine using the following command:

```
git clone https://learn.zone01kisumu.ke/git/mdudi/ascii-art-output.git
cd ascii-art-output
```


## Usage 

 To run the program do the following:

### 1. Writting the output using standard.txt font in our file:

```
go run . --output=output.txt hello standard
cat -e output.txt
```
#### output
```
 _                _    _            $
| |              | |  | |           $
| |__      ___   | |  | |    ___    $
|  _ \    / _ \  | |  | |   / _ \   $
| | | |  |  __/  | |  | |  | (_) |  $
|_| |_|   \___|  |_|  |_|   \___/   $
                                    $
                                    $
```
### 2. Writting the output using shadow.txt font in our file:

```
go run . --output=output.txt hello shadow
cat -e output.txt
```
#### output
```
                                      $
_|                  _|  _|            $
_|_|_|      _|_|    _|  _|    _|_|    $
_|    _|  _|_|_|_|  _|  _|  _|    _|  $
_|    _|  _|        _|  _|  _|    _|  $
_|    _|    _|_|_|  _|  _|    _|_|    $
                                      $
                                      $
```
### 3. Writting the output using thinkertoy.txt font in our file:

```
go run . --output=output.txt hello thinkertoy
cat -e output.txt
```
#### output
```
                      $
o          o  o       $
|          |  |       $
O--o  o-o  |  |  o-o  $
|  |  |-'  |  |  | |  $
o  o  o-o  o  o  o-o  $
                      $
                      $
```
## Testing

1. Get into the folder that contains the tests:
```
cd ascii
```
2. Run the tests:
```
go test -v
```

## Contributions

Contributions to this project are welcome! If you have ideas for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request. Your contributions help make this project better for everyone. Thank you for your support and contributions!

## Authors

[Martin Dudi](https://learn.zone01kisumu.ke/git/mdudi)

[Bernad Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)