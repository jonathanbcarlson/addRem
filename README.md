# addRem
## Add reminders to go to an online class 
## Usage
- Note that you do have to give your terminal application administrative control to use addRem since it uses crontab to remind you.
- First run addRem: `./addRem`
- Then answer questions about:
  - The name of the event
  - The online class link (e.g. Zoom link)
  - Which days and at what time you want to be reminded at 
  - Where you want the reminder scripts to be stored
## Dependencies:
- [terminal-notifier](https://github.com/julienXX/terminal-notifier)
- [cron](https://en.wikipedia.org/wiki/Cron)
- [bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell))
## TODO:
- [x] have user specify location of where reminders are stored
  - [x] have default
- [ ] have user say where terminal-notifier is stored
  - [ ] have mine as default
- [x] give ability for user to change message of notification
- [x] check that Y works for confirming zoom
- [ ] add ability to delete crontabs
- [ ] add abiity for user to specify not just per week for crontab (e.g. a weekly scheduled meeting) but any day of the month (basically use all of crontab)

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



