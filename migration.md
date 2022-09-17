**How to migrate from starport to Ignite CLI**

***Note**: Depending on your user permissions, run below commands with or without `sudo`.*

1. Pull the migration branch or create a branch from dev branch (after the approval of migration pull request) which already have all necessary code changes (updated dependencies and vue files) for Ignite CLI 
2. Remove the Starport cli from your local machine with `rm $(which starport)`. 
3. Download the Ignite CLI using `curl https://get.ignite.com/cli | bash`
4. Move the ignite executable to /usr/local/bin/ using `sudo mv ignite /usr/local/bin/`
5. Test the chain by running it using `ignite chain serve`.
