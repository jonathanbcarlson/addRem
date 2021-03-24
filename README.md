# addRem
## Add reminders to go to a Zoom class 
## Known issues
- If you try to use environmental variables like $HOME when specifying the directory where reminders are stored, you first need to export the variable to use it in addrem. 
- I've currently only added support for $HOME, which you can use by setting it to rem_Home and then exporting rem_Home
  - For example: 
    ```bash 
    rem_Home=$HOME 
    export rem_Home    
    ./addrem
    ...
    Specify where you want to store reminders shell scripts:  $rem_Home
    ```
Uses:
- terminal-notifier
- crontab
- bash

