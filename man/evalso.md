evalso(1) -- Interact with the Eval.so JSON API
===============================================

## SYNOPSIS

cat /path/to/file | `evalso` --language <language>  
`evalso` --status
`evalso` --list-evaluations (admin only)  
`evalso` --kill <evaluation id> (admin only)  
`evalso` --disable-user <username> (admin only)  
`evalso` --enable-user <username> (admin only)  

## DESCRIPTION

**evalso** is the commandline interface to the Eval.so JSON API. It allows you
to evaluate code, check on the status of the system, and (eventually) register
applications which can use the platform.

By default, `evalso` will evaluate code from STDIN.

Site administrators can also use `evalso` to list current evaluations, kill
evaluations, and disable and enable users.

## OPTIONS

  * `--server`:
    Configure which server to interact with. Useful for debugging against local
    instances of the API, or if DNS is going whacky for some reason.

## COPYRIGHT

Eval.so is Copyright (c) 2013 Ricky Elrod <ricky@elrod.me>
