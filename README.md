# simple password manager

## sherlocked
Password manager are no new thing there are many out there however,
I had two days to spare and was interested and learning something about encryption.
Perfect time! The way I was maintaining passwords currently simple was not 2020 like. Hence this small cmd tool.

<!-- ## installation
clone the repository with `git clone https://github.com/KonstantinGasser/sherlocked` then
place the executable `sherlocked/bin/lock` -->

## usage
1. run `touch $HOME/.sherlocked` // this will the vault file the encrypted password will be stored in
2. run `lock password` // this will set the password which will be used to encrypt/decrypt your vault <br>

After the setup you can use
-  `lock add` to add a new password to your vault
    * add `--override` to override an entry in the vault or to change a password for a given username
-  `lock get --user <username-mapped-to-password>`
    * above command will print the password to the terminal and copy it to your clipboard
    * add `--hide` to prevent the password to be print and also copied into your clipboard
- `lock list` lists all username/account names stored in the vault


## things to know
When you add a new password the old vault will be overwritten, however to avoid password loss there will be a backup created of the `.sherlocked` file with the suffix of a unix timestamp. In case I messed up something and the bug is about to ruin your Sunday you can run `mv .sherlocked-unixtimestamp .sherlocked` to get the old vault back.

## lastly
in case of password loss I can not take responsibility for it but would be happy if you tell what happened so I can fix it in post :D
