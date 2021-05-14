# Permutate

Given a file with a list of domains, e.g:

```
google.com
infected-database.org
```

And a list of subdomains you want to prefix with the domain list, e.g.

```
vpn
remote
www
mail
```

`Permutate` will create the following list for you _(the real output is not ordered)_:

```
vpn.google.com
remote.google.com
www.google.com
mail.google.com
vpn.infected-database.org
remote.infected-database.org
www.infected-database.org
mail.infected-database.org
```

Permutate does not perform DNS lookups, it simply generates your list, what happens after is up to you :)

## Contributing
Any feedback or ideas are welcome! Want to improve something? Create a pull request!

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Configure pre commit checks: `pre-commit install`
4. Commit your changes: `git commit -am 'Add some feature'`
5. Push to the branch: `git push origin my-new-feature`
6. Submit a pull request :D