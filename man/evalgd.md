evalgd(1) -- Interact with the Eval.gd JSON API
===============================================

## SYNOPSIS

cat /path/to/file | `evalgd` --language <language>  
`evalgd` --status  
`evalgd` --list-evaluations (admin only)  
`evalgd` --kill <evaluation id> (admin only)  
`evalgd` --disable-user <username> (admin only)  
`evalgd` --enable-user <username> (admin only)  

## DESCRIPTION

**evalgd** is the commandline interface to the Eval.gd JSON API. It allows you
to evaluate code, check on the status of the system, and (eventually) register
applications which can use the platform.

By default, `evalgd` will evaluate code from STDIN.

Site administrators can also use `evalgd` to list current evaluations, kill
evaluations, and disable and enable users.

## COPYRIGHT

Eval.gd is Copyright (c) 2013 Ricky Elrod <ricky@elrod.me>
