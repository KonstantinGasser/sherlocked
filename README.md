Repo has moved to [sherlock](https://github.com/KonstantinGasser/sherlock)

# simple password manager

## sherlocked
Is a very simple password manager I wrote. I came to the conclusion that I could not continue
maintaining passwords the way I was doing it so fare - After all it's 2020! <br>
Thus this command line password manager.


## usage
1. run `lock password` you will get asked to set a vault password - don't forget it
#### Available commands:
* `lock add` enter you vault password then the account name and its password
* `lock get <account-name>` retrieve a password (by default the password will not be printed but only copied to your clipboard
  * use the flag `-v` to print the password
* `lock ls` print all accounts stored in the vault
* `lock del <account-name>` delete account and password from vault
* `lock password` change the password of the vault
*  `lock gen` generate a random password (default length:8).
   * `--length <#>` override default length of password. Shorthand: `-l`
   * `--uppers <#>` set number of upper case letters. Shorthand: `-u`
   * `--numbers <#>` set number of numbers (0-9). Shorthand: `-n`
   * `--specials <#>` set number of special chars `+_-?.@#$%!`. Shorthand: `-s`
   * `--Create <account-name>` if set generated password will be mapped to the given user. Shorthand: `-C`
   * `--ignore char#1,char#2,char#n` if set chars will not be used in password !comma separated list!. Shorthand: `-i`


## things to know
Since I use this command line tool on a daily basis I will extend it with feature over the time - however it must be set this was a Saturday project (yes the commit are over a longer time span -> well it some things where just missing and the code needed some refactoring :)
