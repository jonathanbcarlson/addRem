# addRem
## Add reminders to go to a Zoom class 

## TODO:
- [x] have user specify location of where reminders are stored
  - [x] have default
- [ ] have user say where terminal-notifier is stored
  - [ ] have mine as default
- [x] give ability for user to change message of notification
- [x] check that Y works for confirming zoom
- [ ] add ability to delete crontabs

## Known issues
- If you try to use environmental variables like $HOME when specifying the directory where reminders are stored, you first need to export the variable to use it in addrem. 
- I've currently only added support for $HOME, which you can use by setting it to "remHome" and then exporting remHome
  - For example: 
    ```bash 
    remHome=$HOME 
    export remHome    
    ./addrem
    ...
    Specify where you want to store reminders shell scripts:  $remHome
    ```
Uses:
- terminal-notifier
- crontab
- bash


