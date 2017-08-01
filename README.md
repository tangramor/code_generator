# code_generator
A code (coupon, security code, Anti-counterfeiting code...) generator by golang

Setup uuid first: `go get github.com/nu7hatch/gouuid`

## Usage of code_generator:

  -len int
        The length of code (default 18)
        
  -num int
        The total number of codes (default 100)
        
  -prefix
        If use prefix with 8 charaters (default true)
        
```
code_generator -len=20 -num=200 -prefix=false
```
