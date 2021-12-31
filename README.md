# chorro

Chore scheduling made easy

## Development

Chorro is a
[monorepo](https://en.wikipedia.org/wiki/Monorepo#:~:text=In%20version%20control%20systems%2C%20a,stored%20in%20the%20same%20repository.&text=Many%20attempts%20have%20been%20made,other%2C%20newer%20forms%20of%20monorepos.)
managed by [turborepo](https://turborepo.org/).

```
chorro/
├─ .changeset/      # changeset files live in here
├─ apps/            # deployable apps live in here
│  ├─ foo/
│  ├─ bar/
├─ packages/        # reusable packages live in here
│  ├─ foo/
│  ├─ bar/
├─ scripts/         # node scripts used to manage the repo live in here
```

### Contributing

Chorro has only one main `git` branch - `main`. The `main` branch only accepts
changes made to it via
[pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests).
Chorro uses [changesets](https://github.com/changesets/changesets) to provide
semantic change management for each of the app.

Every pull request (PR) should be accompanied by the addition of a new
changeset. Luckily, adding a new changeset is easy:

```bash
$ yarn changeset
```

**NOTE: PRs should not be merged without also including a new changeset** \
We use [the changeset bot](https://github.com/apps/changeset-bot) to ensure that
our PRs include corresponding changesets.

#### Workflow

1.  Create a new branch off of `main`

    - if creating a new feature, use the `feature/` branch name prefix (e.g.
      `feature/build-the-thing`)
    - if fixing a bug, use the `bugfix/` branch name prefix (e.g.
      `bugfix/fix-the-thing`)
    - if hotfixing an issue, use the `hotfix/` branch name prefix (e.g.
      `hotfix/patch-the-thing`)

    ```bash
    $ git pull
    $ git checkout -b feature/trust-the-process
    ```

2.  Make changes via `git commit`, continually rebasing against `origin/main`

    ```bash
    $ git commit -am 'Made some changes'
    $ git pull --rebase origin master
    ```

3.  Add changesets along the way describing any changes recognizable to humans

    ```bash
    $ yarn changeset
    ```

4.  When you're ready for review, publish your branch to the repo, and create a
    [pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests)
    against `origin/main`

    ```bash
    $ git push
    ```

5.  Once you merge your PR, use `yarn reset` to do post-merge cleanup locally

    ```bash
    $ yarn reset
    ```
