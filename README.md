# addRem
## Add reminders to go to an online class 
## Usage
- Note that you do have to give your terminal application administrative control to use addRem since it uses crontab to remind you.
- First run addRem: `./addRem`
- Then answer questions about the name of the event, online class link (e.g. Zoom link), date and time, and where you want the reminder scripts to be stored.

## TODO:
- [x] have user specify location of where reminders are stored
  - [x] have default
- [ ] have user say where terminal-notifier is stored
  - [ ] have mine as default
- [x] give ability for user to change message of notification
- [x] check that Y works for confirming zoom
- [ ] add ability to delete crontabs

## Known issues
- If you try to use environmental variables like $HOME when specifying the directory where reminders are stored, you first need to export the variable to use it in addRem. 
- I've currently only added support for $HOME, which you can use by setting it to "remHome" and then exporting remHome
  - For example: 
    ```bash 
    remHome=$HOME 
    export remHome    
    ./addRem
    ...
    Specify where you want to store reminders shell scripts:  $remHome
    ```
## Uses:
- terminal-notifier
- crontab
- bash


